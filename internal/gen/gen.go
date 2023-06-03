package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

type goStruct struct {
	Node      *node
	FullName  string
	BuildDef  buildDef
	NextNodes []*node
	Variadic  bool

	MultipleToken bool
}

type buildDef struct {
	MethodName string
	Command    []string
	Parameters []parameter
}

func (d buildDef) hash() uint32 {
	hash := crc32.NewIEEE()
	if err := gob.NewEncoder(hash).Encode(d); err != nil {
		panic(err)
	}
	return hash.Sum32()
}

type parameter struct {
	Name string
	Type string
}

type command struct {
	Group     string     `json:"group"`
	Since     string     `json:"since"`
	Arguments []argument `json:"arguments"`
}

type argument struct {
	Name      any        `json:"name"`
	Type      any        `json:"type"`
	Command   string     `json:"command"`
	Enum      []string   `json:"enum"`
	Block     []argument `json:"block"`
	Arguments []argument `json:"arguments"`
	Multiple  bool       `json:"multiple"`
	Optional  bool       `json:"optional"`
	Variadic  bool       `json:"variadic"`

	MultipleToken bool `json:"multiple_token"`
}

type node struct {
	Group  string
	Parent *node
	Child  *node
	Next   *node
	Cmd    command
	Arg    argument
	Root   bool
}

func (n *node) FindRoot() (root *node) {
	root = n
	for root.Parent != nil {
		root = root.Parent
	}
	return
}

//gocyclo:ignore
func (n *node) GoStructs() (out []goStruct) {
	fn := n.FullName()
	// fix for XGROUP and XADD
	if len(n.Arg.Enum) == 2 && ((n.Arg.Enum[0] == "id" && n.Arg.Enum[1] == "$") || (n.Arg.Enum[0] == "*" && n.Arg.Enum[1] == "ID")) {
		n.Arg.Enum = nil
		n.Arg.Name = "id"
		n.Arg.Type = "string"
		if !strings.HasSuffix(fn, "Id") {
			fn += "Id"
		}
	}

	// fix for TS.MRANGE, TS.MREVRANGE, TS.MGET, TS.QUERYINDEX
	if len(n.Arg.Enum) >= 2 && (n.Arg.Enum[0] == "l=v" && n.Arg.Enum[1] == "l!=v") {
		n.Arg.Enum = nil
		n.Arg.Name = "filter"
		n.Arg.Type = "string"
		if !strings.HasSuffix(fn, "Filter") {
			fn += "Filter"
		}
	}

	if len(n.Arg.Block) > 0 {
		panic("GoStructs should not be called on Block node")
	} else if len(n.Arg.Enum) > 0 {
		for _, e := range n.Arg.Enum {
			s := goStruct{
				Node:     n,
				FullName: fn,
				BuildDef: buildDef{
					MethodName: name(e),
					Command:    nil,
					Parameters: nil,
				},
				Variadic:      n.Variadic() && !n.MultipleToken(),
				MultipleToken: n.MultipleToken(),
				NextNodes:     n.NextNodes(),
			}
			if len(n.Arg.Command) != 0 {
				if !strings.HasSuffix(s.FullName, name(n.Arg.Command)) {
					s.FullName += name(n.Arg.Command)
				}
				s.BuildDef.MethodName = name(n.Arg.Command) + name(e)
				s.BuildDef.Command = strings.Split(n.Arg.Command, " ")
			}
			if !strings.HasSuffix(s.FullName, name(e)) {
				s.FullName += name(e)
			}
			cmds := strings.Split(e, " ")
			if len(cmds) == 1 {
				if cmds[0] == "VECTOR" {
					s.BuildDef.Command = append(s.BuildDef.Command, cmds...)
					s.BuildDef.Parameters = []parameter{
						{Name: "algo", Type: "string"},
						{Name: "nargs", Type: "integer"},
						{Name: "args", Type: "...interface{}"},
					}
				} else {
					s.BuildDef.Command = append(s.BuildDef.Command, cmds...)
					s.BuildDef.Parameters = nil
				}
			} else {
				switch cmds[1] {
				case "name", "category", "pattern":
					s.BuildDef.Command = append(s.BuildDef.Command, cmds[:1]...)
					s.BuildDef.Parameters = []parameter{{Name: lcFirst(name(cmds[1])), Type: "string"}}
				case "seconds", "milliseconds", "timestamp", "milliseconds-timestamp":
					s.BuildDef.Command = append(s.BuildDef.Command, cmds[:1]...)
					s.BuildDef.Parameters = []parameter{{Name: lcFirst(name(cmds[1])), Type: "integer"}}
				case "*":
					// fix for FT.AGGREGATE
					if cmds[0] == "LOAD" {
						s.BuildDef.Command = append(s.BuildDef.Command, cmds...)
						s.BuildDef.Parameters = nil
					}
				case "label1":
					// fix for TS.MRANGE, TS.MREVRANGE, TS.MGET
					if cmds[0] == "SELECTED_LABELS" {
						s.FullName = strings.TrimRight(s.FullName, "Label1")
						s.BuildDef.MethodName = strings.TrimRight(s.BuildDef.MethodName, "Label1")
						s.BuildDef.Command = append(s.BuildDef.Command, cmds[0])
						s.BuildDef.Parameters = []parameter{{Name: "labels", Type: "[]interface{}"}}
					}
				default:
					panic("unknown enum " + cmds[1])
				}
			}
			out = append(out, s)
		}
		return
	} else {
		s := goStruct{
			Node:     n,
			FullName: fn,
			BuildDef: buildDef{
				MethodName: n.Name(),
				Parameters: nil,
			},
			Variadic:      n.Variadic() && !n.MultipleToken(),
			MultipleToken: n.MultipleToken(),
			NextNodes:     n.NextNodes(),
		}
		if len(n.Arg.Command) != 0 {
			s.BuildDef.Command = strings.Split(n.Arg.Command, " ")
		}

		switch nm := n.Arg.Name.(type) {
		case string:
			if s.FullName == "FtCreatePrefixPrefix" && nm == "count" && n.Arg.Command == "PREFIX" {
				s.FullName = "FtCreatePrefixCount"
			}
			s.BuildDef.Parameters = []parameter{{Name: lcFirst(name(nm)), Type: n.Arg.Type.(string)}} // not change to go type here, change at render
		case []any:
			for i, nn := range nm {
				s.BuildDef.Parameters = append(s.BuildDef.Parameters, parameter{Name: lcFirst(name(nn.(string))), Type: n.Arg.Type.([]any)[i].(string)})
			}
		default:
			if n.Arg.Type == nil || (n.Arg.Type != nil && n.Arg.Type.(string) == "command") {
				// ignore
			} else {
				panic("unknown arg name")
			}
		}
		out = append(out, s)
	}
	return
}

func (n *node) MultipleToken() bool {
	return n.Arg.MultipleToken
}

func (n *node) Variadic() bool {
	return n.Arg.Multiple || n.Arg.Variadic
}

func (n *node) FullName() (out string) {
	if n.Parent != nil {
		return n.Parent.FullName() + n.Name()
	}
	return n.Name()
}

func (n *node) Name() (out string) {
	var tokens []string
	if n.Arg.Command != "" {
		tokens = append(tokens, name(n.Arg.Command))
	} else {
		switch n := n.Arg.Name.(type) {
		case string:
			tokens = append(tokens, name(n))
		case []any:
			for _, nn := range n {
				tokens = append(tokens, name(nn.(string)))
			}
		}
		if len(tokens) == 0 {
			if n.Child != nil {
				tokens = append(tokens, n.Child.Name())
			}
		}
	}

	duplicate := map[string]bool{}
	for _, t := range tokens {
		if len(t) == 0 || duplicate[t] {
			continue
		}
		duplicate[t] = true
		out += t
	}
	return
}

func (n *node) NextNodes() (nodes []*node) {
	defer func() {
		deduped := make([]*node, 0, len(nodes))
	next:
		for _, n := range nodes {
			for _, nn := range deduped {
				if n == nn {
					continue next
				}
			}
			deduped = append(deduped, n)
		}
		nodes = deduped
	}()

	if n.Child == nil && n.Variadic() {
		nodes = append(nodes, n)
	}

	if n.Child != nil {
		nodes = append(nodes, blockEntries(n)...)
	}

	parent := n
	for parent != nil {
		next := parent.Next
		for next != nil {
			nodes = append(nodes, next)
			if !next.Arg.Optional {
				return nodes
			}
			next = next.Next
		}

		parent = parent.Parent
		// block variadic
		if parent != nil && parent.Variadic() {
			nodes = append(nodes, blockEntries(parent)...)
		}
		if parent != nil && parent.Root {
			break // don't climb to root
		}
	}
	return nodes
}

func (n *node) Walk(fn func(node *node)) {
	next := n
	for next != nil {
		if next.Child != nil {
			next.Child.Walk(fn)
		} else {
			fn(next)
		}
		next = next.Next
	}
}

func main() {

	var structs = map[string]map[string]goStruct{}

	defs, err := filepath.Glob("*.json")
	if err != nil {
		panic(err)
	}

	for _, p := range defs {
		raw, err := os.ReadFile(p)
		if err != nil {
			panic(err)
		}

		var commands = map[string]command{}
		if err := json.Unmarshal(raw, &commands); err != nil {
			panic(err)
		}

		var roots []string
		nodes := map[string]*node{}
		for k, cmd := range commands {
			if cmd.Group == "" {
				panic(k + " no group")
			}
			root := &node{Group: cmd.Group, Cmd: cmd, Arg: argument{Name: k, Command: k, Type: "command"}, Root: true}
			root.Next = makeChildNodes(root, cmd.Arguments)
			roots = append(roots, k)
			if _, ok := nodes[k]; ok {
				panic(k + " conflict")
			}
			nodes[k] = root
		}
		sort.Strings(roots)

		for _, name := range roots {
			n := nodes[name]
			g := n.Group
			if _, ok := structs[g]; !ok {
				structs[g] = make(map[string]goStruct)
			}
			n.Walk(func(n *node) {
				for _, s := range n.GoStructs() {
					if v, ok := structs[g][s.FullName]; ok {
						panic("struct conflict " + v.FullName)
					}
					structs[g][s.FullName] = s
				}
			})
		}
	}

	for g, structs := range structs {
		gfname := "../cmds/gen_" + g + ".go"
		tfname := "../cmds/gen_" + g + "_test.go"
		gf, err := os.Create(gfname)
		if err != nil {
			panic(err)
		}
		generate(gf, structs)
		if err := gf.Close(); err != nil {
			panic(err)
		}
		if err := exec.Command("gofmt", "-w", gfname).Run(); err != nil {
			panic(err)
		}
		if err := exec.Command("goimports", "-w", gfname).Run(); err != nil {
			panic(err)
		}
		tf, err := os.Create(tfname)
		if err != nil {
			panic(err)
		}
		tests(tf, structs, g)
		if err := tf.Close(); err != nil {
			panic(err)
		}
		if err := exec.Command("gofmt", "-w", tfname).Run(); err != nil {
			panic(err)
		}
		if err := exec.Command("goimports", "-w", tfname).Run(); err != nil {
			panic(err)
		}
	}
}

func tests(f io.Writer, structs map[string]goStruct, prefix string) {
	var names []string
	for name := range structs {
		names = append(names, name)
	}
	sort.Strings(names)

	var pathes [][]goStruct
	for _, name := range names {
		s := structs[name]
		if !s.Node.Root {
			continue
		}
		pathes = makePath(s, nil, pathes)
	}

	fmt.Fprintf(f, "// Code generated DO NOT EDIT\n\npackage cmds\n\n")

	fmt.Fprintf(f, "import \"testing\"\n\n")

	mod := 100

	for i, p := range pathes {
		if i%mod == 0 {
			fmt.Fprintf(f, "func %s%d(s Builder) {\n", prefix, i/mod)
		}
		printPath(f, "s", p, "GetArgs")
		if i%mod == mod-1 || i == len(pathes)-1 {
			fmt.Fprintf(f, "}\n\n")
		}
	}

	fmt.Fprintf(f, "func TestCommand_%s(t *testing.T) {\n", prefix)
	for i := 0; i <= len(pathes)/mod; i++ {
		fmt.Fprintf(f, "\tt.Run(\"%d\", func(t *testing.T) { %s%d(Builder{}) })\n", i, prefix, i)
	}
	fmt.Fprintf(f, "}\n\n")
}

var pathmark = map[string]bool{}
var blockmark = map[*node]bool{}

func makePath(s goStruct, path []goStruct, pathes [][]goStruct) [][]goStruct {
	path = append(path, s)
	nexts := s.Node.NextNodes()
	if pathmark[s.FullName] && allOptional(s.Node, nexts) {
		pathes = append(pathes, path)
		return pathes
	}
	if len(path) < 100 {
		for _, n := range nexts {
			if s.Node == n {
				if !s.Variadic && !s.MultipleToken {
					path = append(path, s)
				}
				continue
			}
			if n.Parent != nil && n.Parent.Child == n {
				if blockmark[n] {
					continue
				}
				blockmark[n] = true
			}
			nodes := []*node{n}
			if n.Child != nil {
				nodes = blockEntries(n)
			}
			for _, nn := range nodes {
				for _, ss := range nn.GoStructs() {
					clone := make([]goStruct, len(path))
					copy(clone, path)
					pathes = makePath(ss, clone, pathes)
					if pathmark[s.FullName] {
						return pathes
					}
				}
			}
		}
		pathmark[s.FullName] = true
	}
	if allOptional(s.Node, nexts) {
		pathes = append(pathes, path)
	}
	return pathes
}

func testParams(defs []parameter) string {
	var params []string
	for _, param := range defs {
		switch toGoType(param.Type) {
		case "[]interface{}":
			params = append(params, `[]interface{}{"1"}`)
		case "interface{}":
			params = append(params, `"1"`)
		case "int64", "uint64", "float64":
			params = append(params, `1`)
		}
	}
	return strings.Join(params, ", ")
}

func printPath(f io.Writer, receiver string, path []goStruct, end string) {
	fmt.Fprintf(f, "\t%s.%s()", receiver, path[0].BuildDef.MethodName)
	for _, s := range path[1:] {
		fmt.Fprintf(f, ".%s(", s.BuildDef.MethodName)
		if (len(s.BuildDef.Parameters) != 1 && s.Variadic) || s.MultipleToken {
			fmt.Fprintf(f, ").%s(", s.BuildDef.MethodName)
		}
		fmt.Fprintf(f, "%s)", testParams(s.BuildDef.Parameters))
		if s.Variadic || s.MultipleToken {
			fmt.Fprintf(f, ".%s(%s)", s.BuildDef.MethodName, testParams(s.BuildDef.Parameters))
		}
	}
	fmt.Fprintf(f, ".%s()\n", end)
}

func generate(f io.Writer, structs map[string]goStruct) {
	var names []string
	for name := range structs {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Fprintf(f, "// Code generated DO NOT EDIT\n\npackage cmds\n\n")
	fmt.Fprintf(f, "import %q\n\n", "strconv")

	for _, name := range names {
		s := structs[name]

		fmt.Fprintf(f, "type %s Completed\n\n", s.FullName)

		if s.Node.Root {
			printRootBuilder(f, s)
		}

		dedupe := make(map[string]uint32)
		for _, next := range s.NextNodes {
			nodes := []*node{next}
			if next.Child != nil {
				nodes = blockEntries(next)
			}
			for _, nn := range nodes {
				for _, ss := range nn.GoStructs() {
					hash := ss.BuildDef.hash()
					if h, ok := dedupe[ss.BuildDef.MethodName]; ok {
						if h != hash {
							panic("same method but different hash")
						}
					} else {
						dedupe[ss.BuildDef.MethodName] = hash
						printBuilder(f, s, ss)
					}
				}
			}
		}

		if allOptional(s.Node, s.NextNodes) {
			printFinalBuilder(f, s, "Build", "Completed")
		}
	}
}

func allOptional(s *node, nodes []*node) bool {
	for _, n := range nodes {
		if s == n {
			continue
		}
		if n.Parent != nil && n.Parent.Child == n {
			continue
		}
		if !n.Arg.Optional {
			return false
		}
	}
	return true
}

func toGoType(paramType string) string {
	switch paramType {
	case "[]interface{}": // TODO hack for TS.MRANGE, TS.MREVRANGE, TS.MGET
		return "[]interface{}"
	case "...interface{}": // TODO hack for FT.CREATE VECTOR
		return "...interface{}"
	case "key", "string", "pattern", "type":
		return "interface{}"
	case "double":
		return "float64"
	case "integer", "posix time":
		return "int64"
	case "unsigned integer":
		return "uint64"
	default:
		panic("unknown param type " + paramType)
	}
}

func toGoName(paramName string) string {
	if paramName == "type" {
		return "typ"
	}
	return paramName
}

func printRootBuilder(w io.Writer, root goStruct) {
	fmt.Fprintf(w, "func (b Builder) %s() (c %s) {\n", root.FullName, root.FullName)

	var appends []string
	for _, cmd := range root.BuildDef.Command {
		appends = append(appends, fmt.Sprintf(`"%s"`, cmd))
	}

	fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
	fmt.Fprintf(w, "\treturn c\n")
	fmt.Fprintf(w, "}\n\n")
}

func printFinalBuilder(w io.Writer, parent goStruct, method, ss string) {
	fmt.Fprintf(w, "func (c %s) %s() %s {\n", parent.FullName, "GetArgs", "[]interface{}")
	fmt.Fprintf(w, "\treturn c.cs\n")
	fmt.Fprintf(w, "}\n\n")
}

//gocyclo:ignore
func printBuilder(w io.Writer, parent, next goStruct) {
	fmt.Fprintf(w, "func (c %s) %s(", parent.FullName, next.BuildDef.MethodName)
	if len(next.BuildDef.Parameters) == 1 && next.Variadic {
		fmt.Fprintf(w, "%s ...%s", toGoName(next.BuildDef.Parameters[0].Name), toGoType(next.BuildDef.Parameters[0].Type))
	} else if (next.Variadic || next.MultipleToken) && parent.FullName != next.FullName {
		// no parameter
	} else {
		for i, param := range next.BuildDef.Parameters {
			fmt.Fprintf(w, "%s %s", toGoName(param.Name), toGoType(param.Type))
			if i != len(next.BuildDef.Parameters)-1 {
				fmt.Fprintf(w, ", ")
			}
		}
	}
	fmt.Fprintf(w, ") %s {\n", next.FullName)

	var appends []string

	for _, cmd := range next.BuildDef.Command {
		if cmd == "" {
			panic("unexpected empty command on " + next.FullName)
		}
		if cmd == `""` {
			appends = append(appends, `""`)
		} else {
			if !(len(next.BuildDef.Parameters) != 1 && next.Variadic && parent.FullName == next.FullName) {
				appends = append(appends, fmt.Sprintf(`"%s"`, cmd))
			}
		}
	}

	if len(appends) == 0 && next.Variadic && len(next.BuildDef.Parameters) == 1 && toGoType(next.BuildDef.Parameters[0].Type) == "interface{}" {
		appends = append(appends, toGoName(next.BuildDef.Parameters[0].Name)+"...")
		fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
	} else if len(next.BuildDef.Parameters) != 1 && (next.Variadic || next.MultipleToken) && parent.FullName != next.FullName {
		// no parameter
		if len(appends) != 0 && !next.MultipleToken {
			fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
		}
	} else if !(next.MultipleToken && parent.FullName != next.FullName) {
		allstring := true
		for _, p := range next.BuildDef.Parameters {
			if toGoType(p.Type) != "interface{}" {
				allstring = false
				break
			}
		}
		if allstring && !next.Variadic {
			for _, p := range next.BuildDef.Parameters {
				appends = append(appends, toGoName(p.Name))
			}
			fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
		} else {
			if len(next.BuildDef.Parameters) == 1 && next.Variadic {
				if len(appends) != 0 {
					fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
				}
				if toGoType(next.BuildDef.Parameters[0].Type) == "interface{}" {
					fmt.Fprintf(w, "\tc.cs = append(c.cs, %s...)\n", toGoName(next.BuildDef.Parameters[0].Name))
				} else {
					fmt.Fprintf(w, "\tfor _, n := range %s {\n", toGoName(next.BuildDef.Parameters[0].Name))
					switch toGoType(next.BuildDef.Parameters[0].Type) {
					case "float64":
						fmt.Fprintf(w, "\t\tc.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))\n")
					case "int64":
						fmt.Fprintf(w, "\t\tc.cs = append(c.cs, strconv.FormatInt(n, 10))\n")
					case "uint64":
						fmt.Fprintf(w, "\t\tc.cs = append(c.cs, strconv.FormatUint(n, 10))\n")
					default:
						panic("unexpected param type " + next.BuildDef.Parameters[0].Type)
					}
					fmt.Fprintf(w, "\t}\n")
				}
			} else {
				var follows []string
				for _, p := range next.BuildDef.Parameters {
					switch toGoType(p.Type) {
					case "float64":
						appends = append(appends, fmt.Sprintf("strconv.FormatFloat(%s, 'f', -1, 64)", toGoName(p.Name)))
					case "int64":
						appends = append(appends, fmt.Sprintf("strconv.FormatInt(%s, 10)", toGoName(p.Name)))
					case "uint64":
						appends = append(appends, fmt.Sprintf("strconv.FormatUint(%s, 10)", toGoName(p.Name)))
					case "interface{}":
						appends = append(appends, toGoName(p.Name))
					case "[]interface{}": // TODO hack for TS.MRANGE, TS.MREVRANGE, TS.MGET
						follows = append(follows, toGoName(p.Name)+"...")
					case "...interface{}": // TODO hack for FT.CREATE VECTOR
						follows = append(follows, toGoName(p.Name)+"...")
					default:
						panic("unexpected param type " + next.BuildDef.Parameters[0].Type)
					}
				}
				fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
				for _, follow := range follows {
					fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", follow)
				}
			}
		}
	}

	if parent.FullName == next.FullName {
		fmt.Fprintf(w, "\treturn c\n")
	} else {
		fmt.Fprintf(w, "\treturn (%s)(c)\n", next.FullName)
	}
	fmt.Fprintf(w, "}\n\n")
}

func makeChildNodes(parent *node, args []argument) (first *node) {
	if len(args) == 0 {
		return nil
	}
	var nodes []*node
	for _, arg := range args {

		if len(arg.Arguments) != 0 {
			arg.Block = arg.Arguments
		}
		if arg.Type == "enum" && len(arg.Enum) == 0 && arg.Command != "" {
			arg.Enum = []string{arg.Command}
			arg.Command = ""
		}
		if arg.Type == "oneof" && len(arg.Block) == 0 && arg.Command != "" {
			arg.Enum = []string{arg.Command}
		}

		nodes = append(nodes, &node{Parent: parent, Arg: arg})
	}
	for i, node := range nodes {
		node.Child = makeChildNodes(node, node.Arg.Block)
		if len(node.Arg.Block) != 0 && node.Arg.Command != "" {
			if node.Child.Arg.Optional {
				panic("unexpected block command with optional child")
			}
			if node.Child.Arg.Command == "" {
				node.Child.Arg.Command = node.Arg.Command
			} else {
				node.Child.Arg.Command = node.Arg.Command + " " + node.Child.Arg.Command
			}
		}
		if i != len(nodes)-1 {
			node.Next = nodes[i+1]
		}
	}
	return nodes[0]
}

func blockEntries(block *node) (nodes []*node) {
	for child := block.Child; child != nil && child.Parent == block; child = child.Next {
		if child.Child != nil {
			nodes = append(nodes, blockEntries(child)...)
		} else {
			nodes = append(nodes, child)
		}
		if !child.Arg.Optional {
			break
		}
	}
	return nodes
}

func name(s string) (name string) {
	switch s {
	case "~":
		return "Almost"
	case "=":
		return "Exact"
	case "$":
		return "Last"
	case "\"\"":
		return "Empty"
	}
	for _, n := range strings.Split(strings.NewReplacer("-", " ", "_", " ", ":", " ", "/", " ", ".", " ", "*", "all").Replace(s), " ") {
		name += ucFirst(strings.ToLower(n))
	}
	return name
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func lcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// Code generated DO NOT EDIT

package cmds

import "strconv"

type GraphConfigGet Completed

func (b Builder) GraphConfigGet() (c GraphConfigGet) {
	c.cs = append(c.cs, "GRAPH.CONFIG", "GET")
	return c
}

func (c GraphConfigGet) Name(name interface{}) GraphConfigGetName {
	c.cs = append(c.cs, name)
	return (GraphConfigGetName)(c)
}

type GraphConfigGetName Completed

func (c GraphConfigGetName) GetArgs() []interface{} {
	return c.cs
}

type GraphConfigSet Completed

func (b Builder) GraphConfigSet() (c GraphConfigSet) {
	c.cs = append(c.cs, "GRAPH.CONFIG", "SET")
	return c
}

func (c GraphConfigSet) Name(name interface{}) GraphConfigSetName {
	c.cs = append(c.cs, name)
	return (GraphConfigSetName)(c)
}

type GraphConfigSetName Completed

func (c GraphConfigSetName) Value(value interface{}) GraphConfigSetValue {
	c.cs = append(c.cs, value)
	return (GraphConfigSetValue)(c)
}

type GraphConfigSetValue Completed

func (c GraphConfigSetValue) GetArgs() []interface{} {
	return c.cs
}

type GraphConstraintCreate Completed

func (b Builder) GraphConstraintCreate() (c GraphConstraintCreate) {
	c.cs = append(c.cs, "GRAPH.CONSTRAINT", "CREATE")
	return c
}

func (c GraphConstraintCreate) GetArgs() []interface{} {
	return c.cs
}

type GraphConstraintDrop Completed

func (b Builder) GraphConstraintDrop() (c GraphConstraintDrop) {
	c.cs = append(c.cs, "GRAPH.CONSTRAINT", "DROP")
	return c
}

func (c GraphConstraintDrop) GetArgs() []interface{} {
	return c.cs
}

type GraphDelete Completed

func (b Builder) GraphDelete() (c GraphDelete) {
	c.cs = append(c.cs, "GRAPH.DELETE")
	return c
}

func (c GraphDelete) Graph(graph interface{}) GraphDeleteGraph {
	c.cs = append(c.cs, graph)
	return (GraphDeleteGraph)(c)
}

type GraphDeleteGraph Completed

func (c GraphDeleteGraph) GetArgs() []interface{} {
	return c.cs
}

type GraphExplain Completed

func (b Builder) GraphExplain() (c GraphExplain) {
	c.cs = append(c.cs, "GRAPH.EXPLAIN")
	return c
}

func (c GraphExplain) Graph(graph interface{}) GraphExplainGraph {
	c.cs = append(c.cs, graph)
	return (GraphExplainGraph)(c)
}

type GraphExplainGraph Completed

func (c GraphExplainGraph) Query(query interface{}) GraphExplainQuery {
	c.cs = append(c.cs, query)
	return (GraphExplainQuery)(c)
}

type GraphExplainQuery Completed

func (c GraphExplainQuery) GetArgs() []interface{} {
	return c.cs
}

type GraphList Completed

func (b Builder) GraphList() (c GraphList) {
	c.cs = append(c.cs, "GRAPH.LIST")
	return c
}

func (c GraphList) GetArgs() []interface{} {
	return c.cs
}

type GraphProfile Completed

func (b Builder) GraphProfile() (c GraphProfile) {
	c.cs = append(c.cs, "GRAPH.PROFILE")
	return c
}

func (c GraphProfile) Graph(graph interface{}) GraphProfileGraph {
	c.cs = append(c.cs, graph)
	return (GraphProfileGraph)(c)
}

type GraphProfileGraph Completed

func (c GraphProfileGraph) Query(query interface{}) GraphProfileQuery {
	c.cs = append(c.cs, query)
	return (GraphProfileQuery)(c)
}

type GraphProfileQuery Completed

func (c GraphProfileQuery) Timeout(timeout int64) GraphProfileTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (GraphProfileTimeout)(c)
}

func (c GraphProfileQuery) GetArgs() []interface{} {
	return c.cs
}

type GraphProfileTimeout Completed

func (c GraphProfileTimeout) GetArgs() []interface{} {
	return c.cs
}

type GraphQuery Completed

func (b Builder) GraphQuery() (c GraphQuery) {
	c.cs = append(c.cs, "GRAPH.QUERY")
	return c
}

func (c GraphQuery) Graph(graph interface{}) GraphQueryGraph {
	c.cs = append(c.cs, graph)
	return (GraphQueryGraph)(c)
}

type GraphQueryGraph Completed

func (c GraphQueryGraph) Query(query interface{}) GraphQueryQuery {
	c.cs = append(c.cs, query)
	return (GraphQueryQuery)(c)
}

type GraphQueryQuery Completed

func (c GraphQueryQuery) Timeout(timeout int64) GraphQueryTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (GraphQueryTimeout)(c)
}

func (c GraphQueryQuery) GetArgs() []interface{} {
	return c.cs
}

type GraphQueryTimeout Completed

func (c GraphQueryTimeout) GetArgs() []interface{} {
	return c.cs
}

type GraphRoQuery Completed

func (b Builder) GraphRoQuery() (c GraphRoQuery) {
	c.cs = append(c.cs, "GRAPH.RO_QUERY")
	return c
}

func (c GraphRoQuery) Graph(graph interface{}) GraphRoQueryGraph {
	c.cs = append(c.cs, graph)
	return (GraphRoQueryGraph)(c)
}

type GraphRoQueryGraph Completed

func (c GraphRoQueryGraph) Query(query interface{}) GraphRoQueryQuery {
	c.cs = append(c.cs, query)
	return (GraphRoQueryQuery)(c)
}

type GraphRoQueryQuery Completed

func (c GraphRoQueryQuery) Timeout(timeout int64) GraphRoQueryTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (GraphRoQueryTimeout)(c)
}

func (c GraphRoQueryQuery) GetArgs() []interface{} {
	return c.cs
}

type GraphRoQueryTimeout Completed

func (c GraphRoQueryTimeout) GetArgs() []interface{} {
	return c.cs
}

type GraphSlowlog Completed

func (b Builder) GraphSlowlog() (c GraphSlowlog) {
	c.cs = append(c.cs, "GRAPH.SLOWLOG")
	return c
}

func (c GraphSlowlog) Graph(graph interface{}) GraphSlowlogGraph {
	c.cs = append(c.cs, graph)
	return (GraphSlowlogGraph)(c)
}

type GraphSlowlogGraph Completed

func (c GraphSlowlogGraph) GetArgs() []interface{} {
	return c.cs
}

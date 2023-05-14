package cmds

import (
	"sync"
)

const ErrBuiltTwice = "a command should not be built twice"
const ErrUnfinished = "a command should be finished by calling Build() or Cache()"

// TODOElena This pool can probably be removed, cause we have pooling handled in go-redis
var pool = &sync.Pool{New: func() any {
	return &CommandSlice{s: make([]string, 0, 2), l: -1}
}}

// CommandSlice is the command container managed by the sync.Pool
type CommandSlice struct {
	s []string
	l int32
	r int32
}

func (cs *CommandSlice) Build() {
	if cs.l != -1 {
		panic(ErrBuiltTwice)
	}
	cs.l = int32(len(cs.s))
}

func (cs *CommandSlice) Verify() {
	if cs.l != int32(len(cs.s)) {
		panic(ErrUnfinished)
	}
}

func newCommandSlice(s []string) *CommandSlice {
	return &CommandSlice{s: s, l: int32(len(s))}
}

// NewBuilder creates a Builder and initializes the internal sync.Pool
func NewBuilder(initSlot uint16) *Builder {
	return &Builder{ks: initSlot}
}

// Builder builds commands by reusing CommandSlice from the sync.Pool
type Builder struct {
	ks uint16
}

func get() *CommandSlice {
	return pool.Get().(*CommandSlice)
}

func Put(cs *CommandSlice) {
	cs.s = cs.s[:0]
	cs.l = -1
	cs.r = 0
	pool.Put(cs)
}

// PutCompleted recycles the Completed
func PutCompleted(c Completed) {
	if c.cs.r == 0 {
		Put(c.cs)
	}
}

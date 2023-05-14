package cmds

const (
	optInTag = uint16(1 << 15)
	blockTag = uint16(1 << 14)
	readonly = uint16(1 << 13)
	noRetTag = uint16(1<<12) | readonly // make noRetTag can also be retried
	mtGetTag = uint16(1<<11) | readonly // make mtGetTag can also be retried
	scrRoTag = uint16(1<<10) | readonly // make scrRoTag can also be retried
	// InitSlot indicates that the command be sent to any redis node in cluster
	InitSlot = uint16(1 << 14)
	// NoSlot indicates that the command has no key slot specified
	NoSlot = uint16(1 << 15)
)

// Completed represents a completed Redis command, should be created by the Build() of command builder.
type Completed struct {
	cs *CommandSlice
	cf uint16
	ks uint16
}

// NewCompleted creates an arbitrary Completed command.
func NewCompleted(ss []string) Completed {
	return Completed{cs: newCommandSlice(ss)}
}

// NewBlockingCompleted creates an arbitrary blocking Completed command.
func NewBlockingCompleted(ss []string) Completed {
	return Completed{cs: newCommandSlice(ss), cf: blockTag}
}

// NewReadOnlyCompleted creates an arbitrary readonly Completed command.
func NewReadOnlyCompleted(ss []string) Completed {
	return Completed{cs: newCommandSlice(ss), cf: readonly}
}

// NewMGetCompleted creates an arbitrary readonly Completed command.
func NewMGetCompleted(ss []string) Completed {
	return Completed{cs: newCommandSlice(ss), cf: mtGetTag}
}

// Pin prevents a Completed to be recycled
func (c Completed) Pin() Completed {
	c.cs.r = 1
	return c
}

// IsEmpty checks if it is an empty command.
func (c *Completed) IsEmpty() bool {
	return c.cs == nil || len(c.cs.s) == 0
}

// IsBlock checks if it is blocking command which needs to be process by dedicated connection.
func (c *Completed) IsBlock() bool {
	return c.cf&blockTag == blockTag
}

// NoReply checks if it is one of the SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE or PUNSUBSCRIBE commands.
func (c *Completed) NoReply() bool {
	return c.cf&noRetTag == noRetTag
}

// IsReadOnly checks if it is readonly command and can be retried when network error.
func (c *Completed) IsReadOnly() bool {
	return c.cf&readonly == readonly
}

// IsWrite checks if it is not readonly command.
func (c *Completed) IsWrite() bool {
	return !c.IsReadOnly()
}

// GetArgs returns the commands as []string.
// Note that the returned []string should not be modified
// and should not be read after passing into the Client interface, because it will be recycled.
func (c *Completed) GetArgs() []string {
	return c.cs.s
}

// Slot returns the command key slot
func (c *Completed) Slot() uint16 {
	return c.ks
}

// NewMultiCompleted creates multiple arbitrary Completed commands.
func NewMultiCompleted(cs [][]string) []Completed {
	ret := make([]Completed, len(cs))
	for i, c := range cs {
		ret[i] = NewCompleted(c)
	}
	return ret
}

func check(prev, new uint16) uint16 {
	if prev == InitSlot || prev == new {
		return new
	}
	panic(multiKeySlotErr)
}

const multiKeySlotErr = "multi key command with different key slots are not allowed"
const multiKeyCacheErr = "client side caching for scripting only supports numkeys=1"

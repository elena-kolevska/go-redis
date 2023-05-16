// Code generated DO NOT EDIT

package cmds

import "testing"

func topk0(s Builder) {
	s.TopkAdd().Key("1").Items("1").Items("1").GetArgs()
	s.TopkCount().Key("1").Item("1").Item("1").GetArgs()
	s.TopkIncrby().Key("1").Item("1").Increment(1).Item("1").Increment(1).GetArgs()
	s.TopkInfo().Key("1").GetArgs()
	s.TopkList().Key("1").Withcount().GetArgs()
	s.TopkList().Key("1").GetArgs()
	s.TopkQuery().Key("1").Item("1").Item("1").GetArgs()
	s.TopkReserve().Key("1").Topk(1).Width(1).Depth(1).Decay(1).GetArgs()
	s.TopkReserve().Key("1").Topk(1).GetArgs()
}

func TestCommand_topk(t *testing.T) {
	t.Run("0", func(t *testing.T) { topk0(Builder{}) })
}

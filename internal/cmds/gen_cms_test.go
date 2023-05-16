// Code generated DO NOT EDIT

package cmds

import "testing"

func cms0(s Builder) {
	s.CmsIncrby().Key("1").Item("1").Increment(1).Item("1").Increment(1).GetArgs()
	s.CmsInfo().Key("1").GetArgs()
	s.CmsInitbydim().Key("1").Width(1).Depth(1).GetArgs()
	s.CmsInitbyprob().Key("1").Error(1).Probability(1).GetArgs()
	s.CmsMerge().Destination("1").Numkeys(1).Source("1").Source("1").Weights().Weight(1).Weight(1).GetArgs()
	s.CmsMerge().Destination("1").Numkeys(1).Source("1").Source("1").GetArgs()
	s.CmsQuery().Key("1").Item("1").Item("1").GetArgs()
}

func TestCommand_cms(t *testing.T) {
	t.Run("0", func(t *testing.T) { cms0(Builder{}) })
}

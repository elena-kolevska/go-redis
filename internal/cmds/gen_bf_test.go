// Code generated DO NOT EDIT

package cmds

import "testing"

func bf0(s Builder) {
	s.BfAdd().Key("1").Item("1").GetArgs()
	s.BfCard().Key("1").GetArgs()
	s.BfExists().Key("1").Item("1").GetArgs()
	s.BfInfo().Key("1").Capacity().GetArgs()
	s.BfInfo().Key("1").Size().GetArgs()
	s.BfInfo().Key("1").Filters().GetArgs()
	s.BfInfo().Key("1").Items().GetArgs()
	s.BfInfo().Key("1").Expansion().GetArgs()
	s.BfInfo().Key("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Nocreate().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Error(1).Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Capacity(1).Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Error(1).Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Nocreate().Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Nonscaling().Items().Item("1").Item("1").GetArgs()
	s.BfInsert().Key("1").Items().Item("1").Item("1").GetArgs()
	s.BfLoadchunk().Key("1").Iterator(1).Data("1").GetArgs()
	s.BfMadd().Key("1").Item("1").Item("1").GetArgs()
	s.BfMexists().Key("1").Item("1").Item("1").GetArgs()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Expansion(1).Nonscaling().GetArgs()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Expansion(1).GetArgs()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Nonscaling().GetArgs()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).GetArgs()
	s.BfScandump().Key("1").Iterator(1).GetArgs()
}

func TestCommand_bf(t *testing.T) {
	t.Run("0", func(t *testing.T) { bf0(Builder{}) })
}

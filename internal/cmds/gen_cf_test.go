// Code generated DO NOT EDIT

package cmds

import "testing"

func cf0(s Builder) {
	s.CfAdd().Key("1").Item("1").GetArgs()
	s.CfAddnx().Key("1").Item("1").GetArgs()
	s.CfCount().Key("1").Item("1").GetArgs()
	s.CfDel().Key("1").Item("1").GetArgs()
	s.CfExists().Key("1").Item("1").GetArgs()
	s.CfInfo().Key("1").GetArgs()
	s.CfInsert().Key("1").Capacity(1).Nocreate().Items().Item("1").Item("1").GetArgs()
	s.CfInsert().Key("1").Capacity(1).Items().Item("1").Item("1").GetArgs()
	s.CfInsert().Key("1").Nocreate().Items().Item("1").Item("1").GetArgs()
	s.CfInsert().Key("1").Items().Item("1").Item("1").GetArgs()
	s.CfInsertnx().Key("1").Capacity(1).Nocreate().Items().Item("1").Item("1").GetArgs()
	s.CfInsertnx().Key("1").Capacity(1).Items().Item("1").Item("1").GetArgs()
	s.CfInsertnx().Key("1").Nocreate().Items().Item("1").Item("1").GetArgs()
	s.CfInsertnx().Key("1").Items().Item("1").Item("1").GetArgs()
	s.CfLoadchunk().Key("1").Iterator(1).Data("1").GetArgs()
	s.CfMexists().Key("1").Item("1").Item("1").GetArgs()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Maxiterations(1).Expansion(1).GetArgs()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Maxiterations(1).GetArgs()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Expansion(1).GetArgs()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).GetArgs()
	s.CfReserve().Key("1").Capacity(1).Maxiterations(1).GetArgs()
	s.CfReserve().Key("1").Capacity(1).Expansion(1).GetArgs()
	s.CfReserve().Key("1").Capacity(1).GetArgs()
	s.CfScandump().Key("1").Iterator(1).GetArgs()
}

func TestCommand_cf(t *testing.T) {
	t.Run("0", func(t *testing.T) { cf0(Builder{}) })
}

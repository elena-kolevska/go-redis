// Code generated DO NOT EDIT

package cmds

import "testing"

func graph0(s Builder) {
	s.GraphConfigGet().Name("1").GetArgs()
	s.GraphConfigSet().Name("1").Value("1").GetArgs()
	s.GraphConstraintCreate().GetArgs()
	s.GraphConstraintDrop().GetArgs()
	s.GraphDelete().Graph("1").GetArgs()
	s.GraphExplain().Graph("1").Query("1").GetArgs()
	s.GraphList().GetArgs()
	s.GraphProfile().Graph("1").Query("1").Timeout(1).GetArgs()
	s.GraphProfile().Graph("1").Query("1").GetArgs()
	s.GraphQuery().Graph("1").Query("1").Timeout(1).GetArgs()
	s.GraphQuery().Graph("1").Query("1").GetArgs()
	s.GraphRoQuery().Graph("1").Query("1").Timeout(1).GetArgs()
	s.GraphRoQuery().Graph("1").Query("1").GetArgs()
	s.GraphSlowlog().Graph("1").GetArgs()
}

func TestCommand_graph(t *testing.T) {
	t.Run("0", func(t *testing.T) { graph0(Builder{}) })
}

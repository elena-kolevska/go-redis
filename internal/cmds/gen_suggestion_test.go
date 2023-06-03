// Code generated DO NOT EDIT

package cmds

import "testing"

func suggestion0(s Builder) {
	s.FtSugadd().Key("1").String("1").Score(1).Incr().Payload("1").GetArgs()
	s.FtSugadd().Key("1").String("1").Score(1).Incr().GetArgs()
	s.FtSugadd().Key("1").String("1").Score(1).Payload("1").GetArgs()
	s.FtSugadd().Key("1").String("1").Score(1).GetArgs()
	s.FtSugdel().Key("1").String("1").GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Withpayloads().Max(1).GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Withpayloads().GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Max(1).GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withpayloads().GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Max(1).GetArgs()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().GetArgs()
	s.FtSugget().Key("1").Prefix("1").Withscores().GetArgs()
	s.FtSugget().Key("1").Prefix("1").Withpayloads().GetArgs()
	s.FtSugget().Key("1").Prefix("1").Max(1).GetArgs()
	s.FtSugget().Key("1").Prefix("1").GetArgs()
	s.FtSuglen().Key("1").GetArgs()
}

func TestCommand_suggestion(t *testing.T) {
	t.Run("0", func(t *testing.T) { suggestion0(Builder{}) })
}

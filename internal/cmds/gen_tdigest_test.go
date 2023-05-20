// Code generated DO NOT EDIT

package cmds

import "testing"

func tdigest0(s Builder) {
	s.TdigestAdd().Key("1").Value(1).Value(1).GetArgs()
	s.TdigestByrank().Key("1").Rank(1).Rank(1).GetArgs()
	s.TdigestByrevrank().Key("1").ReverseRank(1).ReverseRank(1).GetArgs()
	s.TdigestCdf().Key("1").Value(1).Value(1).GetArgs()
	s.TdigestCreate().Key("1").Compression(1).GetArgs()
	s.TdigestCreate().Key("1").GetArgs()
	s.TdigestInfo().Key("1").GetArgs()
	s.TdigestMax().Key("1").GetArgs()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Compression(1).Override().GetArgs()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Compression(1).GetArgs()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Override().GetArgs()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").GetArgs()
	s.TdigestMin().Key("1").GetArgs()
	s.TdigestQuantile().Key("1").Quantile(1).Quantile(1).GetArgs()
	s.TdigestRank().Key("1").Value(1).Value(1).GetArgs()
	s.TdigestReset().Key("1").GetArgs()
	s.TdigestRevrank().Key("1").Value(1).Value(1).GetArgs()
	s.TdigestTrimmedMean().Key("1").LowCutQuantile(1).HighCutQuantile(1).GetArgs()
}

func TestCommand_tdigest(t *testing.T) {
	t.Run("0", func(t *testing.T) { tdigest0(Builder{}) })
}

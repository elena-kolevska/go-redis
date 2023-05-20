// Code generated DO NOT EDIT

package cmds

import "strconv"

type TdigestAdd Completed

func (b Builder) TdigestAdd() (c TdigestAdd) {
	c.cs = append(c.cs, "TDIGEST.ADD")
	return c
}

func (c TdigestAdd) Key(key interface{}) TdigestAddKey {
	c.cs = append(c.cs, key)
	return (TdigestAddKey)(c)
}

type TdigestAddKey Completed

func (c TdigestAddKey) Value(value float64) TdigestAddValuesValue {
	c.cs = append(c.cs, strconv.FormatFloat(value, 'f', -1, 64))
	return (TdigestAddValuesValue)(c)
}

type TdigestAddValuesValue Completed

func (c TdigestAddValuesValue) Value(value float64) TdigestAddValuesValue {
	c.cs = append(c.cs, strconv.FormatFloat(value, 'f', -1, 64))
	return c
}

func (c TdigestAddValuesValue) GetArgs() []interface{} {
	return c.cs
}

type TdigestByrank Completed

func (b Builder) TdigestByrank() (c TdigestByrank) {
	c.cs = append(c.cs, "TDIGEST.BYRANK")
	return c
}

func (c TdigestByrank) Key(key interface{}) TdigestByrankKey {
	c.cs = append(c.cs, key)
	return (TdigestByrankKey)(c)
}

type TdigestByrankKey Completed

func (c TdigestByrankKey) Rank(rank ...float64) TdigestByrankRank {
	for _, n := range rank {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (TdigestByrankRank)(c)
}

type TdigestByrankRank Completed

func (c TdigestByrankRank) Rank(rank ...float64) TdigestByrankRank {
	for _, n := range rank {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c TdigestByrankRank) GetArgs() []interface{} {
	return c.cs
}

type TdigestByrevrank Completed

func (b Builder) TdigestByrevrank() (c TdigestByrevrank) {
	c.cs = append(c.cs, "TDIGEST.BYREVRANK")
	return c
}

func (c TdigestByrevrank) Key(key interface{}) TdigestByrevrankKey {
	c.cs = append(c.cs, key)
	return (TdigestByrevrankKey)(c)
}

type TdigestByrevrankKey Completed

func (c TdigestByrevrankKey) ReverseRank(reverseRank ...float64) TdigestByrevrankReverseRank {
	for _, n := range reverseRank {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (TdigestByrevrankReverseRank)(c)
}

type TdigestByrevrankReverseRank Completed

func (c TdigestByrevrankReverseRank) ReverseRank(reverseRank ...float64) TdigestByrevrankReverseRank {
	for _, n := range reverseRank {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c TdigestByrevrankReverseRank) GetArgs() []interface{} {
	return c.cs
}

type TdigestCdf Completed

func (b Builder) TdigestCdf() (c TdigestCdf) {
	c.cs = append(c.cs, "TDIGEST.CDF")
	return c
}

func (c TdigestCdf) Key(key interface{}) TdigestCdfKey {
	c.cs = append(c.cs, key)
	return (TdigestCdfKey)(c)
}

type TdigestCdfKey Completed

func (c TdigestCdfKey) Value(value ...float64) TdigestCdfValue {
	for _, n := range value {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (TdigestCdfValue)(c)
}

type TdigestCdfValue Completed

func (c TdigestCdfValue) Value(value ...float64) TdigestCdfValue {
	for _, n := range value {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c TdigestCdfValue) GetArgs() []interface{} {
	return c.cs
}

type TdigestCreate Completed

func (b Builder) TdigestCreate() (c TdigestCreate) {
	c.cs = append(c.cs, "TDIGEST.CREATE")
	return c
}

func (c TdigestCreate) Key(key interface{}) TdigestCreateKey {
	c.cs = append(c.cs, key)
	return (TdigestCreateKey)(c)
}

type TdigestCreateCompression Completed

func (c TdigestCreateCompression) GetArgs() []interface{} {
	return c.cs
}

type TdigestCreateKey Completed

func (c TdigestCreateKey) Compression(compression int64) TdigestCreateCompression {
	c.cs = append(c.cs, "COMPRESSION", strconv.FormatInt(compression, 10))
	return (TdigestCreateCompression)(c)
}

func (c TdigestCreateKey) GetArgs() []interface{} {
	return c.cs
}

type TdigestInfo Completed

func (b Builder) TdigestInfo() (c TdigestInfo) {
	c.cs = append(c.cs, "TDIGEST.INFO")
	return c
}

func (c TdigestInfo) Key(key interface{}) TdigestInfoKey {
	c.cs = append(c.cs, key)
	return (TdigestInfoKey)(c)
}

type TdigestInfoKey Completed

func (c TdigestInfoKey) GetArgs() []interface{} {
	return c.cs
}

type TdigestMax Completed

func (b Builder) TdigestMax() (c TdigestMax) {
	c.cs = append(c.cs, "TDIGEST.MAX")
	return c
}

func (c TdigestMax) Key(key interface{}) TdigestMaxKey {
	c.cs = append(c.cs, key)
	return (TdigestMaxKey)(c)
}

type TdigestMaxKey Completed

func (c TdigestMaxKey) GetArgs() []interface{} {
	return c.cs
}

type TdigestMerge Completed

func (b Builder) TdigestMerge() (c TdigestMerge) {
	c.cs = append(c.cs, "TDIGEST.MERGE")
	return c
}

func (c TdigestMerge) DestinationKey(destinationKey interface{}) TdigestMergeDestinationKey {
	c.cs = append(c.cs, destinationKey)
	return (TdigestMergeDestinationKey)(c)
}

type TdigestMergeConfigCompression Completed

func (c TdigestMergeConfigCompression) Override() TdigestMergeOverride {
	c.cs = append(c.cs, "OVERRIDE")
	return (TdigestMergeOverride)(c)
}

func (c TdigestMergeConfigCompression) GetArgs() []interface{} {
	return c.cs
}

type TdigestMergeDestinationKey Completed

func (c TdigestMergeDestinationKey) Numkeys(numkeys int64) TdigestMergeNumkeys {
	c.cs = append(c.cs, strconv.FormatInt(numkeys, 10))
	return (TdigestMergeNumkeys)(c)
}

type TdigestMergeNumkeys Completed

func (c TdigestMergeNumkeys) SourceKey(sourceKey ...interface{}) TdigestMergeSourceKey {
	c.cs = append(c.cs, sourceKey...)
	return (TdigestMergeSourceKey)(c)
}

type TdigestMergeOverride Completed

func (c TdigestMergeOverride) GetArgs() []interface{} {
	return c.cs
}

type TdigestMergeSourceKey Completed

func (c TdigestMergeSourceKey) SourceKey(sourceKey ...interface{}) TdigestMergeSourceKey {
	c.cs = append(c.cs, sourceKey...)
	return c
}

func (c TdigestMergeSourceKey) Compression(compression int64) TdigestMergeConfigCompression {
	c.cs = append(c.cs, "COMPRESSION", strconv.FormatInt(compression, 10))
	return (TdigestMergeConfigCompression)(c)
}

func (c TdigestMergeSourceKey) Override() TdigestMergeOverride {
	c.cs = append(c.cs, "OVERRIDE")
	return (TdigestMergeOverride)(c)
}

func (c TdigestMergeSourceKey) GetArgs() []interface{} {
	return c.cs
}

type TdigestMin Completed

func (b Builder) TdigestMin() (c TdigestMin) {
	c.cs = append(c.cs, "TDIGEST.MIN")
	return c
}

func (c TdigestMin) Key(key interface{}) TdigestMinKey {
	c.cs = append(c.cs, key)
	return (TdigestMinKey)(c)
}

type TdigestMinKey Completed

func (c TdigestMinKey) GetArgs() []interface{} {
	return c.cs
}

type TdigestQuantile Completed

func (b Builder) TdigestQuantile() (c TdigestQuantile) {
	c.cs = append(c.cs, "TDIGEST.QUANTILE")
	return c
}

func (c TdigestQuantile) Key(key interface{}) TdigestQuantileKey {
	c.cs = append(c.cs, key)
	return (TdigestQuantileKey)(c)
}

type TdigestQuantileKey Completed

func (c TdigestQuantileKey) Quantile(quantile ...float64) TdigestQuantileQuantile {
	for _, n := range quantile {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (TdigestQuantileQuantile)(c)
}

type TdigestQuantileQuantile Completed

func (c TdigestQuantileQuantile) Quantile(quantile ...float64) TdigestQuantileQuantile {
	for _, n := range quantile {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c TdigestQuantileQuantile) GetArgs() []interface{} {
	return c.cs
}

type TdigestRank Completed

func (b Builder) TdigestRank() (c TdigestRank) {
	c.cs = append(c.cs, "TDIGEST.RANK")
	return c
}

func (c TdigestRank) Key(key interface{}) TdigestRankKey {
	c.cs = append(c.cs, key)
	return (TdigestRankKey)(c)
}

type TdigestRankKey Completed

func (c TdigestRankKey) Value(value ...float64) TdigestRankValue {
	for _, n := range value {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (TdigestRankValue)(c)
}

type TdigestRankValue Completed

func (c TdigestRankValue) Value(value ...float64) TdigestRankValue {
	for _, n := range value {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c TdigestRankValue) GetArgs() []interface{} {
	return c.cs
}

type TdigestReset Completed

func (b Builder) TdigestReset() (c TdigestReset) {
	c.cs = append(c.cs, "TDIGEST.RESET")
	return c
}

func (c TdigestReset) Key(key interface{}) TdigestResetKey {
	c.cs = append(c.cs, key)
	return (TdigestResetKey)(c)
}

type TdigestResetKey Completed

func (c TdigestResetKey) GetArgs() []interface{} {
	return c.cs
}

type TdigestRevrank Completed

func (b Builder) TdigestRevrank() (c TdigestRevrank) {
	c.cs = append(c.cs, "TDIGEST.REVRANK")
	return c
}

func (c TdigestRevrank) Key(key interface{}) TdigestRevrankKey {
	c.cs = append(c.cs, key)
	return (TdigestRevrankKey)(c)
}

type TdigestRevrankKey Completed

func (c TdigestRevrankKey) Value(value ...float64) TdigestRevrankValue {
	for _, n := range value {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (TdigestRevrankValue)(c)
}

type TdigestRevrankValue Completed

func (c TdigestRevrankValue) Value(value ...float64) TdigestRevrankValue {
	for _, n := range value {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c TdigestRevrankValue) GetArgs() []interface{} {
	return c.cs
}

type TdigestTrimmedMean Completed

func (b Builder) TdigestTrimmedMean() (c TdigestTrimmedMean) {
	c.cs = append(c.cs, "TDIGEST.TRIMMED_MEAN")
	return c
}

func (c TdigestTrimmedMean) Key(key interface{}) TdigestTrimmedMeanKey {
	c.cs = append(c.cs, key)
	return (TdigestTrimmedMeanKey)(c)
}

type TdigestTrimmedMeanHighCutQuantile Completed

func (c TdigestTrimmedMeanHighCutQuantile) GetArgs() []interface{} {
	return c.cs
}

type TdigestTrimmedMeanKey Completed

func (c TdigestTrimmedMeanKey) LowCutQuantile(lowCutQuantile float64) TdigestTrimmedMeanLowCutQuantile {
	c.cs = append(c.cs, strconv.FormatFloat(lowCutQuantile, 'f', -1, 64))
	return (TdigestTrimmedMeanLowCutQuantile)(c)
}

type TdigestTrimmedMeanLowCutQuantile Completed

func (c TdigestTrimmedMeanLowCutQuantile) HighCutQuantile(highCutQuantile float64) TdigestTrimmedMeanHighCutQuantile {
	c.cs = append(c.cs, strconv.FormatFloat(highCutQuantile, 'f', -1, 64))
	return (TdigestTrimmedMeanHighCutQuantile)(c)
}

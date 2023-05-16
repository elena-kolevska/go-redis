// Code generated DO NOT EDIT

package cmds

import "strconv"

type CmsIncrby Completed

func (b Builder) CmsIncrby() (c CmsIncrby) {
	c.cs = append(c.cs, "CMS.INCRBY")
	return c
}

func (c CmsIncrby) Key(key interface{}) CmsIncrbyKey {
	c.cs = append(c.cs, key)
	return (CmsIncrbyKey)(c)
}

type CmsIncrbyItemsIncrement Completed

func (c CmsIncrbyItemsIncrement) Item(item interface{}) CmsIncrbyItemsItem {
	c.cs = append(c.cs, item)
	return (CmsIncrbyItemsItem)(c)
}

func (c CmsIncrbyItemsIncrement) GetArgs() []interface{} {
	return c.cs
}

type CmsIncrbyItemsItem Completed

func (c CmsIncrbyItemsItem) Increment(increment int64) CmsIncrbyItemsIncrement {
	c.cs = append(c.cs, strconv.FormatInt(increment, 10))
	return (CmsIncrbyItemsIncrement)(c)
}

type CmsIncrbyKey Completed

func (c CmsIncrbyKey) Item(item interface{}) CmsIncrbyItemsItem {
	c.cs = append(c.cs, item)
	return (CmsIncrbyItemsItem)(c)
}

type CmsInfo Completed

func (b Builder) CmsInfo() (c CmsInfo) {
	c.cs = append(c.cs, "CMS.INFO")
	return c
}

func (c CmsInfo) Key(key interface{}) CmsInfoKey {
	c.cs = append(c.cs, key)
	return (CmsInfoKey)(c)
}

type CmsInfoKey Completed

func (c CmsInfoKey) GetArgs() []interface{} {
	return c.cs
}

type CmsInitbydim Completed

func (b Builder) CmsInitbydim() (c CmsInitbydim) {
	c.cs = append(c.cs, "CMS.INITBYDIM")
	return c
}

func (c CmsInitbydim) Key(key interface{}) CmsInitbydimKey {
	c.cs = append(c.cs, key)
	return (CmsInitbydimKey)(c)
}

type CmsInitbydimDepth Completed

func (c CmsInitbydimDepth) GetArgs() []interface{} {
	return c.cs
}

type CmsInitbydimKey Completed

func (c CmsInitbydimKey) Width(width int64) CmsInitbydimWidth {
	c.cs = append(c.cs, strconv.FormatInt(width, 10))
	return (CmsInitbydimWidth)(c)
}

type CmsInitbydimWidth Completed

func (c CmsInitbydimWidth) Depth(depth int64) CmsInitbydimDepth {
	c.cs = append(c.cs, strconv.FormatInt(depth, 10))
	return (CmsInitbydimDepth)(c)
}

type CmsInitbyprob Completed

func (b Builder) CmsInitbyprob() (c CmsInitbyprob) {
	c.cs = append(c.cs, "CMS.INITBYPROB")
	return c
}

func (c CmsInitbyprob) Key(key interface{}) CmsInitbyprobKey {
	c.cs = append(c.cs, key)
	return (CmsInitbyprobKey)(c)
}

type CmsInitbyprobError Completed

func (c CmsInitbyprobError) Probability(probability float64) CmsInitbyprobProbability {
	c.cs = append(c.cs, strconv.FormatFloat(probability, 'f', -1, 64))
	return (CmsInitbyprobProbability)(c)
}

type CmsInitbyprobKey Completed

func (c CmsInitbyprobKey) Error(error float64) CmsInitbyprobError {
	c.cs = append(c.cs, strconv.FormatFloat(error, 'f', -1, 64))
	return (CmsInitbyprobError)(c)
}

type CmsInitbyprobProbability Completed

func (c CmsInitbyprobProbability) GetArgs() []interface{} {
	return c.cs
}

type CmsMerge Completed

func (b Builder) CmsMerge() (c CmsMerge) {
	c.cs = append(c.cs, "CMS.MERGE")
	return c
}

func (c CmsMerge) Destination(destination interface{}) CmsMergeDestination {
	c.cs = append(c.cs, destination)
	return (CmsMergeDestination)(c)
}

type CmsMergeDestination Completed

func (c CmsMergeDestination) Numkeys(numkeys int64) CmsMergeNumkeys {
	c.cs = append(c.cs, strconv.FormatInt(numkeys, 10))
	return (CmsMergeNumkeys)(c)
}

type CmsMergeNumkeys Completed

func (c CmsMergeNumkeys) Source(source ...interface{}) CmsMergeSource {
	c.cs = append(c.cs, source...)
	return (CmsMergeSource)(c)
}

type CmsMergeSource Completed

func (c CmsMergeSource) Source(source ...interface{}) CmsMergeSource {
	c.cs = append(c.cs, source...)
	return c
}

func (c CmsMergeSource) Weights() CmsMergeWeightWeights {
	c.cs = append(c.cs, "WEIGHTS")
	return (CmsMergeWeightWeights)(c)
}

func (c CmsMergeSource) GetArgs() []interface{} {
	return c.cs
}

type CmsMergeWeightWeight Completed

func (c CmsMergeWeightWeight) Weight(weight ...float64) CmsMergeWeightWeight {
	for _, n := range weight {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return c
}

func (c CmsMergeWeightWeight) GetArgs() []interface{} {
	return c.cs
}

type CmsMergeWeightWeights Completed

func (c CmsMergeWeightWeights) Weight(weight ...float64) CmsMergeWeightWeight {
	for _, n := range weight {
		c.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))
	}
	return (CmsMergeWeightWeight)(c)
}

type CmsQuery Completed

func (b Builder) CmsQuery() (c CmsQuery) {
	c.cs = append(c.cs, "CMS.QUERY")
	return c
}

func (c CmsQuery) Key(key interface{}) CmsQueryKey {
	c.cs = append(c.cs, key)
	return (CmsQueryKey)(c)
}

type CmsQueryItem Completed

func (c CmsQueryItem) Item(item ...interface{}) CmsQueryItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c CmsQueryItem) GetArgs() []interface{} {
	return c.cs
}

type CmsQueryKey Completed

func (c CmsQueryKey) Item(item ...interface{}) CmsQueryItem {
	c.cs = append(c.cs, item...)
	return (CmsQueryItem)(c)
}

// Code generated DO NOT EDIT

package cmds

import "strconv"

type TopkAdd Completed

func (b Builder) TopkAdd() (c TopkAdd) {
	c.cs = append(c.cs, "TOPK.ADD")
	return c
}

func (c TopkAdd) Key(key interface{}) TopkAddKey {
	c.cs = append(c.cs, key)
	return (TopkAddKey)(c)
}

type TopkAddItems Completed

func (c TopkAddItems) Items(items ...interface{}) TopkAddItems {
	c.cs = append(c.cs, items...)
	return c
}

func (c TopkAddItems) GetArgs() []interface{} {
	return c.cs
}

type TopkAddKey Completed

func (c TopkAddKey) Items(items ...interface{}) TopkAddItems {
	c.cs = append(c.cs, items...)
	return (TopkAddItems)(c)
}

type TopkCount Completed

func (b Builder) TopkCount() (c TopkCount) {
	c.cs = append(c.cs, "TOPK.COUNT")
	return c
}

func (c TopkCount) Key(key interface{}) TopkCountKey {
	c.cs = append(c.cs, key)
	return (TopkCountKey)(c)
}

type TopkCountItem Completed

func (c TopkCountItem) Item(item ...interface{}) TopkCountItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c TopkCountItem) GetArgs() []interface{} {
	return c.cs
}

type TopkCountKey Completed

func (c TopkCountKey) Item(item ...interface{}) TopkCountItem {
	c.cs = append(c.cs, item...)
	return (TopkCountItem)(c)
}

type TopkIncrby Completed

func (b Builder) TopkIncrby() (c TopkIncrby) {
	c.cs = append(c.cs, "TOPK.INCRBY")
	return c
}

func (c TopkIncrby) Key(key interface{}) TopkIncrbyKey {
	c.cs = append(c.cs, key)
	return (TopkIncrbyKey)(c)
}

type TopkIncrbyItemsIncrement Completed

func (c TopkIncrbyItemsIncrement) Item(item interface{}) TopkIncrbyItemsItem {
	c.cs = append(c.cs, item)
	return (TopkIncrbyItemsItem)(c)
}

func (c TopkIncrbyItemsIncrement) GetArgs() []interface{} {
	return c.cs
}

type TopkIncrbyItemsItem Completed

func (c TopkIncrbyItemsItem) Increment(increment int64) TopkIncrbyItemsIncrement {
	c.cs = append(c.cs, strconv.FormatInt(increment, 10))
	return (TopkIncrbyItemsIncrement)(c)
}

type TopkIncrbyKey Completed

func (c TopkIncrbyKey) Item(item interface{}) TopkIncrbyItemsItem {
	c.cs = append(c.cs, item)
	return (TopkIncrbyItemsItem)(c)
}

type TopkInfo Completed

func (b Builder) TopkInfo() (c TopkInfo) {
	c.cs = append(c.cs, "TOPK.INFO")
	return c
}

func (c TopkInfo) Key(key interface{}) TopkInfoKey {
	c.cs = append(c.cs, key)
	return (TopkInfoKey)(c)
}

type TopkInfoKey Completed

func (c TopkInfoKey) GetArgs() []interface{} {
	return c.cs
}

type TopkList Completed

func (b Builder) TopkList() (c TopkList) {
	c.cs = append(c.cs, "TOPK.LIST")
	return c
}

func (c TopkList) Key(key interface{}) TopkListKey {
	c.cs = append(c.cs, key)
	return (TopkListKey)(c)
}

type TopkListKey Completed

func (c TopkListKey) Withcount() TopkListWithcount {
	c.cs = append(c.cs, "WITHCOUNT")
	return (TopkListWithcount)(c)
}

func (c TopkListKey) GetArgs() []interface{} {
	return c.cs
}

type TopkListWithcount Completed

func (c TopkListWithcount) GetArgs() []interface{} {
	return c.cs
}

type TopkQuery Completed

func (b Builder) TopkQuery() (c TopkQuery) {
	c.cs = append(c.cs, "TOPK.QUERY")
	return c
}

func (c TopkQuery) Key(key interface{}) TopkQueryKey {
	c.cs = append(c.cs, key)
	return (TopkQueryKey)(c)
}

type TopkQueryItem Completed

func (c TopkQueryItem) Item(item ...interface{}) TopkQueryItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c TopkQueryItem) GetArgs() []interface{} {
	return c.cs
}

type TopkQueryKey Completed

func (c TopkQueryKey) Item(item ...interface{}) TopkQueryItem {
	c.cs = append(c.cs, item...)
	return (TopkQueryItem)(c)
}

type TopkReserve Completed

func (b Builder) TopkReserve() (c TopkReserve) {
	c.cs = append(c.cs, "TOPK.RESERVE")
	return c
}

func (c TopkReserve) Key(key interface{}) TopkReserveKey {
	c.cs = append(c.cs, key)
	return (TopkReserveKey)(c)
}

type TopkReserveKey Completed

func (c TopkReserveKey) Topk(topk int64) TopkReserveTopk {
	c.cs = append(c.cs, strconv.FormatInt(topk, 10))
	return (TopkReserveTopk)(c)
}

type TopkReserveParamsDecay Completed

func (c TopkReserveParamsDecay) GetArgs() []interface{} {
	return c.cs
}

type TopkReserveParamsDepth Completed

func (c TopkReserveParamsDepth) Decay(decay float64) TopkReserveParamsDecay {
	c.cs = append(c.cs, strconv.FormatFloat(decay, 'f', -1, 64))
	return (TopkReserveParamsDecay)(c)
}

type TopkReserveParamsWidth Completed

func (c TopkReserveParamsWidth) Depth(depth int64) TopkReserveParamsDepth {
	c.cs = append(c.cs, strconv.FormatInt(depth, 10))
	return (TopkReserveParamsDepth)(c)
}

type TopkReserveTopk Completed

func (c TopkReserveTopk) Width(width int64) TopkReserveParamsWidth {
	c.cs = append(c.cs, strconv.FormatInt(width, 10))
	return (TopkReserveParamsWidth)(c)
}

func (c TopkReserveTopk) GetArgs() []interface{} {
	return c.cs
}

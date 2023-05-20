// Code generated DO NOT EDIT

package cmds

import "strconv"

type BfAdd Completed

func (b Builder) BfAdd() (c BfAdd) {
	c.cs = append(c.cs, "BF.ADD")
	return c
}

func (c BfAdd) Key(key interface{}) BfAddKey {
	c.cs = append(c.cs, key)
	return (BfAddKey)(c)
}

type BfAddItem Completed

func (c BfAddItem) GetArgs() []interface{} {
	return c.cs
}

type BfAddKey Completed

func (c BfAddKey) Item(item interface{}) BfAddItem {
	c.cs = append(c.cs, item)
	return (BfAddItem)(c)
}

type BfCard Completed

func (b Builder) BfCard() (c BfCard) {
	c.cs = append(c.cs, "BF.CARD")
	return c
}

func (c BfCard) Key(key interface{}) BfCardKey {
	c.cs = append(c.cs, key)
	return (BfCardKey)(c)
}

type BfCardKey Completed

func (c BfCardKey) GetArgs() []interface{} {
	return c.cs
}

type BfExists Completed

func (b Builder) BfExists() (c BfExists) {
	c.cs = append(c.cs, "BF.EXISTS")
	return c
}

func (c BfExists) Key(key interface{}) BfExistsKey {
	c.cs = append(c.cs, key)
	return (BfExistsKey)(c)
}

type BfExistsItem Completed

func (c BfExistsItem) GetArgs() []interface{} {
	return c.cs
}

type BfExistsKey Completed

func (c BfExistsKey) Item(item interface{}) BfExistsItem {
	c.cs = append(c.cs, item)
	return (BfExistsItem)(c)
}

type BfInfo Completed

func (b Builder) BfInfo() (c BfInfo) {
	c.cs = append(c.cs, "BF.INFO")
	return c
}

func (c BfInfo) Key(key interface{}) BfInfoKey {
	c.cs = append(c.cs, key)
	return (BfInfoKey)(c)
}

type BfInfoKey Completed

func (c BfInfoKey) Capacity() BfInfoSingleValueCapacity {
	c.cs = append(c.cs, "CAPACITY")
	return (BfInfoSingleValueCapacity)(c)
}

func (c BfInfoKey) Size() BfInfoSingleValueSize {
	c.cs = append(c.cs, "SIZE")
	return (BfInfoSingleValueSize)(c)
}

func (c BfInfoKey) Filters() BfInfoSingleValueFilters {
	c.cs = append(c.cs, "FILTERS")
	return (BfInfoSingleValueFilters)(c)
}

func (c BfInfoKey) Items() BfInfoSingleValueItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInfoSingleValueItems)(c)
}

func (c BfInfoKey) Expansion() BfInfoSingleValueExpansion {
	c.cs = append(c.cs, "EXPANSION")
	return (BfInfoSingleValueExpansion)(c)
}

func (c BfInfoKey) GetArgs() []interface{} {
	return c.cs
}

type BfInfoSingleValueCapacity Completed

func (c BfInfoSingleValueCapacity) GetArgs() []interface{} {
	return c.cs
}

type BfInfoSingleValueExpansion Completed

func (c BfInfoSingleValueExpansion) GetArgs() []interface{} {
	return c.cs
}

type BfInfoSingleValueFilters Completed

func (c BfInfoSingleValueFilters) GetArgs() []interface{} {
	return c.cs
}

type BfInfoSingleValueItems Completed

func (c BfInfoSingleValueItems) GetArgs() []interface{} {
	return c.cs
}

type BfInfoSingleValueSize Completed

func (c BfInfoSingleValueSize) GetArgs() []interface{} {
	return c.cs
}

type BfInsert Completed

func (b Builder) BfInsert() (c BfInsert) {
	c.cs = append(c.cs, "BF.INSERT")
	return c
}

func (c BfInsert) Key(key interface{}) BfInsertKey {
	c.cs = append(c.cs, key)
	return (BfInsertKey)(c)
}

type BfInsertCapacity Completed

func (c BfInsertCapacity) Error(error float64) BfInsertError {
	c.cs = append(c.cs, "ERROR", strconv.FormatFloat(error, 'f', -1, 64))
	return (BfInsertError)(c)
}

func (c BfInsertCapacity) Expansion(expansion int64) BfInsertExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (BfInsertExpansion)(c)
}

func (c BfInsertCapacity) Nocreate() BfInsertNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (BfInsertNocreate)(c)
}

func (c BfInsertCapacity) Nonscaling() BfInsertNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfInsertNonscaling)(c)
}

func (c BfInsertCapacity) Items() BfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInsertItems)(c)
}

type BfInsertError Completed

func (c BfInsertError) Expansion(expansion int64) BfInsertExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (BfInsertExpansion)(c)
}

func (c BfInsertError) Nocreate() BfInsertNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (BfInsertNocreate)(c)
}

func (c BfInsertError) Nonscaling() BfInsertNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfInsertNonscaling)(c)
}

func (c BfInsertError) Items() BfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInsertItems)(c)
}

type BfInsertExpansion Completed

func (c BfInsertExpansion) Nocreate() BfInsertNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (BfInsertNocreate)(c)
}

func (c BfInsertExpansion) Nonscaling() BfInsertNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfInsertNonscaling)(c)
}

func (c BfInsertExpansion) Items() BfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInsertItems)(c)
}

type BfInsertItem Completed

func (c BfInsertItem) Item(item ...interface{}) BfInsertItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c BfInsertItem) GetArgs() []interface{} {
	return c.cs
}

type BfInsertItems Completed

func (c BfInsertItems) Item(item ...interface{}) BfInsertItem {
	c.cs = append(c.cs, item...)
	return (BfInsertItem)(c)
}

type BfInsertKey Completed

func (c BfInsertKey) Capacity(capacity int64) BfInsertCapacity {
	c.cs = append(c.cs, "CAPACITY", strconv.FormatInt(capacity, 10))
	return (BfInsertCapacity)(c)
}

func (c BfInsertKey) Error(error float64) BfInsertError {
	c.cs = append(c.cs, "ERROR", strconv.FormatFloat(error, 'f', -1, 64))
	return (BfInsertError)(c)
}

func (c BfInsertKey) Expansion(expansion int64) BfInsertExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (BfInsertExpansion)(c)
}

func (c BfInsertKey) Nocreate() BfInsertNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (BfInsertNocreate)(c)
}

func (c BfInsertKey) Nonscaling() BfInsertNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfInsertNonscaling)(c)
}

func (c BfInsertKey) Items() BfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInsertItems)(c)
}

type BfInsertNocreate Completed

func (c BfInsertNocreate) Nonscaling() BfInsertNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfInsertNonscaling)(c)
}

func (c BfInsertNocreate) Items() BfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInsertItems)(c)
}

type BfInsertNonscaling Completed

func (c BfInsertNonscaling) Items() BfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (BfInsertItems)(c)
}

type BfLoadchunk Completed

func (b Builder) BfLoadchunk() (c BfLoadchunk) {
	c.cs = append(c.cs, "BF.LOADCHUNK")
	return c
}

func (c BfLoadchunk) Key(key interface{}) BfLoadchunkKey {
	c.cs = append(c.cs, key)
	return (BfLoadchunkKey)(c)
}

type BfLoadchunkData Completed

func (c BfLoadchunkData) GetArgs() []interface{} {
	return c.cs
}

type BfLoadchunkIterator Completed

func (c BfLoadchunkIterator) Data(data interface{}) BfLoadchunkData {
	c.cs = append(c.cs, data)
	return (BfLoadchunkData)(c)
}

type BfLoadchunkKey Completed

func (c BfLoadchunkKey) Iterator(iterator int64) BfLoadchunkIterator {
	c.cs = append(c.cs, strconv.FormatInt(iterator, 10))
	return (BfLoadchunkIterator)(c)
}

type BfMadd Completed

func (b Builder) BfMadd() (c BfMadd) {
	c.cs = append(c.cs, "BF.MADD")
	return c
}

func (c BfMadd) Key(key interface{}) BfMaddKey {
	c.cs = append(c.cs, key)
	return (BfMaddKey)(c)
}

type BfMaddItem Completed

func (c BfMaddItem) Item(item ...interface{}) BfMaddItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c BfMaddItem) GetArgs() []interface{} {
	return c.cs
}

type BfMaddKey Completed

func (c BfMaddKey) Item(item ...interface{}) BfMaddItem {
	c.cs = append(c.cs, item...)
	return (BfMaddItem)(c)
}

type BfMexists Completed

func (b Builder) BfMexists() (c BfMexists) {
	c.cs = append(c.cs, "BF.MEXISTS")
	return c
}

func (c BfMexists) Key(key interface{}) BfMexistsKey {
	c.cs = append(c.cs, key)
	return (BfMexistsKey)(c)
}

type BfMexistsItem Completed

func (c BfMexistsItem) Item(item ...interface{}) BfMexistsItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c BfMexistsItem) GetArgs() []interface{} {
	return c.cs
}

type BfMexistsKey Completed

func (c BfMexistsKey) Item(item ...interface{}) BfMexistsItem {
	c.cs = append(c.cs, item...)
	return (BfMexistsItem)(c)
}

type BfReserve Completed

func (b Builder) BfReserve() (c BfReserve) {
	c.cs = append(c.cs, "BF.RESERVE")
	return c
}

func (c BfReserve) Key(key interface{}) BfReserveKey {
	c.cs = append(c.cs, key)
	return (BfReserveKey)(c)
}

type BfReserveCapacity Completed

func (c BfReserveCapacity) Expansion(expansion int64) BfReserveExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (BfReserveExpansion)(c)
}

func (c BfReserveCapacity) Nonscaling() BfReserveNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfReserveNonscaling)(c)
}

func (c BfReserveCapacity) GetArgs() []interface{} {
	return c.cs
}

type BfReserveErrorRate Completed

func (c BfReserveErrorRate) Capacity(capacity int64) BfReserveCapacity {
	c.cs = append(c.cs, strconv.FormatInt(capacity, 10))
	return (BfReserveCapacity)(c)
}

type BfReserveExpansion Completed

func (c BfReserveExpansion) Nonscaling() BfReserveNonscaling {
	c.cs = append(c.cs, "NONSCALING")
	return (BfReserveNonscaling)(c)
}

func (c BfReserveExpansion) GetArgs() []interface{} {
	return c.cs
}

type BfReserveKey Completed

func (c BfReserveKey) ErrorRate(errorRate float64) BfReserveErrorRate {
	c.cs = append(c.cs, strconv.FormatFloat(errorRate, 'f', -1, 64))
	return (BfReserveErrorRate)(c)
}

type BfReserveNonscaling Completed

func (c BfReserveNonscaling) GetArgs() []interface{} {
	return c.cs
}

type BfScandump Completed

func (b Builder) BfScandump() (c BfScandump) {
	c.cs = append(c.cs, "BF.SCANDUMP")
	return c
}

func (c BfScandump) Key(key interface{}) BfScandumpKey {
	c.cs = append(c.cs, key)
	return (BfScandumpKey)(c)
}

type BfScandumpIterator Completed

func (c BfScandumpIterator) GetArgs() []interface{} {
	return c.cs
}

type BfScandumpKey Completed

func (c BfScandumpKey) Iterator(iterator int64) BfScandumpIterator {
	c.cs = append(c.cs, strconv.FormatInt(iterator, 10))
	return (BfScandumpIterator)(c)
}

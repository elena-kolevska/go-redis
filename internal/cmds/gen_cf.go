// Code generated DO NOT EDIT

package cmds

import "strconv"

type CfAdd Completed

func (b Builder) CfAdd() (c CfAdd) {
	c.cs = append(c.cs, "CF.ADD")
	return c
}

func (c CfAdd) Key(key interface{}) CfAddKey {
	c.cs = append(c.cs, key)
	return (CfAddKey)(c)
}

type CfAddItem Completed

func (c CfAddItem) GetArgs() []interface{} {
	return c.cs
}

type CfAddKey Completed

func (c CfAddKey) Item(item interface{}) CfAddItem {
	c.cs = append(c.cs, item)
	return (CfAddItem)(c)
}

type CfAddnx Completed

func (b Builder) CfAddnx() (c CfAddnx) {
	c.cs = append(c.cs, "CF.ADDNX")
	return c
}

func (c CfAddnx) Key(key interface{}) CfAddnxKey {
	c.cs = append(c.cs, key)
	return (CfAddnxKey)(c)
}

type CfAddnxItem Completed

func (c CfAddnxItem) GetArgs() []interface{} {
	return c.cs
}

type CfAddnxKey Completed

func (c CfAddnxKey) Item(item interface{}) CfAddnxItem {
	c.cs = append(c.cs, item)
	return (CfAddnxItem)(c)
}

type CfCount Completed

func (b Builder) CfCount() (c CfCount) {
	c.cs = append(c.cs, "CF.COUNT")
	return c
}

func (c CfCount) Key(key interface{}) CfCountKey {
	c.cs = append(c.cs, key)
	return (CfCountKey)(c)
}

type CfCountItem Completed

func (c CfCountItem) GetArgs() []interface{} {
	return c.cs
}

type CfCountKey Completed

func (c CfCountKey) Item(item interface{}) CfCountItem {
	c.cs = append(c.cs, item)
	return (CfCountItem)(c)
}

type CfDel Completed

func (b Builder) CfDel() (c CfDel) {
	c.cs = append(c.cs, "CF.DEL")
	return c
}

func (c CfDel) Key(key interface{}) CfDelKey {
	c.cs = append(c.cs, key)
	return (CfDelKey)(c)
}

type CfDelItem Completed

func (c CfDelItem) GetArgs() []interface{} {
	return c.cs
}

type CfDelKey Completed

func (c CfDelKey) Item(item interface{}) CfDelItem {
	c.cs = append(c.cs, item)
	return (CfDelItem)(c)
}

type CfExists Completed

func (b Builder) CfExists() (c CfExists) {
	c.cs = append(c.cs, "CF.EXISTS")
	return c
}

func (c CfExists) Key(key interface{}) CfExistsKey {
	c.cs = append(c.cs, key)
	return (CfExistsKey)(c)
}

type CfExistsItem Completed

func (c CfExistsItem) GetArgs() []interface{} {
	return c.cs
}

type CfExistsKey Completed

func (c CfExistsKey) Item(item interface{}) CfExistsItem {
	c.cs = append(c.cs, item)
	return (CfExistsItem)(c)
}

type CfInfo Completed

func (b Builder) CfInfo() (c CfInfo) {
	c.cs = append(c.cs, "CF.INFO")
	return c
}

func (c CfInfo) Key(key interface{}) CfInfoKey {
	c.cs = append(c.cs, key)
	return (CfInfoKey)(c)
}

type CfInfoKey Completed

func (c CfInfoKey) GetArgs() []interface{} {
	return c.cs
}

type CfInsert Completed

func (b Builder) CfInsert() (c CfInsert) {
	c.cs = append(c.cs, "CF.INSERT")
	return c
}

func (c CfInsert) Key(key interface{}) CfInsertKey {
	c.cs = append(c.cs, key)
	return (CfInsertKey)(c)
}

type CfInsertCapacity Completed

func (c CfInsertCapacity) Nocreate() CfInsertNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (CfInsertNocreate)(c)
}

func (c CfInsertCapacity) Items() CfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (CfInsertItems)(c)
}

type CfInsertItem Completed

func (c CfInsertItem) Item(item ...interface{}) CfInsertItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c CfInsertItem) GetArgs() []interface{} {
	return c.cs
}

type CfInsertItems Completed

func (c CfInsertItems) Item(item ...interface{}) CfInsertItem {
	c.cs = append(c.cs, item...)
	return (CfInsertItem)(c)
}

type CfInsertKey Completed

func (c CfInsertKey) Capacity(capacity int64) CfInsertCapacity {
	c.cs = append(c.cs, "CAPACITY", strconv.FormatInt(capacity, 10))
	return (CfInsertCapacity)(c)
}

func (c CfInsertKey) Nocreate() CfInsertNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (CfInsertNocreate)(c)
}

func (c CfInsertKey) Items() CfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (CfInsertItems)(c)
}

type CfInsertNocreate Completed

func (c CfInsertNocreate) Items() CfInsertItems {
	c.cs = append(c.cs, "ITEMS")
	return (CfInsertItems)(c)
}

type CfInsertnx Completed

func (b Builder) CfInsertnx() (c CfInsertnx) {
	c.cs = append(c.cs, "CF.INSERTNX")
	return c
}

func (c CfInsertnx) Key(key interface{}) CfInsertnxKey {
	c.cs = append(c.cs, key)
	return (CfInsertnxKey)(c)
}

type CfInsertnxCapacity Completed

func (c CfInsertnxCapacity) Nocreate() CfInsertnxNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (CfInsertnxNocreate)(c)
}

func (c CfInsertnxCapacity) Items() CfInsertnxItems {
	c.cs = append(c.cs, "ITEMS")
	return (CfInsertnxItems)(c)
}

type CfInsertnxItem Completed

func (c CfInsertnxItem) Item(item ...interface{}) CfInsertnxItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c CfInsertnxItem) GetArgs() []interface{} {
	return c.cs
}

type CfInsertnxItems Completed

func (c CfInsertnxItems) Item(item ...interface{}) CfInsertnxItem {
	c.cs = append(c.cs, item...)
	return (CfInsertnxItem)(c)
}

type CfInsertnxKey Completed

func (c CfInsertnxKey) Capacity(capacity int64) CfInsertnxCapacity {
	c.cs = append(c.cs, "CAPACITY", strconv.FormatInt(capacity, 10))
	return (CfInsertnxCapacity)(c)
}

func (c CfInsertnxKey) Nocreate() CfInsertnxNocreate {
	c.cs = append(c.cs, "NOCREATE")
	return (CfInsertnxNocreate)(c)
}

func (c CfInsertnxKey) Items() CfInsertnxItems {
	c.cs = append(c.cs, "ITEMS")
	return (CfInsertnxItems)(c)
}

type CfInsertnxNocreate Completed

func (c CfInsertnxNocreate) Items() CfInsertnxItems {
	c.cs = append(c.cs, "ITEMS")
	return (CfInsertnxItems)(c)
}

type CfLoadchunk Completed

func (b Builder) CfLoadchunk() (c CfLoadchunk) {
	c.cs = append(c.cs, "CF.LOADCHUNK")
	return c
}

func (c CfLoadchunk) Key(key interface{}) CfLoadchunkKey {
	c.cs = append(c.cs, key)
	return (CfLoadchunkKey)(c)
}

type CfLoadchunkData Completed

func (c CfLoadchunkData) GetArgs() []interface{} {
	return c.cs
}

type CfLoadchunkIterator Completed

func (c CfLoadchunkIterator) Data(data interface{}) CfLoadchunkData {
	c.cs = append(c.cs, data)
	return (CfLoadchunkData)(c)
}

type CfLoadchunkKey Completed

func (c CfLoadchunkKey) Iterator(iterator int64) CfLoadchunkIterator {
	c.cs = append(c.cs, strconv.FormatInt(iterator, 10))
	return (CfLoadchunkIterator)(c)
}

type CfMexists Completed

func (b Builder) CfMexists() (c CfMexists) {
	c.cs = append(c.cs, "CF.MEXISTS")
	return c
}

func (c CfMexists) Key(key interface{}) CfMexistsKey {
	c.cs = append(c.cs, key)
	return (CfMexistsKey)(c)
}

type CfMexistsItem Completed

func (c CfMexistsItem) Item(item ...interface{}) CfMexistsItem {
	c.cs = append(c.cs, item...)
	return c
}

func (c CfMexistsItem) GetArgs() []interface{} {
	return c.cs
}

type CfMexistsKey Completed

func (c CfMexistsKey) Item(item ...interface{}) CfMexistsItem {
	c.cs = append(c.cs, item...)
	return (CfMexistsItem)(c)
}

type CfReserve Completed

func (b Builder) CfReserve() (c CfReserve) {
	c.cs = append(c.cs, "CF.RESERVE")
	return c
}

func (c CfReserve) Key(key interface{}) CfReserveKey {
	c.cs = append(c.cs, key)
	return (CfReserveKey)(c)
}

type CfReserveBucketsize Completed

func (c CfReserveBucketsize) Maxiterations(maxiterations int64) CfReserveMaxiterations {
	c.cs = append(c.cs, "MAXITERATIONS", strconv.FormatInt(maxiterations, 10))
	return (CfReserveMaxiterations)(c)
}

func (c CfReserveBucketsize) Expansion(expansion int64) CfReserveExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (CfReserveExpansion)(c)
}

func (c CfReserveBucketsize) GetArgs() []interface{} {
	return c.cs
}

type CfReserveCapacity Completed

func (c CfReserveCapacity) Bucketsize(bucketsize int64) CfReserveBucketsize {
	c.cs = append(c.cs, "BUCKETSIZE", strconv.FormatInt(bucketsize, 10))
	return (CfReserveBucketsize)(c)
}

func (c CfReserveCapacity) Maxiterations(maxiterations int64) CfReserveMaxiterations {
	c.cs = append(c.cs, "MAXITERATIONS", strconv.FormatInt(maxiterations, 10))
	return (CfReserveMaxiterations)(c)
}

func (c CfReserveCapacity) Expansion(expansion int64) CfReserveExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (CfReserveExpansion)(c)
}

func (c CfReserveCapacity) GetArgs() []interface{} {
	return c.cs
}

type CfReserveExpansion Completed

func (c CfReserveExpansion) GetArgs() []interface{} {
	return c.cs
}

type CfReserveKey Completed

func (c CfReserveKey) Capacity(capacity int64) CfReserveCapacity {
	c.cs = append(c.cs, strconv.FormatInt(capacity, 10))
	return (CfReserveCapacity)(c)
}

type CfReserveMaxiterations Completed

func (c CfReserveMaxiterations) Expansion(expansion int64) CfReserveExpansion {
	c.cs = append(c.cs, "EXPANSION", strconv.FormatInt(expansion, 10))
	return (CfReserveExpansion)(c)
}

func (c CfReserveMaxiterations) GetArgs() []interface{} {
	return c.cs
}

type CfScandump Completed

func (b Builder) CfScandump() (c CfScandump) {
	c.cs = append(c.cs, "CF.SCANDUMP")
	return c
}

func (c CfScandump) Key(key interface{}) CfScandumpKey {
	c.cs = append(c.cs, key)
	return (CfScandumpKey)(c)
}

type CfScandumpIterator Completed

func (c CfScandumpIterator) GetArgs() []interface{} {
	return c.cs
}

type CfScandumpKey Completed

func (c CfScandumpKey) Iterator(iterator int64) CfScandumpIterator {
	c.cs = append(c.cs, strconv.FormatInt(iterator, 10))
	return (CfScandumpIterator)(c)
}

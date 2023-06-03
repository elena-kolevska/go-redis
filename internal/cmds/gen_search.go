// Code generated DO NOT EDIT

package cmds

import "strconv"

type FtAggregate Completed

func (b Builder) FtAggregate() (c FtAggregate) {
	c.cs = append(c.cs, "FT.AGGREGATE")
	return c
}

func (c FtAggregate) Index(index interface{}) FtAggregateIndex {
	c.cs = append(c.cs, index)
	return (FtAggregateIndex)(c)
}

type FtAggregateCursorCount Completed

func (c FtAggregateCursorCount) Maxidle(idleTime int64) FtAggregateCursorMaxidle {
	c.cs = append(c.cs, "MAXIDLE", strconv.FormatInt(idleTime, 10))
	return (FtAggregateCursorMaxidle)(c)
}

func (c FtAggregateCursorCount) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateCursorCount) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateCursorCount) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateCursorMaxidle Completed

func (c FtAggregateCursorMaxidle) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateCursorMaxidle) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateCursorMaxidle) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateCursorWithcursor Completed

func (c FtAggregateCursorWithcursor) Count(readSize int64) FtAggregateCursorCount {
	c.cs = append(c.cs, "COUNT", strconv.FormatInt(readSize, 10))
	return (FtAggregateCursorCount)(c)
}

func (c FtAggregateCursorWithcursor) Maxidle(idleTime int64) FtAggregateCursorMaxidle {
	c.cs = append(c.cs, "MAXIDLE", strconv.FormatInt(idleTime, 10))
	return (FtAggregateCursorMaxidle)(c)
}

func (c FtAggregateCursorWithcursor) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateCursorWithcursor) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateCursorWithcursor) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateDialect Completed

func (c FtAggregateDialect) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateIndex Completed

func (c FtAggregateIndex) Query(query interface{}) FtAggregateQuery {
	c.cs = append(c.cs, query)
	return (FtAggregateQuery)(c)
}

type FtAggregateOpApplyApply Completed

func (c FtAggregateOpApplyApply) As(name interface{}) FtAggregateOpApplyAs {
	c.cs = append(c.cs, "AS", name)
	return (FtAggregateOpApplyAs)(c)
}

type FtAggregateOpApplyAs Completed

func (c FtAggregateOpApplyAs) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpApplyAs) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpApplyAs) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpApplyAs) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpApplyAs) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpApplyAs) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpApplyAs) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpApplyAs) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpApplyAs) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpApplyAs) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpApplyAs) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpFilter Completed

func (c FtAggregateOpFilter) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpFilter) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpFilter) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpFilter) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpFilter) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpFilter) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpFilter) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return c
}

func (c FtAggregateOpFilter) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpFilter) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpFilter) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpFilter) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyGroupby Completed

func (c FtAggregateOpGroupbyGroupby) Property(property ...interface{}) FtAggregateOpGroupbyProperty {
	c.cs = append(c.cs, property...)
	return (FtAggregateOpGroupbyProperty)(c)
}

func (c FtAggregateOpGroupbyGroupby) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyGroupby) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return c
}

func (c FtAggregateOpGroupbyGroupby) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyGroupby) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyGroupby) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyGroupby) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyGroupby) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyGroupby) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyGroupby) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyGroupby) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyGroupby) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyGroupby) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyProperty Completed

func (c FtAggregateOpGroupbyProperty) Property(property ...interface{}) FtAggregateOpGroupbyProperty {
	c.cs = append(c.cs, property...)
	return c
}

func (c FtAggregateOpGroupbyProperty) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyProperty) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyProperty) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyProperty) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyProperty) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyProperty) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyProperty) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyProperty) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyProperty) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyProperty) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyProperty) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyProperty) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceArg Completed

func (c FtAggregateOpGroupbyReduceArg) Arg(arg ...interface{}) FtAggregateOpGroupbyReduceArg {
	c.cs = append(c.cs, arg...)
	return c
}

func (c FtAggregateOpGroupbyReduceArg) As(name interface{}) FtAggregateOpGroupbyReduceAs {
	c.cs = append(c.cs, "AS", name)
	return (FtAggregateOpGroupbyReduceAs)(c)
}

func (c FtAggregateOpGroupbyReduceArg) By(by interface{}) FtAggregateOpGroupbyReduceBy {
	c.cs = append(c.cs, "BY", by)
	return (FtAggregateOpGroupbyReduceBy)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	c.cs = append(c.cs, "ASC")
	return (FtAggregateOpGroupbyReduceOrderAsc)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	c.cs = append(c.cs, "DESC")
	return (FtAggregateOpGroupbyReduceOrderDesc)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyReduceArg) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyReduceArg) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyReduceArg) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceAs Completed

func (c FtAggregateOpGroupbyReduceAs) By(by interface{}) FtAggregateOpGroupbyReduceBy {
	c.cs = append(c.cs, "BY", by)
	return (FtAggregateOpGroupbyReduceBy)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	c.cs = append(c.cs, "ASC")
	return (FtAggregateOpGroupbyReduceOrderAsc)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	c.cs = append(c.cs, "DESC")
	return (FtAggregateOpGroupbyReduceOrderDesc)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyReduceAs) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyReduceAs) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyReduceAs) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceBy Completed

func (c FtAggregateOpGroupbyReduceBy) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	c.cs = append(c.cs, "ASC")
	return (FtAggregateOpGroupbyReduceOrderAsc)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	c.cs = append(c.cs, "DESC")
	return (FtAggregateOpGroupbyReduceOrderDesc)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyReduceBy) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyReduceBy) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyReduceBy) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceNargs Completed

func (c FtAggregateOpGroupbyReduceNargs) Arg(arg ...interface{}) FtAggregateOpGroupbyReduceArg {
	c.cs = append(c.cs, arg...)
	return (FtAggregateOpGroupbyReduceArg)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) As(name interface{}) FtAggregateOpGroupbyReduceAs {
	c.cs = append(c.cs, "AS", name)
	return (FtAggregateOpGroupbyReduceAs)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) By(by interface{}) FtAggregateOpGroupbyReduceBy {
	c.cs = append(c.cs, "BY", by)
	return (FtAggregateOpGroupbyReduceBy)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	c.cs = append(c.cs, "ASC")
	return (FtAggregateOpGroupbyReduceOrderAsc)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	c.cs = append(c.cs, "DESC")
	return (FtAggregateOpGroupbyReduceOrderDesc)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyReduceNargs) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceOrderAsc Completed

func (c FtAggregateOpGroupbyReduceOrderAsc) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceOrderDesc Completed

func (c FtAggregateOpGroupbyReduceOrderDesc) Reduce(function interface{}) FtAggregateOpGroupbyReduceReduce {
	c.cs = append(c.cs, "REDUCE", function)
	return (FtAggregateOpGroupbyReduceReduce)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpGroupbyReduceReduce Completed

func (c FtAggregateOpGroupbyReduceReduce) Nargs(nargs int64) FtAggregateOpGroupbyReduceNargs {
	c.cs = append(c.cs, strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyReduceNargs)(c)
}

type FtAggregateOpLimitLimit Completed

func (c FtAggregateOpLimitLimit) OffsetNum(offset int64, num int64) FtAggregateOpLimitOffsetNum {
	c.cs = append(c.cs, strconv.FormatInt(offset, 10), strconv.FormatInt(num, 10))
	return (FtAggregateOpLimitOffsetNum)(c)
}

type FtAggregateOpLimitOffsetNum Completed

func (c FtAggregateOpLimitOffsetNum) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpLimitOffsetNum) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpLimitOffsetNum) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpLimitOffsetNum) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpLimitOffsetNum) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpLimitOffsetNum) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpLimitOffsetNum) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpLimitOffsetNum) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpLimitOffsetNum) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpLimitOffsetNum) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpLimitOffsetNum) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpLoadField Completed

func (c FtAggregateOpLoadField) Field(field ...interface{}) FtAggregateOpLoadField {
	c.cs = append(c.cs, field...)
	return c
}

func (c FtAggregateOpLoadField) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpLoadField) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpLoadField) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpLoadField) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpLoadField) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpLoadField) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpLoadField) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpLoadField) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpLoadField) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpLoadField) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpLoadField) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpLoadLoad Completed

func (c FtAggregateOpLoadLoad) Field(field ...interface{}) FtAggregateOpLoadField {
	c.cs = append(c.cs, field...)
	return (FtAggregateOpLoadField)(c)
}

type FtAggregateOpLoadallLoadAll Completed

func (c FtAggregateOpLoadallLoadAll) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpLoadallLoadAll) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpLoadallLoadAll) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpLoadallLoadAll) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpLoadallLoadAll) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpLoadallLoadAll) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpLoadallLoadAll) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return c
}

func (c FtAggregateOpLoadallLoadAll) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpLoadallLoadAll) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpLoadallLoadAll) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpLoadallLoadAll) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpSortbyFieldsOrderAsc Completed

func (c FtAggregateOpSortbyFieldsOrderAsc) Property(property interface{}) FtAggregateOpSortbyFieldsProperty {
	c.cs = append(c.cs, property)
	return (FtAggregateOpSortbyFieldsProperty)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Max(num int64) FtAggregateOpSortbyMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(num, 10))
	return (FtAggregateOpSortbyMax)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpSortbyFieldsOrderDesc Completed

func (c FtAggregateOpSortbyFieldsOrderDesc) Property(property interface{}) FtAggregateOpSortbyFieldsProperty {
	c.cs = append(c.cs, property)
	return (FtAggregateOpSortbyFieldsProperty)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Max(num int64) FtAggregateOpSortbyMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(num, 10))
	return (FtAggregateOpSortbyMax)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpSortbyFieldsProperty Completed

func (c FtAggregateOpSortbyFieldsProperty) Asc() FtAggregateOpSortbyFieldsOrderAsc {
	c.cs = append(c.cs, "ASC")
	return (FtAggregateOpSortbyFieldsOrderAsc)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Desc() FtAggregateOpSortbyFieldsOrderDesc {
	c.cs = append(c.cs, "DESC")
	return (FtAggregateOpSortbyFieldsOrderDesc)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Property(property interface{}) FtAggregateOpSortbyFieldsProperty {
	c.cs = append(c.cs, property)
	return c
}

func (c FtAggregateOpSortbyFieldsProperty) Max(num int64) FtAggregateOpSortbyMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(num, 10))
	return (FtAggregateOpSortbyMax)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpSortbyFieldsProperty) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpSortbyMax Completed

func (c FtAggregateOpSortbyMax) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpSortbyMax) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpSortbyMax) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpSortbyMax) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpSortbyMax) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpSortbyMax) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpSortbyMax) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateOpSortbyMax) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpSortbyMax) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpSortbyMax) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpSortbyMax) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateOpSortbySortby Completed

func (c FtAggregateOpSortbySortby) Property(property interface{}) FtAggregateOpSortbyFieldsProperty {
	c.cs = append(c.cs, property)
	return (FtAggregateOpSortbyFieldsProperty)(c)
}

func (c FtAggregateOpSortbySortby) Max(num int64) FtAggregateOpSortbyMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(num, 10))
	return (FtAggregateOpSortbyMax)(c)
}

func (c FtAggregateOpSortbySortby) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateOpSortbySortby) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateOpSortbySortby) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateOpSortbySortby) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateOpSortbySortby) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateOpSortbySortby) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateOpSortbySortby) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return c
}

func (c FtAggregateOpSortbySortby) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateOpSortbySortby) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateOpSortbySortby) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateOpSortbySortby) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateParamsNameValue Completed

func (c FtAggregateParamsNameValue) NameValue(name interface{}, value interface{}) FtAggregateParamsNameValue {
	c.cs = append(c.cs, name, value)
	return c
}

func (c FtAggregateParamsNameValue) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateParamsNameValue) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateParamsNargs Completed

func (c FtAggregateParamsNargs) NameValue() FtAggregateParamsNameValue {
	return (FtAggregateParamsNameValue)(c)
}

type FtAggregateParamsParams Completed

func (c FtAggregateParamsParams) Nargs(nargs int64) FtAggregateParamsNargs {
	c.cs = append(c.cs, strconv.FormatInt(nargs, 10))
	return (FtAggregateParamsNargs)(c)
}

type FtAggregateQuery Completed

func (c FtAggregateQuery) Verbatim() FtAggregateVerbatim {
	c.cs = append(c.cs, "VERBATIM")
	return (FtAggregateVerbatim)(c)
}

func (c FtAggregateQuery) Timeout(timeout int64) FtAggregateTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtAggregateTimeout)(c)
}

func (c FtAggregateQuery) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateQuery) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateQuery) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateQuery) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateQuery) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateQuery) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateQuery) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateQuery) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateQuery) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateQuery) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateQuery) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateTimeout Completed

func (c FtAggregateTimeout) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateTimeout) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateTimeout) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateTimeout) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateTimeout) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateTimeout) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateTimeout) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateTimeout) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateTimeout) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateTimeout) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateTimeout) GetArgs() []interface{} {
	return c.cs
}

type FtAggregateVerbatim Completed

func (c FtAggregateVerbatim) Timeout(timeout int64) FtAggregateTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtAggregateTimeout)(c)
}

func (c FtAggregateVerbatim) LoadAll() FtAggregateOpLoadallLoadAll {
	c.cs = append(c.cs, "LOAD", "*")
	return (FtAggregateOpLoadallLoadAll)(c)
}

func (c FtAggregateVerbatim) Load(count int64) FtAggregateOpLoadLoad {
	c.cs = append(c.cs, "LOAD", strconv.FormatInt(count, 10))
	return (FtAggregateOpLoadLoad)(c)
}

func (c FtAggregateVerbatim) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	c.cs = append(c.cs, "GROUPBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpGroupbyGroupby)(c)
}

func (c FtAggregateVerbatim) Sortby(nargs int64) FtAggregateOpSortbySortby {
	c.cs = append(c.cs, "SORTBY", strconv.FormatInt(nargs, 10))
	return (FtAggregateOpSortbySortby)(c)
}

func (c FtAggregateVerbatim) Apply(expression interface{}) FtAggregateOpApplyApply {
	c.cs = append(c.cs, "APPLY", expression)
	return (FtAggregateOpApplyApply)(c)
}

func (c FtAggregateVerbatim) Limit() FtAggregateOpLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtAggregateOpLimitLimit)(c)
}

func (c FtAggregateVerbatim) Filter(filter interface{}) FtAggregateOpFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtAggregateOpFilter)(c)
}

func (c FtAggregateVerbatim) Withcursor() FtAggregateCursorWithcursor {
	c.cs = append(c.cs, "WITHCURSOR")
	return (FtAggregateCursorWithcursor)(c)
}

func (c FtAggregateVerbatim) Params() FtAggregateParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtAggregateParamsParams)(c)
}

func (c FtAggregateVerbatim) Dialect(dialect int64) FtAggregateDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtAggregateDialect)(c)
}

func (c FtAggregateVerbatim) GetArgs() []interface{} {
	return c.cs
}

type FtAliasadd Completed

func (b Builder) FtAliasadd() (c FtAliasadd) {
	c.cs = append(c.cs, "FT.ALIASADD")
	return c
}

func (c FtAliasadd) Alias(alias interface{}) FtAliasaddAlias {
	c.cs = append(c.cs, alias)
	return (FtAliasaddAlias)(c)
}

type FtAliasaddAlias Completed

func (c FtAliasaddAlias) Index(index interface{}) FtAliasaddIndex {
	c.cs = append(c.cs, index)
	return (FtAliasaddIndex)(c)
}

type FtAliasaddIndex Completed

func (c FtAliasaddIndex) GetArgs() []interface{} {
	return c.cs
}

type FtAliasdel Completed

func (b Builder) FtAliasdel() (c FtAliasdel) {
	c.cs = append(c.cs, "FT.ALIASDEL")
	return c
}

func (c FtAliasdel) Alias(alias interface{}) FtAliasdelAlias {
	c.cs = append(c.cs, alias)
	return (FtAliasdelAlias)(c)
}

type FtAliasdelAlias Completed

func (c FtAliasdelAlias) GetArgs() []interface{} {
	return c.cs
}

type FtAliasupdate Completed

func (b Builder) FtAliasupdate() (c FtAliasupdate) {
	c.cs = append(c.cs, "FT.ALIASUPDATE")
	return c
}

func (c FtAliasupdate) Alias(alias interface{}) FtAliasupdateAlias {
	c.cs = append(c.cs, alias)
	return (FtAliasupdateAlias)(c)
}

type FtAliasupdateAlias Completed

func (c FtAliasupdateAlias) Index(index interface{}) FtAliasupdateIndex {
	c.cs = append(c.cs, index)
	return (FtAliasupdateIndex)(c)
}

type FtAliasupdateIndex Completed

func (c FtAliasupdateIndex) GetArgs() []interface{} {
	return c.cs
}

type FtAlter Completed

func (b Builder) FtAlter() (c FtAlter) {
	c.cs = append(c.cs, "FT.ALTER")
	return c
}

func (c FtAlter) Index(index interface{}) FtAlterIndex {
	c.cs = append(c.cs, index)
	return (FtAlterIndex)(c)
}

type FtAlterAdd Completed

func (c FtAlterAdd) Field(field interface{}) FtAlterField {
	c.cs = append(c.cs, field)
	return (FtAlterField)(c)
}

type FtAlterField Completed

func (c FtAlterField) Options(options interface{}) FtAlterOptions {
	c.cs = append(c.cs, options)
	return (FtAlterOptions)(c)
}

type FtAlterIndex Completed

func (c FtAlterIndex) Skipinitialscan() FtAlterSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtAlterSkipinitialscan)(c)
}

func (c FtAlterIndex) Schema() FtAlterSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtAlterSchema)(c)
}

type FtAlterOptions Completed

func (c FtAlterOptions) GetArgs() []interface{} {
	return c.cs
}

type FtAlterSchema Completed

func (c FtAlterSchema) Add() FtAlterAdd {
	c.cs = append(c.cs, "ADD")
	return (FtAlterAdd)(c)
}

type FtAlterSkipinitialscan Completed

func (c FtAlterSkipinitialscan) Schema() FtAlterSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtAlterSchema)(c)
}

type FtConfigGet Completed

func (b Builder) FtConfigGet() (c FtConfigGet) {
	c.cs = append(c.cs, "FT.CONFIG", "GET")
	return c
}

func (c FtConfigGet) Option(option interface{}) FtConfigGetOption {
	c.cs = append(c.cs, option)
	return (FtConfigGetOption)(c)
}

type FtConfigGetOption Completed

func (c FtConfigGetOption) GetArgs() []interface{} {
	return c.cs
}

type FtConfigHelp Completed

func (b Builder) FtConfigHelp() (c FtConfigHelp) {
	c.cs = append(c.cs, "FT.CONFIG", "HELP")
	return c
}

func (c FtConfigHelp) Option(option interface{}) FtConfigHelpOption {
	c.cs = append(c.cs, option)
	return (FtConfigHelpOption)(c)
}

type FtConfigHelpOption Completed

func (c FtConfigHelpOption) GetArgs() []interface{} {
	return c.cs
}

type FtConfigSet Completed

func (b Builder) FtConfigSet() (c FtConfigSet) {
	c.cs = append(c.cs, "FT.CONFIG", "SET")
	return c
}

func (c FtConfigSet) Option(option interface{}) FtConfigSetOption {
	c.cs = append(c.cs, option)
	return (FtConfigSetOption)(c)
}

type FtConfigSetOption Completed

func (c FtConfigSetOption) Value(value interface{}) FtConfigSetValue {
	c.cs = append(c.cs, value)
	return (FtConfigSetValue)(c)
}

type FtConfigSetValue Completed

func (c FtConfigSetValue) GetArgs() []interface{} {
	return c.cs
}

type FtCreate Completed

func (b Builder) FtCreate() (c FtCreate) {
	c.cs = append(c.cs, "FT.CREATE")
	return c
}

func (c FtCreate) Index(index interface{}) FtCreateIndex {
	c.cs = append(c.cs, index)
	return (FtCreateIndex)(c)
}

type FtCreateFieldAs Completed

func (c FtCreateFieldAs) Text() FtCreateFieldFieldTypeText {
	c.cs = append(c.cs, "TEXT")
	return (FtCreateFieldFieldTypeText)(c)
}

func (c FtCreateFieldAs) Tag() FtCreateFieldFieldTypeTag {
	c.cs = append(c.cs, "TAG")
	return (FtCreateFieldFieldTypeTag)(c)
}

func (c FtCreateFieldAs) Numeric() FtCreateFieldFieldTypeNumeric {
	c.cs = append(c.cs, "NUMERIC")
	return (FtCreateFieldFieldTypeNumeric)(c)
}

func (c FtCreateFieldAs) Geo() FtCreateFieldFieldTypeGeo {
	c.cs = append(c.cs, "GEO")
	return (FtCreateFieldFieldTypeGeo)(c)
}

func (c FtCreateFieldAs) Vector(algo interface{}, nargs int64, args ...interface{}) FtCreateFieldFieldTypeVector {
	c.cs = append(c.cs, "VECTOR", algo, strconv.FormatInt(nargs, 10))
	c.cs = append(c.cs, args...)
	return (FtCreateFieldFieldTypeVector)(c)
}

type FtCreateFieldFieldName Completed

func (c FtCreateFieldFieldName) As(alias interface{}) FtCreateFieldAs {
	c.cs = append(c.cs, "AS", alias)
	return (FtCreateFieldAs)(c)
}

func (c FtCreateFieldFieldName) Text() FtCreateFieldFieldTypeText {
	c.cs = append(c.cs, "TEXT")
	return (FtCreateFieldFieldTypeText)(c)
}

func (c FtCreateFieldFieldName) Tag() FtCreateFieldFieldTypeTag {
	c.cs = append(c.cs, "TAG")
	return (FtCreateFieldFieldTypeTag)(c)
}

func (c FtCreateFieldFieldName) Numeric() FtCreateFieldFieldTypeNumeric {
	c.cs = append(c.cs, "NUMERIC")
	return (FtCreateFieldFieldTypeNumeric)(c)
}

func (c FtCreateFieldFieldName) Geo() FtCreateFieldFieldTypeGeo {
	c.cs = append(c.cs, "GEO")
	return (FtCreateFieldFieldTypeGeo)(c)
}

func (c FtCreateFieldFieldName) Vector(algo interface{}, nargs int64, args ...interface{}) FtCreateFieldFieldTypeVector {
	c.cs = append(c.cs, "VECTOR", algo, strconv.FormatInt(nargs, 10))
	c.cs = append(c.cs, args...)
	return (FtCreateFieldFieldTypeVector)(c)
}

type FtCreateFieldFieldTypeGeo Completed

func (c FtCreateFieldFieldTypeGeo) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldFieldTypeGeo) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldFieldTypeGeo) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldFieldTypeGeo) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldFieldTypeGeo) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldFieldTypeGeo) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldFieldTypeGeo) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldFieldTypeGeo) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldFieldTypeGeo) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldFieldTypeGeo) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldFieldTypeNumeric Completed

func (c FtCreateFieldFieldTypeNumeric) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldFieldTypeNumeric) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldFieldTypeNumeric) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldFieldTypeNumeric) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldFieldTypeTag Completed

func (c FtCreateFieldFieldTypeTag) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldFieldTypeTag) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldFieldTypeTag) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldFieldTypeTag) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldFieldTypeTag) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldFieldTypeTag) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldFieldTypeTag) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldFieldTypeTag) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldFieldTypeTag) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldFieldTypeTag) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldFieldTypeText Completed

func (c FtCreateFieldFieldTypeText) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldFieldTypeText) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldFieldTypeText) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldFieldTypeText) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldFieldTypeText) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldFieldTypeText) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldFieldTypeText) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldFieldTypeText) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldFieldTypeText) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldFieldTypeText) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldFieldTypeVector Completed

func (c FtCreateFieldFieldTypeVector) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldFieldTypeVector) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldFieldTypeVector) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldFieldTypeVector) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldFieldTypeVector) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldFieldTypeVector) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldFieldTypeVector) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldFieldTypeVector) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldFieldTypeVector) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldFieldTypeVector) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionCasesensitive Completed

func (c FtCreateFieldOptionCasesensitive) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionCasesensitive) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionCasesensitive) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionCasesensitive) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionCasesensitive) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionCasesensitive) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionCasesensitive) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionCasesensitive) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return c
}

func (c FtCreateFieldOptionCasesensitive) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionCasesensitive) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionNoindex Completed

func (c FtCreateFieldOptionNoindex) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionNoindex) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionNoindex) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionNoindex) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionNoindex) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionNoindex) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionNoindex) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionNoindex) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return c
}

func (c FtCreateFieldOptionNoindex) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionNoindex) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionNostem Completed

func (c FtCreateFieldOptionNostem) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionNostem) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionNostem) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionNostem) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionNostem) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionNostem) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionNostem) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionNostem) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return c
}

func (c FtCreateFieldOptionNostem) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionNostem) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionPhonetic Completed

func (c FtCreateFieldOptionPhonetic) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionPhonetic) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionPhonetic) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionPhonetic) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionPhonetic) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionPhonetic) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionPhonetic) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionPhonetic) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return c
}

func (c FtCreateFieldOptionPhonetic) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionPhonetic) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionSeparator Completed

func (c FtCreateFieldOptionSeparator) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionSeparator) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionSeparator) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionSeparator) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionSeparator) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionSeparator) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionSeparator) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionSeparator) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return c
}

func (c FtCreateFieldOptionSeparator) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionSeparator) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionSortableSortable Completed

func (c FtCreateFieldOptionSortableSortable) Unf() FtCreateFieldOptionSortableUnf {
	c.cs = append(c.cs, "UNF")
	return (FtCreateFieldOptionSortableUnf)(c)
}

func (c FtCreateFieldOptionSortableSortable) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionSortableSortable) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionSortableSortable) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionSortableSortable) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionSortableSortable) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionSortableSortable) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionSortableSortable) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionSortableSortable) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return c
}

func (c FtCreateFieldOptionSortableSortable) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionSortableSortable) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionSortableUnf Completed

func (c FtCreateFieldOptionSortableUnf) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionSortableUnf) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionSortableUnf) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionSortableUnf) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionSortableUnf) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionSortableUnf) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionSortableUnf) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionSortableUnf) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionSortableUnf) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionSortableUnf) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionWeight Completed

func (c FtCreateFieldOptionWeight) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionWeight) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionWeight) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return (FtCreateFieldOptionWithsuffixtrie)(c)
}

func (c FtCreateFieldOptionWeight) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionWeight) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionWeight) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionWeight) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionWeight) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return c
}

func (c FtCreateFieldOptionWeight) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionWeight) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFieldOptionWithsuffixtrie Completed

func (c FtCreateFieldOptionWithsuffixtrie) Sortable() FtCreateFieldOptionSortableSortable {
	c.cs = append(c.cs, "SORTABLE")
	return (FtCreateFieldOptionSortableSortable)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Noindex() FtCreateFieldOptionNoindex {
	c.cs = append(c.cs, "NOINDEX")
	return (FtCreateFieldOptionNoindex)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Nostem() FtCreateFieldOptionNostem {
	c.cs = append(c.cs, "NOSTEM")
	return (FtCreateFieldOptionNostem)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Phonetic(phonetic interface{}) FtCreateFieldOptionPhonetic {
	c.cs = append(c.cs, "PHONETIC", phonetic)
	return (FtCreateFieldOptionPhonetic)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Weight(weight float64) FtCreateFieldOptionWeight {
	c.cs = append(c.cs, "WEIGHT", strconv.FormatFloat(weight, 'f', -1, 64))
	return (FtCreateFieldOptionWeight)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Separator(separator interface{}) FtCreateFieldOptionSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtCreateFieldOptionSeparator)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Casesensitive() FtCreateFieldOptionCasesensitive {
	c.cs = append(c.cs, "CASESENSITIVE")
	return (FtCreateFieldOptionCasesensitive)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	c.cs = append(c.cs, "WITHSUFFIXTRIE")
	return c
}

func (c FtCreateFieldOptionWithsuffixtrie) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

func (c FtCreateFieldOptionWithsuffixtrie) GetArgs() []interface{} {
	return c.cs
}

type FtCreateFilter Completed

func (c FtCreateFilter) Language(defaultLang interface{}) FtCreateLanguage {
	c.cs = append(c.cs, "LANGUAGE", defaultLang)
	return (FtCreateLanguage)(c)
}

func (c FtCreateFilter) LanguageField(langAttribute interface{}) FtCreateLanguageField {
	c.cs = append(c.cs, "LANGUAGE_FIELD", langAttribute)
	return (FtCreateLanguageField)(c)
}

func (c FtCreateFilter) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreateFilter) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateFilter) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateFilter) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateFilter) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateFilter) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateFilter) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateFilter) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateFilter) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateFilter) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateFilter) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateFilter) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateIndex Completed

func (c FtCreateIndex) OnHash() FtCreateOnHash {
	c.cs = append(c.cs, "ON", "HASH")
	return (FtCreateOnHash)(c)
}

func (c FtCreateIndex) OnJson() FtCreateOnJson {
	c.cs = append(c.cs, "ON", "JSON")
	return (FtCreateOnJson)(c)
}

func (c FtCreateIndex) Prefix(count int64) FtCreatePrefixCount {
	c.cs = append(c.cs, "PREFIX", strconv.FormatInt(count, 10))
	return (FtCreatePrefixCount)(c)
}

func (c FtCreateIndex) Filter(filter interface{}) FtCreateFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtCreateFilter)(c)
}

func (c FtCreateIndex) Language(defaultLang interface{}) FtCreateLanguage {
	c.cs = append(c.cs, "LANGUAGE", defaultLang)
	return (FtCreateLanguage)(c)
}

func (c FtCreateIndex) LanguageField(langAttribute interface{}) FtCreateLanguageField {
	c.cs = append(c.cs, "LANGUAGE_FIELD", langAttribute)
	return (FtCreateLanguageField)(c)
}

func (c FtCreateIndex) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreateIndex) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateIndex) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateIndex) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateIndex) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateIndex) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateIndex) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateIndex) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateIndex) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateIndex) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateIndex) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateIndex) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateLanguage Completed

func (c FtCreateLanguage) LanguageField(langAttribute interface{}) FtCreateLanguageField {
	c.cs = append(c.cs, "LANGUAGE_FIELD", langAttribute)
	return (FtCreateLanguageField)(c)
}

func (c FtCreateLanguage) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreateLanguage) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateLanguage) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateLanguage) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateLanguage) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateLanguage) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateLanguage) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateLanguage) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateLanguage) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateLanguage) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateLanguage) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateLanguage) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateLanguageField Completed

func (c FtCreateLanguageField) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreateLanguageField) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateLanguageField) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateLanguageField) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateLanguageField) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateLanguageField) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateLanguageField) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateLanguageField) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateLanguageField) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateLanguageField) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateLanguageField) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateLanguageField) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateMaxtextfields Completed

func (c FtCreateMaxtextfields) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateMaxtextfields) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateMaxtextfields) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateMaxtextfields) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateMaxtextfields) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateMaxtextfields) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateMaxtextfields) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateMaxtextfields) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateNofields Completed

func (c FtCreateNofields) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateNofields) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateNofields) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateNofields) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateNofreqs Completed

func (c FtCreateNofreqs) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateNofreqs) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateNofreqs) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateNohl Completed

func (c FtCreateNohl) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateNohl) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateNohl) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateNohl) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateNohl) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateNooffsets Completed

func (c FtCreateNooffsets) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateNooffsets) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateNooffsets) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateNooffsets) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateNooffsets) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateNooffsets) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateOnHash Completed

func (c FtCreateOnHash) Prefix(count int64) FtCreatePrefixCount {
	c.cs = append(c.cs, "PREFIX", strconv.FormatInt(count, 10))
	return (FtCreatePrefixCount)(c)
}

func (c FtCreateOnHash) Filter(filter interface{}) FtCreateFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtCreateFilter)(c)
}

func (c FtCreateOnHash) Language(defaultLang interface{}) FtCreateLanguage {
	c.cs = append(c.cs, "LANGUAGE", defaultLang)
	return (FtCreateLanguage)(c)
}

func (c FtCreateOnHash) LanguageField(langAttribute interface{}) FtCreateLanguageField {
	c.cs = append(c.cs, "LANGUAGE_FIELD", langAttribute)
	return (FtCreateLanguageField)(c)
}

func (c FtCreateOnHash) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreateOnHash) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateOnHash) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateOnHash) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateOnHash) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateOnHash) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateOnHash) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateOnHash) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateOnHash) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateOnHash) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateOnHash) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateOnHash) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateOnJson Completed

func (c FtCreateOnJson) Prefix(count int64) FtCreatePrefixCount {
	c.cs = append(c.cs, "PREFIX", strconv.FormatInt(count, 10))
	return (FtCreatePrefixCount)(c)
}

func (c FtCreateOnJson) Filter(filter interface{}) FtCreateFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtCreateFilter)(c)
}

func (c FtCreateOnJson) Language(defaultLang interface{}) FtCreateLanguage {
	c.cs = append(c.cs, "LANGUAGE", defaultLang)
	return (FtCreateLanguage)(c)
}

func (c FtCreateOnJson) LanguageField(langAttribute interface{}) FtCreateLanguageField {
	c.cs = append(c.cs, "LANGUAGE_FIELD", langAttribute)
	return (FtCreateLanguageField)(c)
}

func (c FtCreateOnJson) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreateOnJson) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateOnJson) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateOnJson) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateOnJson) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateOnJson) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateOnJson) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateOnJson) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateOnJson) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateOnJson) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateOnJson) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateOnJson) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreatePayloadField Completed

func (c FtCreatePayloadField) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreatePayloadField) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreatePayloadField) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreatePayloadField) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreatePayloadField) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreatePayloadField) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreatePayloadField) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreatePayloadField) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreatePayloadField) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreatePrefixCount Completed

func (c FtCreatePrefixCount) Prefix(prefix ...interface{}) FtCreatePrefixPrefix {
	c.cs = append(c.cs, prefix...)
	return (FtCreatePrefixPrefix)(c)
}

type FtCreatePrefixPrefix Completed

func (c FtCreatePrefixPrefix) Prefix(prefix ...interface{}) FtCreatePrefixPrefix {
	c.cs = append(c.cs, prefix...)
	return c
}

func (c FtCreatePrefixPrefix) Filter(filter interface{}) FtCreateFilter {
	c.cs = append(c.cs, "FILTER", filter)
	return (FtCreateFilter)(c)
}

func (c FtCreatePrefixPrefix) Language(defaultLang interface{}) FtCreateLanguage {
	c.cs = append(c.cs, "LANGUAGE", defaultLang)
	return (FtCreateLanguage)(c)
}

func (c FtCreatePrefixPrefix) LanguageField(langAttribute interface{}) FtCreateLanguageField {
	c.cs = append(c.cs, "LANGUAGE_FIELD", langAttribute)
	return (FtCreateLanguageField)(c)
}

func (c FtCreatePrefixPrefix) Score(defaultScore float64) FtCreateScore {
	c.cs = append(c.cs, "SCORE", strconv.FormatFloat(defaultScore, 'f', -1, 64))
	return (FtCreateScore)(c)
}

func (c FtCreatePrefixPrefix) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreatePrefixPrefix) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreatePrefixPrefix) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreatePrefixPrefix) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreatePrefixPrefix) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreatePrefixPrefix) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreatePrefixPrefix) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreatePrefixPrefix) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreatePrefixPrefix) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreatePrefixPrefix) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreatePrefixPrefix) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateSchema Completed

func (c FtCreateSchema) FieldName(fieldName interface{}) FtCreateFieldFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtCreateFieldFieldName)(c)
}

type FtCreateScore Completed

func (c FtCreateScore) ScoreField(scoreAttribute interface{}) FtCreateScoreField {
	c.cs = append(c.cs, "SCORE_FIELD", scoreAttribute)
	return (FtCreateScoreField)(c)
}

func (c FtCreateScore) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateScore) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateScore) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateScore) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateScore) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateScore) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateScore) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateScore) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateScore) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateScore) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateScoreField Completed

func (c FtCreateScoreField) PayloadField(payloadAttribute interface{}) FtCreatePayloadField {
	c.cs = append(c.cs, "PAYLOAD_FIELD", payloadAttribute)
	return (FtCreatePayloadField)(c)
}

func (c FtCreateScoreField) Maxtextfields() FtCreateMaxtextfields {
	c.cs = append(c.cs, "MAXTEXTFIELDS")
	return (FtCreateMaxtextfields)(c)
}

func (c FtCreateScoreField) Temporary(seconds float64) FtCreateTemporary {
	c.cs = append(c.cs, "TEMPORARY", strconv.FormatFloat(seconds, 'f', -1, 64))
	return (FtCreateTemporary)(c)
}

func (c FtCreateScoreField) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateScoreField) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateScoreField) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateScoreField) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateScoreField) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateScoreField) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateScoreField) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateSkipinitialscan Completed

func (c FtCreateSkipinitialscan) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateStopwordsStopword Completed

func (c FtCreateStopwordsStopword) Stopword(stopword ...interface{}) FtCreateStopwordsStopword {
	c.cs = append(c.cs, stopword...)
	return c
}

func (c FtCreateStopwordsStopword) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateStopwordsStopword) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateStopwordsStopwords Completed

func (c FtCreateStopwordsStopwords) Stopword(stopword ...interface{}) FtCreateStopwordsStopword {
	c.cs = append(c.cs, stopword...)
	return (FtCreateStopwordsStopword)(c)
}

func (c FtCreateStopwordsStopwords) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateStopwordsStopwords) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCreateTemporary Completed

func (c FtCreateTemporary) Nooffsets() FtCreateNooffsets {
	c.cs = append(c.cs, "NOOFFSETS")
	return (FtCreateNooffsets)(c)
}

func (c FtCreateTemporary) Nohl() FtCreateNohl {
	c.cs = append(c.cs, "NOHL")
	return (FtCreateNohl)(c)
}

func (c FtCreateTemporary) Nofields() FtCreateNofields {
	c.cs = append(c.cs, "NOFIELDS")
	return (FtCreateNofields)(c)
}

func (c FtCreateTemporary) Nofreqs() FtCreateNofreqs {
	c.cs = append(c.cs, "NOFREQS")
	return (FtCreateNofreqs)(c)
}

func (c FtCreateTemporary) Stopwords(count int64) FtCreateStopwordsStopwords {
	c.cs = append(c.cs, "STOPWORDS", strconv.FormatInt(count, 10))
	return (FtCreateStopwordsStopwords)(c)
}

func (c FtCreateTemporary) Skipinitialscan() FtCreateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtCreateSkipinitialscan)(c)
}

func (c FtCreateTemporary) Schema() FtCreateSchema {
	c.cs = append(c.cs, "SCHEMA")
	return (FtCreateSchema)(c)
}

type FtCursorDel Completed

func (b Builder) FtCursorDel() (c FtCursorDel) {
	c.cs = append(c.cs, "FT.CURSOR", "DEL")
	return c
}

func (c FtCursorDel) Index(index interface{}) FtCursorDelIndex {
	c.cs = append(c.cs, index)
	return (FtCursorDelIndex)(c)
}

type FtCursorDelCursorId Completed

func (c FtCursorDelCursorId) GetArgs() []interface{} {
	return c.cs
}

type FtCursorDelIndex Completed

func (c FtCursorDelIndex) CursorId(cursorId int64) FtCursorDelCursorId {
	c.cs = append(c.cs, strconv.FormatInt(cursorId, 10))
	return (FtCursorDelCursorId)(c)
}

type FtCursorRead Completed

func (b Builder) FtCursorRead() (c FtCursorRead) {
	c.cs = append(c.cs, "FT.CURSOR", "READ")
	return c
}

func (c FtCursorRead) Index(index interface{}) FtCursorReadIndex {
	c.cs = append(c.cs, index)
	return (FtCursorReadIndex)(c)
}

type FtCursorReadCount Completed

func (c FtCursorReadCount) GetArgs() []interface{} {
	return c.cs
}

type FtCursorReadCursorId Completed

func (c FtCursorReadCursorId) Count(readSize int64) FtCursorReadCount {
	c.cs = append(c.cs, "COUNT", strconv.FormatInt(readSize, 10))
	return (FtCursorReadCount)(c)
}

func (c FtCursorReadCursorId) GetArgs() []interface{} {
	return c.cs
}

type FtCursorReadIndex Completed

func (c FtCursorReadIndex) CursorId(cursorId int64) FtCursorReadCursorId {
	c.cs = append(c.cs, strconv.FormatInt(cursorId, 10))
	return (FtCursorReadCursorId)(c)
}

type FtDictadd Completed

func (b Builder) FtDictadd() (c FtDictadd) {
	c.cs = append(c.cs, "FT.DICTADD")
	return c
}

func (c FtDictadd) Dict(dict interface{}) FtDictaddDict {
	c.cs = append(c.cs, dict)
	return (FtDictaddDict)(c)
}

type FtDictaddDict Completed

func (c FtDictaddDict) Term(term ...interface{}) FtDictaddTerm {
	c.cs = append(c.cs, term...)
	return (FtDictaddTerm)(c)
}

type FtDictaddTerm Completed

func (c FtDictaddTerm) Term(term ...interface{}) FtDictaddTerm {
	c.cs = append(c.cs, term...)
	return c
}

func (c FtDictaddTerm) GetArgs() []interface{} {
	return c.cs
}

type FtDictdel Completed

func (b Builder) FtDictdel() (c FtDictdel) {
	c.cs = append(c.cs, "FT.DICTDEL")
	return c
}

func (c FtDictdel) Dict(dict interface{}) FtDictdelDict {
	c.cs = append(c.cs, dict)
	return (FtDictdelDict)(c)
}

type FtDictdelDict Completed

func (c FtDictdelDict) Term(term ...interface{}) FtDictdelTerm {
	c.cs = append(c.cs, term...)
	return (FtDictdelTerm)(c)
}

type FtDictdelTerm Completed

func (c FtDictdelTerm) Term(term ...interface{}) FtDictdelTerm {
	c.cs = append(c.cs, term...)
	return c
}

func (c FtDictdelTerm) GetArgs() []interface{} {
	return c.cs
}

type FtDictdump Completed

func (b Builder) FtDictdump() (c FtDictdump) {
	c.cs = append(c.cs, "FT.DICTDUMP")
	return c
}

func (c FtDictdump) Dict(dict interface{}) FtDictdumpDict {
	c.cs = append(c.cs, dict)
	return (FtDictdumpDict)(c)
}

type FtDictdumpDict Completed

func (c FtDictdumpDict) GetArgs() []interface{} {
	return c.cs
}

type FtDropindex Completed

func (b Builder) FtDropindex() (c FtDropindex) {
	c.cs = append(c.cs, "FT.DROPINDEX")
	return c
}

func (c FtDropindex) Index(index interface{}) FtDropindexIndex {
	c.cs = append(c.cs, index)
	return (FtDropindexIndex)(c)
}

type FtDropindexDeleteDocsDd Completed

func (c FtDropindexDeleteDocsDd) GetArgs() []interface{} {
	return c.cs
}

type FtDropindexIndex Completed

func (c FtDropindexIndex) Dd() FtDropindexDeleteDocsDd {
	c.cs = append(c.cs, "DD")
	return (FtDropindexDeleteDocsDd)(c)
}

func (c FtDropindexIndex) GetArgs() []interface{} {
	return c.cs
}

type FtExplain Completed

func (b Builder) FtExplain() (c FtExplain) {
	c.cs = append(c.cs, "FT.EXPLAIN")
	return c
}

func (c FtExplain) Index(index interface{}) FtExplainIndex {
	c.cs = append(c.cs, index)
	return (FtExplainIndex)(c)
}

type FtExplainDialect Completed

func (c FtExplainDialect) GetArgs() []interface{} {
	return c.cs
}

type FtExplainIndex Completed

func (c FtExplainIndex) Query(query interface{}) FtExplainQuery {
	c.cs = append(c.cs, query)
	return (FtExplainQuery)(c)
}

type FtExplainQuery Completed

func (c FtExplainQuery) Dialect(dialect int64) FtExplainDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtExplainDialect)(c)
}

func (c FtExplainQuery) GetArgs() []interface{} {
	return c.cs
}

type FtExplaincli Completed

func (b Builder) FtExplaincli() (c FtExplaincli) {
	c.cs = append(c.cs, "FT.EXPLAINCLI")
	return c
}

func (c FtExplaincli) Index(index interface{}) FtExplaincliIndex {
	c.cs = append(c.cs, index)
	return (FtExplaincliIndex)(c)
}

type FtExplaincliDialect Completed

func (c FtExplaincliDialect) GetArgs() []interface{} {
	return c.cs
}

type FtExplaincliIndex Completed

func (c FtExplaincliIndex) Query(query interface{}) FtExplaincliQuery {
	c.cs = append(c.cs, query)
	return (FtExplaincliQuery)(c)
}

type FtExplaincliQuery Completed

func (c FtExplaincliQuery) Dialect(dialect int64) FtExplaincliDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtExplaincliDialect)(c)
}

func (c FtExplaincliQuery) GetArgs() []interface{} {
	return c.cs
}

type FtInfo Completed

func (b Builder) FtInfo() (c FtInfo) {
	c.cs = append(c.cs, "FT.INFO")
	return c
}

func (c FtInfo) Index(index interface{}) FtInfoIndex {
	c.cs = append(c.cs, index)
	return (FtInfoIndex)(c)
}

type FtInfoIndex Completed

func (c FtInfoIndex) GetArgs() []interface{} {
	return c.cs
}

type FtList Completed

func (b Builder) FtList() (c FtList) {
	c.cs = append(c.cs, "FT._LIST")
	return c
}

func (c FtList) GetArgs() []interface{} {
	return c.cs
}

type FtProfile Completed

func (b Builder) FtProfile() (c FtProfile) {
	c.cs = append(c.cs, "FT.PROFILE")
	return c
}

func (c FtProfile) Index(index interface{}) FtProfileIndex {
	c.cs = append(c.cs, index)
	return (FtProfileIndex)(c)
}

type FtProfileIndex Completed

func (c FtProfileIndex) Search() FtProfileQuerytypeSearch {
	c.cs = append(c.cs, "SEARCH")
	return (FtProfileQuerytypeSearch)(c)
}

func (c FtProfileIndex) Aggregate() FtProfileQuerytypeAggregate {
	c.cs = append(c.cs, "AGGREGATE")
	return (FtProfileQuerytypeAggregate)(c)
}

type FtProfileLimited Completed

func (c FtProfileLimited) Query(query interface{}) FtProfileQuery {
	c.cs = append(c.cs, "QUERY", query)
	return (FtProfileQuery)(c)
}

type FtProfileQuery Completed

func (c FtProfileQuery) GetArgs() []interface{} {
	return c.cs
}

type FtProfileQuerytypeAggregate Completed

func (c FtProfileQuerytypeAggregate) Limited() FtProfileLimited {
	c.cs = append(c.cs, "LIMITED")
	return (FtProfileLimited)(c)
}

func (c FtProfileQuerytypeAggregate) Query(query interface{}) FtProfileQuery {
	c.cs = append(c.cs, "QUERY", query)
	return (FtProfileQuery)(c)
}

type FtProfileQuerytypeSearch Completed

func (c FtProfileQuerytypeSearch) Limited() FtProfileLimited {
	c.cs = append(c.cs, "LIMITED")
	return (FtProfileLimited)(c)
}

func (c FtProfileQuerytypeSearch) Query(query interface{}) FtProfileQuery {
	c.cs = append(c.cs, "QUERY", query)
	return (FtProfileQuery)(c)
}

type FtSearch Completed

func (b Builder) FtSearch() (c FtSearch) {
	c.cs = append(c.cs, "FT.SEARCH")
	return c
}

func (c FtSearch) Index(index interface{}) FtSearchIndex {
	c.cs = append(c.cs, index)
	return (FtSearchIndex)(c)
}

type FtSearchDialect Completed

func (c FtSearchDialect) GetArgs() []interface{} {
	return c.cs
}

type FtSearchExpander Completed

func (c FtSearchExpander) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchExpander) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchExpander) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchExpander) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchExpander) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchExpander) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchExpander) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchExpander) GetArgs() []interface{} {
	return c.cs
}

type FtSearchExplainscore Completed

func (c FtSearchExplainscore) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchExplainscore) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchExplainscore) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchExplainscore) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchExplainscore) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchExplainscore) GetArgs() []interface{} {
	return c.cs
}

type FtSearchFilterFilter Completed

func (c FtSearchFilterFilter) Min(min float64) FtSearchFilterMin {
	c.cs = append(c.cs, strconv.FormatFloat(min, 'f', -1, 64))
	return (FtSearchFilterMin)(c)
}

type FtSearchFilterMax Completed

func (c FtSearchFilterMax) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchFilterMax) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchFilterMax) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchFilterMax) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchFilterMax) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchFilterMax) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchFilterMax) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchFilterMax) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchFilterMax) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchFilterMax) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchFilterMax) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchFilterMax) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchFilterMax) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchFilterMax) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchFilterMax) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchFilterMax) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchFilterMax) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchFilterMax) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchFilterMax) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchFilterMax) GetArgs() []interface{} {
	return c.cs
}

type FtSearchFilterMin Completed

func (c FtSearchFilterMin) Max(max float64) FtSearchFilterMax {
	c.cs = append(c.cs, strconv.FormatFloat(max, 'f', -1, 64))
	return (FtSearchFilterMax)(c)
}

type FtSearchGeoFilterGeofilter Completed

func (c FtSearchGeoFilterGeofilter) Lon(lon float64) FtSearchGeoFilterLon {
	c.cs = append(c.cs, strconv.FormatFloat(lon, 'f', -1, 64))
	return (FtSearchGeoFilterLon)(c)
}

type FtSearchGeoFilterLat Completed

func (c FtSearchGeoFilterLat) Radius(radius float64) FtSearchGeoFilterRadius {
	c.cs = append(c.cs, strconv.FormatFloat(radius, 'f', -1, 64))
	return (FtSearchGeoFilterRadius)(c)
}

type FtSearchGeoFilterLon Completed

func (c FtSearchGeoFilterLon) Lat(lat float64) FtSearchGeoFilterLat {
	c.cs = append(c.cs, strconv.FormatFloat(lat, 'f', -1, 64))
	return (FtSearchGeoFilterLat)(c)
}

type FtSearchGeoFilterRadius Completed

func (c FtSearchGeoFilterRadius) M() FtSearchGeoFilterRadiusTypeM {
	c.cs = append(c.cs, "m")
	return (FtSearchGeoFilterRadiusTypeM)(c)
}

func (c FtSearchGeoFilterRadius) Km() FtSearchGeoFilterRadiusTypeKm {
	c.cs = append(c.cs, "km")
	return (FtSearchGeoFilterRadiusTypeKm)(c)
}

func (c FtSearchGeoFilterRadius) Mi() FtSearchGeoFilterRadiusTypeMi {
	c.cs = append(c.cs, "mi")
	return (FtSearchGeoFilterRadiusTypeMi)(c)
}

func (c FtSearchGeoFilterRadius) Ft() FtSearchGeoFilterRadiusTypeFt {
	c.cs = append(c.cs, "ft")
	return (FtSearchGeoFilterRadiusTypeFt)(c)
}

type FtSearchGeoFilterRadiusTypeFt Completed

func (c FtSearchGeoFilterRadiusTypeFt) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchGeoFilterRadiusTypeFt) GetArgs() []interface{} {
	return c.cs
}

type FtSearchGeoFilterRadiusTypeKm Completed

func (c FtSearchGeoFilterRadiusTypeKm) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchGeoFilterRadiusTypeKm) GetArgs() []interface{} {
	return c.cs
}

type FtSearchGeoFilterRadiusTypeM Completed

func (c FtSearchGeoFilterRadiusTypeM) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchGeoFilterRadiusTypeM) GetArgs() []interface{} {
	return c.cs
}

type FtSearchGeoFilterRadiusTypeMi Completed

func (c FtSearchGeoFilterRadiusTypeMi) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchGeoFilterRadiusTypeMi) GetArgs() []interface{} {
	return c.cs
}

type FtSearchHighlightFieldsField Completed

func (c FtSearchHighlightFieldsField) Field(field ...interface{}) FtSearchHighlightFieldsField {
	c.cs = append(c.cs, field...)
	return c
}

func (c FtSearchHighlightFieldsField) Tags() FtSearchHighlightTagsTags {
	c.cs = append(c.cs, "TAGS")
	return (FtSearchHighlightTagsTags)(c)
}

func (c FtSearchHighlightFieldsField) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchHighlightFieldsField) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchHighlightFieldsField) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchHighlightFieldsField) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchHighlightFieldsField) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchHighlightFieldsField) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchHighlightFieldsField) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchHighlightFieldsField) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchHighlightFieldsField) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchHighlightFieldsField) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchHighlightFieldsField) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchHighlightFieldsField) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchHighlightFieldsField) GetArgs() []interface{} {
	return c.cs
}

type FtSearchHighlightFieldsFields Completed

func (c FtSearchHighlightFieldsFields) Field(field ...interface{}) FtSearchHighlightFieldsField {
	c.cs = append(c.cs, field...)
	return (FtSearchHighlightFieldsField)(c)
}

type FtSearchHighlightHighlight Completed

func (c FtSearchHighlightHighlight) Fields(count interface{}) FtSearchHighlightFieldsFields {
	c.cs = append(c.cs, "FIELDS", count)
	return (FtSearchHighlightFieldsFields)(c)
}

func (c FtSearchHighlightHighlight) Tags() FtSearchHighlightTagsTags {
	c.cs = append(c.cs, "TAGS")
	return (FtSearchHighlightTagsTags)(c)
}

func (c FtSearchHighlightHighlight) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchHighlightHighlight) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchHighlightHighlight) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchHighlightHighlight) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchHighlightHighlight) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchHighlightHighlight) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchHighlightHighlight) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchHighlightHighlight) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchHighlightHighlight) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchHighlightHighlight) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchHighlightHighlight) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchHighlightHighlight) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchHighlightHighlight) GetArgs() []interface{} {
	return c.cs
}

type FtSearchHighlightTagsOpenClose Completed

func (c FtSearchHighlightTagsOpenClose) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchHighlightTagsOpenClose) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchHighlightTagsOpenClose) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchHighlightTagsOpenClose) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchHighlightTagsOpenClose) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchHighlightTagsOpenClose) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchHighlightTagsOpenClose) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchHighlightTagsOpenClose) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchHighlightTagsOpenClose) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchHighlightTagsOpenClose) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchHighlightTagsOpenClose) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchHighlightTagsOpenClose) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchHighlightTagsOpenClose) GetArgs() []interface{} {
	return c.cs
}

type FtSearchHighlightTagsTags Completed

func (c FtSearchHighlightTagsTags) OpenClose(open interface{}, close interface{}) FtSearchHighlightTagsOpenClose {
	c.cs = append(c.cs, open, close)
	return (FtSearchHighlightTagsOpenClose)(c)
}

type FtSearchInFieldsField Completed

func (c FtSearchInFieldsField) Field(field ...interface{}) FtSearchInFieldsField {
	c.cs = append(c.cs, field...)
	return c
}

func (c FtSearchInFieldsField) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchInFieldsField) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchInFieldsField) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchInFieldsField) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchInFieldsField) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchInFieldsField) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchInFieldsField) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchInFieldsField) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchInFieldsField) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchInFieldsField) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchInFieldsField) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchInFieldsField) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchInFieldsField) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchInFieldsField) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchInFieldsField) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchInFieldsField) GetArgs() []interface{} {
	return c.cs
}

type FtSearchInFieldsInfields Completed

func (c FtSearchInFieldsInfields) Field(field ...interface{}) FtSearchInFieldsField {
	c.cs = append(c.cs, field...)
	return (FtSearchInFieldsField)(c)
}

type FtSearchInKeysInkeys Completed

func (c FtSearchInKeysInkeys) Key(key ...interface{}) FtSearchInKeysKey {
	c.cs = append(c.cs, key...)
	return (FtSearchInKeysKey)(c)
}

type FtSearchInKeysKey Completed

func (c FtSearchInKeysKey) Key(key ...interface{}) FtSearchInKeysKey {
	c.cs = append(c.cs, key...)
	return c
}

func (c FtSearchInKeysKey) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchInKeysKey) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchInKeysKey) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchInKeysKey) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchInKeysKey) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchInKeysKey) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchInKeysKey) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchInKeysKey) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchInKeysKey) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchInKeysKey) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchInKeysKey) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchInKeysKey) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchInKeysKey) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchInKeysKey) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchInKeysKey) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchInKeysKey) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchInKeysKey) GetArgs() []interface{} {
	return c.cs
}

type FtSearchIndex Completed

func (c FtSearchIndex) Query(query interface{}) FtSearchQuery {
	c.cs = append(c.cs, query)
	return (FtSearchQuery)(c)
}

type FtSearchLanguage Completed

func (c FtSearchLanguage) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchLanguage) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchLanguage) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchLanguage) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchLanguage) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchLanguage) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchLanguage) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchLanguage) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchLanguage) GetArgs() []interface{} {
	return c.cs
}

type FtSearchLimitLimit Completed

func (c FtSearchLimitLimit) OffsetNum(offset int64, num int64) FtSearchLimitOffsetNum {
	c.cs = append(c.cs, strconv.FormatInt(offset, 10), strconv.FormatInt(num, 10))
	return (FtSearchLimitOffsetNum)(c)
}

type FtSearchLimitOffsetNum Completed

func (c FtSearchLimitOffsetNum) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchLimitOffsetNum) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchLimitOffsetNum) GetArgs() []interface{} {
	return c.cs
}

type FtSearchNocontent Completed

func (c FtSearchNocontent) Verbatim() FtSearchVerbatim {
	c.cs = append(c.cs, "VERBATIM")
	return (FtSearchVerbatim)(c)
}

func (c FtSearchNocontent) Nostopwords() FtSearchNostopwords {
	c.cs = append(c.cs, "NOSTOPWORDS")
	return (FtSearchNostopwords)(c)
}

func (c FtSearchNocontent) Withscores() FtSearchWithscores {
	c.cs = append(c.cs, "WITHSCORES")
	return (FtSearchWithscores)(c)
}

func (c FtSearchNocontent) Withpayloads() FtSearchWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSearchWithpayloads)(c)
}

func (c FtSearchNocontent) Withsortkeys() FtSearchWithsortkeys {
	c.cs = append(c.cs, "WITHSORTKEYS")
	return (FtSearchWithsortkeys)(c)
}

func (c FtSearchNocontent) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchNocontent) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchNocontent) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchNocontent) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchNocontent) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchNocontent) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchNocontent) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchNocontent) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchNocontent) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchNocontent) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchNocontent) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchNocontent) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchNocontent) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchNocontent) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchNocontent) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchNocontent) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchNocontent) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchNocontent) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchNocontent) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchNocontent) GetArgs() []interface{} {
	return c.cs
}

type FtSearchNostopwords Completed

func (c FtSearchNostopwords) Withscores() FtSearchWithscores {
	c.cs = append(c.cs, "WITHSCORES")
	return (FtSearchWithscores)(c)
}

func (c FtSearchNostopwords) Withpayloads() FtSearchWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSearchWithpayloads)(c)
}

func (c FtSearchNostopwords) Withsortkeys() FtSearchWithsortkeys {
	c.cs = append(c.cs, "WITHSORTKEYS")
	return (FtSearchWithsortkeys)(c)
}

func (c FtSearchNostopwords) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchNostopwords) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchNostopwords) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchNostopwords) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchNostopwords) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchNostopwords) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchNostopwords) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchNostopwords) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchNostopwords) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchNostopwords) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchNostopwords) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchNostopwords) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchNostopwords) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchNostopwords) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchNostopwords) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchNostopwords) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchNostopwords) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchNostopwords) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchNostopwords) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchNostopwords) GetArgs() []interface{} {
	return c.cs
}

type FtSearchParamsNameValue Completed

func (c FtSearchParamsNameValue) NameValue(name interface{}, value interface{}) FtSearchParamsNameValue {
	c.cs = append(c.cs, name, value)
	return c
}

func (c FtSearchParamsNameValue) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchParamsNameValue) GetArgs() []interface{} {
	return c.cs
}

type FtSearchParamsNargs Completed

func (c FtSearchParamsNargs) NameValue() FtSearchParamsNameValue {
	return (FtSearchParamsNameValue)(c)
}

type FtSearchParamsParams Completed

func (c FtSearchParamsParams) Nargs(nargs int64) FtSearchParamsNargs {
	c.cs = append(c.cs, strconv.FormatInt(nargs, 10))
	return (FtSearchParamsNargs)(c)
}

type FtSearchPayload Completed

func (c FtSearchPayload) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchPayload) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchPayload) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchPayload) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchPayload) GetArgs() []interface{} {
	return c.cs
}

type FtSearchQuery Completed

func (c FtSearchQuery) Nocontent() FtSearchNocontent {
	c.cs = append(c.cs, "NOCONTENT")
	return (FtSearchNocontent)(c)
}

func (c FtSearchQuery) Verbatim() FtSearchVerbatim {
	c.cs = append(c.cs, "VERBATIM")
	return (FtSearchVerbatim)(c)
}

func (c FtSearchQuery) Nostopwords() FtSearchNostopwords {
	c.cs = append(c.cs, "NOSTOPWORDS")
	return (FtSearchNostopwords)(c)
}

func (c FtSearchQuery) Withscores() FtSearchWithscores {
	c.cs = append(c.cs, "WITHSCORES")
	return (FtSearchWithscores)(c)
}

func (c FtSearchQuery) Withpayloads() FtSearchWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSearchWithpayloads)(c)
}

func (c FtSearchQuery) Withsortkeys() FtSearchWithsortkeys {
	c.cs = append(c.cs, "WITHSORTKEYS")
	return (FtSearchWithsortkeys)(c)
}

func (c FtSearchQuery) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchQuery) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchQuery) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchQuery) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchQuery) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchQuery) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchQuery) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchQuery) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchQuery) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchQuery) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchQuery) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchQuery) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchQuery) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchQuery) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchQuery) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchQuery) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchQuery) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchQuery) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchQuery) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchQuery) GetArgs() []interface{} {
	return c.cs
}

type FtSearchReturnIdentifiersAs Completed

func (c FtSearchReturnIdentifiersAs) Identifier(identifier interface{}) FtSearchReturnIdentifiersIdentifier {
	c.cs = append(c.cs, identifier)
	return (FtSearchReturnIdentifiersIdentifier)(c)
}

func (c FtSearchReturnIdentifiersAs) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchReturnIdentifiersAs) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchReturnIdentifiersAs) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchReturnIdentifiersAs) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchReturnIdentifiersAs) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchReturnIdentifiersAs) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchReturnIdentifiersAs) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchReturnIdentifiersAs) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchReturnIdentifiersAs) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchReturnIdentifiersAs) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchReturnIdentifiersAs) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchReturnIdentifiersAs) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchReturnIdentifiersAs) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchReturnIdentifiersAs) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchReturnIdentifiersAs) GetArgs() []interface{} {
	return c.cs
}

type FtSearchReturnIdentifiersIdentifier Completed

func (c FtSearchReturnIdentifiersIdentifier) As(property interface{}) FtSearchReturnIdentifiersAs {
	c.cs = append(c.cs, "AS", property)
	return (FtSearchReturnIdentifiersAs)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Identifier(identifier interface{}) FtSearchReturnIdentifiersIdentifier {
	c.cs = append(c.cs, identifier)
	return c
}

func (c FtSearchReturnIdentifiersIdentifier) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchReturnIdentifiersIdentifier) GetArgs() []interface{} {
	return c.cs
}

type FtSearchReturnReturn Completed

func (c FtSearchReturnReturn) Identifier(identifier interface{}) FtSearchReturnIdentifiersIdentifier {
	c.cs = append(c.cs, identifier)
	return (FtSearchReturnIdentifiersIdentifier)(c)
}

type FtSearchScorer Completed

func (c FtSearchScorer) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchScorer) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchScorer) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchScorer) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchScorer) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchScorer) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchScorer) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSlop Completed

func (c FtSearchSlop) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchSlop) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchSlop) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchSlop) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchSlop) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchSlop) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchSlop) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchSlop) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchSlop) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSlop) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSlop) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSlop) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSortbyOrderAsc Completed

func (c FtSearchSortbyOrderAsc) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSortbyOrderAsc) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSortbyOrderAsc) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSortbyOrderAsc) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSortbyOrderDesc Completed

func (c FtSearchSortbyOrderDesc) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSortbyOrderDesc) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSortbyOrderDesc) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSortbyOrderDesc) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSortbySortby Completed

func (c FtSearchSortbySortby) Asc() FtSearchSortbyOrderAsc {
	c.cs = append(c.cs, "ASC")
	return (FtSearchSortbyOrderAsc)(c)
}

func (c FtSearchSortbySortby) Desc() FtSearchSortbyOrderDesc {
	c.cs = append(c.cs, "DESC")
	return (FtSearchSortbyOrderDesc)(c)
}

func (c FtSearchSortbySortby) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSortbySortby) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSortbySortby) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSortbySortby) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSummarizeFieldsField Completed

func (c FtSearchSummarizeFieldsField) Field(field ...interface{}) FtSearchSummarizeFieldsField {
	c.cs = append(c.cs, field...)
	return c
}

func (c FtSearchSummarizeFieldsField) Frags(num int64) FtSearchSummarizeFrags {
	c.cs = append(c.cs, "FRAGS", strconv.FormatInt(num, 10))
	return (FtSearchSummarizeFrags)(c)
}

func (c FtSearchSummarizeFieldsField) Len(fragsize int64) FtSearchSummarizeLen {
	c.cs = append(c.cs, "LEN", strconv.FormatInt(fragsize, 10))
	return (FtSearchSummarizeLen)(c)
}

func (c FtSearchSummarizeFieldsField) Separator(separator interface{}) FtSearchSummarizeSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtSearchSummarizeSeparator)(c)
}

func (c FtSearchSummarizeFieldsField) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchSummarizeFieldsField) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchSummarizeFieldsField) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchSummarizeFieldsField) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchSummarizeFieldsField) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchSummarizeFieldsField) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchSummarizeFieldsField) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchSummarizeFieldsField) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchSummarizeFieldsField) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchSummarizeFieldsField) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchSummarizeFieldsField) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSummarizeFieldsField) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSummarizeFieldsField) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSummarizeFieldsField) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSummarizeFieldsFields Completed

func (c FtSearchSummarizeFieldsFields) Field(field ...interface{}) FtSearchSummarizeFieldsField {
	c.cs = append(c.cs, field...)
	return (FtSearchSummarizeFieldsField)(c)
}

type FtSearchSummarizeFrags Completed

func (c FtSearchSummarizeFrags) Len(fragsize int64) FtSearchSummarizeLen {
	c.cs = append(c.cs, "LEN", strconv.FormatInt(fragsize, 10))
	return (FtSearchSummarizeLen)(c)
}

func (c FtSearchSummarizeFrags) Separator(separator interface{}) FtSearchSummarizeSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtSearchSummarizeSeparator)(c)
}

func (c FtSearchSummarizeFrags) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchSummarizeFrags) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchSummarizeFrags) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchSummarizeFrags) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchSummarizeFrags) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchSummarizeFrags) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchSummarizeFrags) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchSummarizeFrags) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchSummarizeFrags) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchSummarizeFrags) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchSummarizeFrags) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSummarizeFrags) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSummarizeFrags) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSummarizeFrags) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSummarizeLen Completed

func (c FtSearchSummarizeLen) Separator(separator interface{}) FtSearchSummarizeSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtSearchSummarizeSeparator)(c)
}

func (c FtSearchSummarizeLen) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchSummarizeLen) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchSummarizeLen) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchSummarizeLen) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchSummarizeLen) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchSummarizeLen) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchSummarizeLen) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchSummarizeLen) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchSummarizeLen) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchSummarizeLen) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchSummarizeLen) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSummarizeLen) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSummarizeLen) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSummarizeLen) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSummarizeSeparator Completed

func (c FtSearchSummarizeSeparator) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchSummarizeSeparator) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchSummarizeSeparator) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchSummarizeSeparator) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchSummarizeSeparator) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchSummarizeSeparator) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchSummarizeSeparator) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchSummarizeSeparator) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchSummarizeSeparator) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchSummarizeSeparator) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchSummarizeSeparator) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSummarizeSeparator) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSummarizeSeparator) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSummarizeSeparator) GetArgs() []interface{} {
	return c.cs
}

type FtSearchSummarizeSummarize Completed

func (c FtSearchSummarizeSummarize) Fields(count interface{}) FtSearchSummarizeFieldsFields {
	c.cs = append(c.cs, "FIELDS", count)
	return (FtSearchSummarizeFieldsFields)(c)
}

func (c FtSearchSummarizeSummarize) Frags(num int64) FtSearchSummarizeFrags {
	c.cs = append(c.cs, "FRAGS", strconv.FormatInt(num, 10))
	return (FtSearchSummarizeFrags)(c)
}

func (c FtSearchSummarizeSummarize) Len(fragsize int64) FtSearchSummarizeLen {
	c.cs = append(c.cs, "LEN", strconv.FormatInt(fragsize, 10))
	return (FtSearchSummarizeLen)(c)
}

func (c FtSearchSummarizeSummarize) Separator(separator interface{}) FtSearchSummarizeSeparator {
	c.cs = append(c.cs, "SEPARATOR", separator)
	return (FtSearchSummarizeSeparator)(c)
}

func (c FtSearchSummarizeSummarize) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchSummarizeSummarize) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchSummarizeSummarize) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchSummarizeSummarize) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchSummarizeSummarize) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchSummarizeSummarize) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchSummarizeSummarize) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchSummarizeSummarize) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchSummarizeSummarize) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchSummarizeSummarize) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchSummarizeSummarize) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchSummarizeSummarize) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchSummarizeSummarize) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchSummarizeSummarize) GetArgs() []interface{} {
	return c.cs
}

type FtSearchTagsInorder Completed

func (c FtSearchTagsInorder) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchTagsInorder) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchTagsInorder) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchTagsInorder) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchTagsInorder) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchTagsInorder) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchTagsInorder) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchTagsInorder) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchTagsInorder) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchTagsInorder) GetArgs() []interface{} {
	return c.cs
}

type FtSearchTimeout Completed

func (c FtSearchTimeout) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchTimeout) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchTimeout) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchTimeout) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchTimeout) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchTimeout) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchTimeout) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchTimeout) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchTimeout) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchTimeout) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchTimeout) GetArgs() []interface{} {
	return c.cs
}

type FtSearchVerbatim Completed

func (c FtSearchVerbatim) Nostopwords() FtSearchNostopwords {
	c.cs = append(c.cs, "NOSTOPWORDS")
	return (FtSearchNostopwords)(c)
}

func (c FtSearchVerbatim) Withscores() FtSearchWithscores {
	c.cs = append(c.cs, "WITHSCORES")
	return (FtSearchWithscores)(c)
}

func (c FtSearchVerbatim) Withpayloads() FtSearchWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSearchWithpayloads)(c)
}

func (c FtSearchVerbatim) Withsortkeys() FtSearchWithsortkeys {
	c.cs = append(c.cs, "WITHSORTKEYS")
	return (FtSearchWithsortkeys)(c)
}

func (c FtSearchVerbatim) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchVerbatim) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchVerbatim) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchVerbatim) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchVerbatim) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchVerbatim) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchVerbatim) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchVerbatim) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchVerbatim) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchVerbatim) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchVerbatim) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchVerbatim) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchVerbatim) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchVerbatim) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchVerbatim) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchVerbatim) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchVerbatim) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchVerbatim) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchVerbatim) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchVerbatim) GetArgs() []interface{} {
	return c.cs
}

type FtSearchWithpayloads Completed

func (c FtSearchWithpayloads) Withsortkeys() FtSearchWithsortkeys {
	c.cs = append(c.cs, "WITHSORTKEYS")
	return (FtSearchWithsortkeys)(c)
}

func (c FtSearchWithpayloads) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchWithpayloads) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchWithpayloads) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchWithpayloads) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchWithpayloads) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchWithpayloads) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchWithpayloads) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchWithpayloads) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchWithpayloads) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchWithpayloads) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchWithpayloads) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchWithpayloads) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchWithpayloads) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchWithpayloads) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchWithpayloads) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchWithpayloads) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchWithpayloads) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchWithpayloads) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchWithpayloads) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchWithpayloads) GetArgs() []interface{} {
	return c.cs
}

type FtSearchWithscores Completed

func (c FtSearchWithscores) Withpayloads() FtSearchWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSearchWithpayloads)(c)
}

func (c FtSearchWithscores) Withsortkeys() FtSearchWithsortkeys {
	c.cs = append(c.cs, "WITHSORTKEYS")
	return (FtSearchWithsortkeys)(c)
}

func (c FtSearchWithscores) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchWithscores) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchWithscores) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchWithscores) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchWithscores) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchWithscores) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchWithscores) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchWithscores) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchWithscores) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchWithscores) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchWithscores) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchWithscores) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchWithscores) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchWithscores) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchWithscores) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchWithscores) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchWithscores) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchWithscores) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchWithscores) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchWithscores) GetArgs() []interface{} {
	return c.cs
}

type FtSearchWithsortkeys Completed

func (c FtSearchWithsortkeys) Filter(numericField interface{}) FtSearchFilterFilter {
	c.cs = append(c.cs, "FILTER", numericField)
	return (FtSearchFilterFilter)(c)
}

func (c FtSearchWithsortkeys) Geofilter(geoField interface{}) FtSearchGeoFilterGeofilter {
	c.cs = append(c.cs, "GEOFILTER", geoField)
	return (FtSearchGeoFilterGeofilter)(c)
}

func (c FtSearchWithsortkeys) Inkeys(count interface{}) FtSearchInKeysInkeys {
	c.cs = append(c.cs, "INKEYS", count)
	return (FtSearchInKeysInkeys)(c)
}

func (c FtSearchWithsortkeys) Infields(count interface{}) FtSearchInFieldsInfields {
	c.cs = append(c.cs, "INFIELDS", count)
	return (FtSearchInFieldsInfields)(c)
}

func (c FtSearchWithsortkeys) Return(count interface{}) FtSearchReturnReturn {
	c.cs = append(c.cs, "RETURN", count)
	return (FtSearchReturnReturn)(c)
}

func (c FtSearchWithsortkeys) Summarize() FtSearchSummarizeSummarize {
	c.cs = append(c.cs, "SUMMARIZE")
	return (FtSearchSummarizeSummarize)(c)
}

func (c FtSearchWithsortkeys) Highlight() FtSearchHighlightHighlight {
	c.cs = append(c.cs, "HIGHLIGHT")
	return (FtSearchHighlightHighlight)(c)
}

func (c FtSearchWithsortkeys) Slop(slop int64) FtSearchSlop {
	c.cs = append(c.cs, "SLOP", strconv.FormatInt(slop, 10))
	return (FtSearchSlop)(c)
}

func (c FtSearchWithsortkeys) Timeout(timeout int64) FtSearchTimeout {
	c.cs = append(c.cs, "TIMEOUT", strconv.FormatInt(timeout, 10))
	return (FtSearchTimeout)(c)
}

func (c FtSearchWithsortkeys) Inorder() FtSearchTagsInorder {
	c.cs = append(c.cs, "INORDER")
	return (FtSearchTagsInorder)(c)
}

func (c FtSearchWithsortkeys) Language(language interface{}) FtSearchLanguage {
	c.cs = append(c.cs, "LANGUAGE", language)
	return (FtSearchLanguage)(c)
}

func (c FtSearchWithsortkeys) Expander(expander interface{}) FtSearchExpander {
	c.cs = append(c.cs, "EXPANDER", expander)
	return (FtSearchExpander)(c)
}

func (c FtSearchWithsortkeys) Scorer(scorer interface{}) FtSearchScorer {
	c.cs = append(c.cs, "SCORER", scorer)
	return (FtSearchScorer)(c)
}

func (c FtSearchWithsortkeys) Explainscore() FtSearchExplainscore {
	c.cs = append(c.cs, "EXPLAINSCORE")
	return (FtSearchExplainscore)(c)
}

func (c FtSearchWithsortkeys) Payload(payload interface{}) FtSearchPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSearchPayload)(c)
}

func (c FtSearchWithsortkeys) Sortby(sortby interface{}) FtSearchSortbySortby {
	c.cs = append(c.cs, "SORTBY", sortby)
	return (FtSearchSortbySortby)(c)
}

func (c FtSearchWithsortkeys) Limit() FtSearchLimitLimit {
	c.cs = append(c.cs, "LIMIT")
	return (FtSearchLimitLimit)(c)
}

func (c FtSearchWithsortkeys) Params() FtSearchParamsParams {
	c.cs = append(c.cs, "PARAMS")
	return (FtSearchParamsParams)(c)
}

func (c FtSearchWithsortkeys) Dialect(dialect int64) FtSearchDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSearchDialect)(c)
}

func (c FtSearchWithsortkeys) GetArgs() []interface{} {
	return c.cs
}

type FtSpellcheck Completed

func (b Builder) FtSpellcheck() (c FtSpellcheck) {
	c.cs = append(c.cs, "FT.SPELLCHECK")
	return c
}

func (c FtSpellcheck) Index(index interface{}) FtSpellcheckIndex {
	c.cs = append(c.cs, index)
	return (FtSpellcheckIndex)(c)
}

type FtSpellcheckDialect Completed

func (c FtSpellcheckDialect) GetArgs() []interface{} {
	return c.cs
}

type FtSpellcheckDistance Completed

func (c FtSpellcheckDistance) TermsInclude() FtSpellcheckTermsTermsInclude {
	c.cs = append(c.cs, "TERMS", "INCLUDE")
	return (FtSpellcheckTermsTermsInclude)(c)
}

func (c FtSpellcheckDistance) TermsExclude() FtSpellcheckTermsTermsExclude {
	c.cs = append(c.cs, "TERMS", "EXCLUDE")
	return (FtSpellcheckTermsTermsExclude)(c)
}

func (c FtSpellcheckDistance) Dialect(dialect int64) FtSpellcheckDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSpellcheckDialect)(c)
}

func (c FtSpellcheckDistance) GetArgs() []interface{} {
	return c.cs
}

type FtSpellcheckIndex Completed

func (c FtSpellcheckIndex) Query(query interface{}) FtSpellcheckQuery {
	c.cs = append(c.cs, query)
	return (FtSpellcheckQuery)(c)
}

type FtSpellcheckQuery Completed

func (c FtSpellcheckQuery) Distance(distance int64) FtSpellcheckDistance {
	c.cs = append(c.cs, "DISTANCE", strconv.FormatInt(distance, 10))
	return (FtSpellcheckDistance)(c)
}

func (c FtSpellcheckQuery) TermsInclude() FtSpellcheckTermsTermsInclude {
	c.cs = append(c.cs, "TERMS", "INCLUDE")
	return (FtSpellcheckTermsTermsInclude)(c)
}

func (c FtSpellcheckQuery) TermsExclude() FtSpellcheckTermsTermsExclude {
	c.cs = append(c.cs, "TERMS", "EXCLUDE")
	return (FtSpellcheckTermsTermsExclude)(c)
}

func (c FtSpellcheckQuery) Dialect(dialect int64) FtSpellcheckDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSpellcheckDialect)(c)
}

func (c FtSpellcheckQuery) GetArgs() []interface{} {
	return c.cs
}

type FtSpellcheckTermsDictionary Completed

func (c FtSpellcheckTermsDictionary) Terms(terms ...interface{}) FtSpellcheckTermsTerms {
	c.cs = append(c.cs, terms...)
	return (FtSpellcheckTermsTerms)(c)
}

func (c FtSpellcheckTermsDictionary) Dialect(dialect int64) FtSpellcheckDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSpellcheckDialect)(c)
}

func (c FtSpellcheckTermsDictionary) GetArgs() []interface{} {
	return c.cs
}

type FtSpellcheckTermsTerms Completed

func (c FtSpellcheckTermsTerms) Terms(terms ...interface{}) FtSpellcheckTermsTerms {
	c.cs = append(c.cs, terms...)
	return c
}

func (c FtSpellcheckTermsTerms) Dialect(dialect int64) FtSpellcheckDialect {
	c.cs = append(c.cs, "DIALECT", strconv.FormatInt(dialect, 10))
	return (FtSpellcheckDialect)(c)
}

func (c FtSpellcheckTermsTerms) GetArgs() []interface{} {
	return c.cs
}

type FtSpellcheckTermsTermsExclude Completed

func (c FtSpellcheckTermsTermsExclude) Dictionary(dictionary interface{}) FtSpellcheckTermsDictionary {
	c.cs = append(c.cs, dictionary)
	return (FtSpellcheckTermsDictionary)(c)
}

type FtSpellcheckTermsTermsInclude Completed

func (c FtSpellcheckTermsTermsInclude) Dictionary(dictionary interface{}) FtSpellcheckTermsDictionary {
	c.cs = append(c.cs, dictionary)
	return (FtSpellcheckTermsDictionary)(c)
}

type FtSyndump Completed

func (b Builder) FtSyndump() (c FtSyndump) {
	c.cs = append(c.cs, "FT.SYNDUMP")
	return c
}

func (c FtSyndump) Index(index interface{}) FtSyndumpIndex {
	c.cs = append(c.cs, index)
	return (FtSyndumpIndex)(c)
}

type FtSyndumpIndex Completed

func (c FtSyndumpIndex) GetArgs() []interface{} {
	return c.cs
}

type FtSynupdate Completed

func (b Builder) FtSynupdate() (c FtSynupdate) {
	c.cs = append(c.cs, "FT.SYNUPDATE")
	return c
}

func (c FtSynupdate) Index(index interface{}) FtSynupdateIndex {
	c.cs = append(c.cs, index)
	return (FtSynupdateIndex)(c)
}

type FtSynupdateIndex Completed

func (c FtSynupdateIndex) SynonymGroupId(synonymGroupId interface{}) FtSynupdateSynonymGroupId {
	c.cs = append(c.cs, synonymGroupId)
	return (FtSynupdateSynonymGroupId)(c)
}

type FtSynupdateSkipinitialscan Completed

func (c FtSynupdateSkipinitialscan) Term(term ...interface{}) FtSynupdateTerm {
	c.cs = append(c.cs, term...)
	return (FtSynupdateTerm)(c)
}

type FtSynupdateSynonymGroupId Completed

func (c FtSynupdateSynonymGroupId) Skipinitialscan() FtSynupdateSkipinitialscan {
	c.cs = append(c.cs, "SKIPINITIALSCAN")
	return (FtSynupdateSkipinitialscan)(c)
}

func (c FtSynupdateSynonymGroupId) Term(term ...interface{}) FtSynupdateTerm {
	c.cs = append(c.cs, term...)
	return (FtSynupdateTerm)(c)
}

type FtSynupdateTerm Completed

func (c FtSynupdateTerm) Term(term ...interface{}) FtSynupdateTerm {
	c.cs = append(c.cs, term...)
	return c
}

func (c FtSynupdateTerm) GetArgs() []interface{} {
	return c.cs
}

type FtTagvals Completed

func (b Builder) FtTagvals() (c FtTagvals) {
	c.cs = append(c.cs, "FT.TAGVALS")
	return c
}

func (c FtTagvals) Index(index interface{}) FtTagvalsIndex {
	c.cs = append(c.cs, index)
	return (FtTagvalsIndex)(c)
}

type FtTagvalsFieldName Completed

func (c FtTagvalsFieldName) GetArgs() []interface{} {
	return c.cs
}

type FtTagvalsIndex Completed

func (c FtTagvalsIndex) FieldName(fieldName interface{}) FtTagvalsFieldName {
	c.cs = append(c.cs, fieldName)
	return (FtTagvalsFieldName)(c)
}

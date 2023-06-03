// Code generated DO NOT EDIT

package cmds

import "strconv"

type JsonArrappend Completed

func (b Builder) JsonArrappend() (c JsonArrappend) {
	c.cs = append(c.cs, "JSON.ARRAPPEND")
	return c
}

func (c JsonArrappend) Key(key interface{}) JsonArrappendKey {
	c.cs = append(c.cs, key)
	return (JsonArrappendKey)(c)
}

type JsonArrappendKey Completed

func (c JsonArrappendKey) Path(path interface{}) JsonArrappendPath {
	c.cs = append(c.cs, path)
	return (JsonArrappendPath)(c)
}

func (c JsonArrappendKey) Value(value ...interface{}) JsonArrappendValue {
	c.cs = append(c.cs, value...)
	return (JsonArrappendValue)(c)
}

type JsonArrappendPath Completed

func (c JsonArrappendPath) Value(value ...interface{}) JsonArrappendValue {
	c.cs = append(c.cs, value...)
	return (JsonArrappendValue)(c)
}

type JsonArrappendValue Completed

func (c JsonArrappendValue) Value(value ...interface{}) JsonArrappendValue {
	c.cs = append(c.cs, value...)
	return c
}

func (c JsonArrappendValue) GetArgs() []interface{} {
	return c.cs
}

type JsonArrindex Completed

func (b Builder) JsonArrindex() (c JsonArrindex) {
	c.cs = append(c.cs, "JSON.ARRINDEX")
	return c
}

func (c JsonArrindex) Key(key interface{}) JsonArrindexKey {
	c.cs = append(c.cs, key)
	return (JsonArrindexKey)(c)
}

type JsonArrindexKey Completed

func (c JsonArrindexKey) Path(path interface{}) JsonArrindexPath {
	c.cs = append(c.cs, path)
	return (JsonArrindexPath)(c)
}

type JsonArrindexPath Completed

func (c JsonArrindexPath) Value(value interface{}) JsonArrindexValue {
	c.cs = append(c.cs, value)
	return (JsonArrindexValue)(c)
}

type JsonArrindexStartStart Completed

func (c JsonArrindexStartStart) Stop(stop int64) JsonArrindexStartStop {
	c.cs = append(c.cs, strconv.FormatInt(stop, 10))
	return (JsonArrindexStartStop)(c)
}

func (c JsonArrindexStartStart) GetArgs() []interface{} {
	return c.cs
}

type JsonArrindexStartStop Completed

func (c JsonArrindexStartStop) GetArgs() []interface{} {
	return c.cs
}

type JsonArrindexValue Completed

func (c JsonArrindexValue) Start(start int64) JsonArrindexStartStart {
	c.cs = append(c.cs, strconv.FormatInt(start, 10))
	return (JsonArrindexStartStart)(c)
}

func (c JsonArrindexValue) GetArgs() []interface{} {
	return c.cs
}

type JsonArrinsert Completed

func (b Builder) JsonArrinsert() (c JsonArrinsert) {
	c.cs = append(c.cs, "JSON.ARRINSERT")
	return c
}

func (c JsonArrinsert) Key(key interface{}) JsonArrinsertKey {
	c.cs = append(c.cs, key)
	return (JsonArrinsertKey)(c)
}

type JsonArrinsertIndex Completed

func (c JsonArrinsertIndex) Value(value ...interface{}) JsonArrinsertValue {
	c.cs = append(c.cs, value...)
	return (JsonArrinsertValue)(c)
}

type JsonArrinsertKey Completed

func (c JsonArrinsertKey) Path(path interface{}) JsonArrinsertPath {
	c.cs = append(c.cs, path)
	return (JsonArrinsertPath)(c)
}

type JsonArrinsertPath Completed

func (c JsonArrinsertPath) Index(index int64) JsonArrinsertIndex {
	c.cs = append(c.cs, strconv.FormatInt(index, 10))
	return (JsonArrinsertIndex)(c)
}

type JsonArrinsertValue Completed

func (c JsonArrinsertValue) Value(value ...interface{}) JsonArrinsertValue {
	c.cs = append(c.cs, value...)
	return c
}

func (c JsonArrinsertValue) GetArgs() []interface{} {
	return c.cs
}

type JsonArrlen Completed

func (b Builder) JsonArrlen() (c JsonArrlen) {
	c.cs = append(c.cs, "JSON.ARRLEN")
	return c
}

func (c JsonArrlen) Key(key interface{}) JsonArrlenKey {
	c.cs = append(c.cs, key)
	return (JsonArrlenKey)(c)
}

type JsonArrlenKey Completed

func (c JsonArrlenKey) Path(path interface{}) JsonArrlenPath {
	c.cs = append(c.cs, path)
	return (JsonArrlenPath)(c)
}

func (c JsonArrlenKey) GetArgs() []interface{} {
	return c.cs
}

type JsonArrlenPath Completed

func (c JsonArrlenPath) GetArgs() []interface{} {
	return c.cs
}

type JsonArrpop Completed

func (b Builder) JsonArrpop() (c JsonArrpop) {
	c.cs = append(c.cs, "JSON.ARRPOP")
	return c
}

func (c JsonArrpop) Key(key interface{}) JsonArrpopKey {
	c.cs = append(c.cs, key)
	return (JsonArrpopKey)(c)
}

type JsonArrpopKey Completed

func (c JsonArrpopKey) Path(path interface{}) JsonArrpopPathPath {
	c.cs = append(c.cs, path)
	return (JsonArrpopPathPath)(c)
}

func (c JsonArrpopKey) GetArgs() []interface{} {
	return c.cs
}

type JsonArrpopPathIndex Completed

func (c JsonArrpopPathIndex) GetArgs() []interface{} {
	return c.cs
}

type JsonArrpopPathPath Completed

func (c JsonArrpopPathPath) Index(index int64) JsonArrpopPathIndex {
	c.cs = append(c.cs, strconv.FormatInt(index, 10))
	return (JsonArrpopPathIndex)(c)
}

func (c JsonArrpopPathPath) GetArgs() []interface{} {
	return c.cs
}

type JsonArrtrim Completed

func (b Builder) JsonArrtrim() (c JsonArrtrim) {
	c.cs = append(c.cs, "JSON.ARRTRIM")
	return c
}

func (c JsonArrtrim) Key(key interface{}) JsonArrtrimKey {
	c.cs = append(c.cs, key)
	return (JsonArrtrimKey)(c)
}

type JsonArrtrimKey Completed

func (c JsonArrtrimKey) Path(path interface{}) JsonArrtrimPath {
	c.cs = append(c.cs, path)
	return (JsonArrtrimPath)(c)
}

type JsonArrtrimPath Completed

func (c JsonArrtrimPath) Start(start int64) JsonArrtrimStart {
	c.cs = append(c.cs, strconv.FormatInt(start, 10))
	return (JsonArrtrimStart)(c)
}

type JsonArrtrimStart Completed

func (c JsonArrtrimStart) Stop(stop int64) JsonArrtrimStop {
	c.cs = append(c.cs, strconv.FormatInt(stop, 10))
	return (JsonArrtrimStop)(c)
}

type JsonArrtrimStop Completed

func (c JsonArrtrimStop) GetArgs() []interface{} {
	return c.cs
}

type JsonClear Completed

func (b Builder) JsonClear() (c JsonClear) {
	c.cs = append(c.cs, "JSON.CLEAR")
	return c
}

func (c JsonClear) Key(key interface{}) JsonClearKey {
	c.cs = append(c.cs, key)
	return (JsonClearKey)(c)
}

type JsonClearKey Completed

func (c JsonClearKey) Path(path interface{}) JsonClearPath {
	c.cs = append(c.cs, path)
	return (JsonClearPath)(c)
}

func (c JsonClearKey) GetArgs() []interface{} {
	return c.cs
}

type JsonClearPath Completed

func (c JsonClearPath) GetArgs() []interface{} {
	return c.cs
}

type JsonDebugHelp Completed

func (b Builder) JsonDebugHelp() (c JsonDebugHelp) {
	c.cs = append(c.cs, "JSON.DEBUG", "HELP")
	return c
}

func (c JsonDebugHelp) GetArgs() []interface{} {
	return c.cs
}

type JsonDebugMemory Completed

func (b Builder) JsonDebugMemory() (c JsonDebugMemory) {
	c.cs = append(c.cs, "JSON.DEBUG", "MEMORY")
	return c
}

func (c JsonDebugMemory) Key(key interface{}) JsonDebugMemoryKey {
	c.cs = append(c.cs, key)
	return (JsonDebugMemoryKey)(c)
}

type JsonDebugMemoryKey Completed

func (c JsonDebugMemoryKey) Path(path interface{}) JsonDebugMemoryPath {
	c.cs = append(c.cs, path)
	return (JsonDebugMemoryPath)(c)
}

func (c JsonDebugMemoryKey) GetArgs() []interface{} {
	return c.cs
}

type JsonDebugMemoryPath Completed

func (c JsonDebugMemoryPath) GetArgs() []interface{} {
	return c.cs
}

type JsonDel Completed

func (b Builder) JsonDel() (c JsonDel) {
	c.cs = append(c.cs, "JSON.DEL")
	return c
}

func (c JsonDel) Key(key interface{}) JsonDelKey {
	c.cs = append(c.cs, key)
	return (JsonDelKey)(c)
}

type JsonDelKey Completed

func (c JsonDelKey) Path(path interface{}) JsonDelPath {
	c.cs = append(c.cs, path)
	return (JsonDelPath)(c)
}

func (c JsonDelKey) GetArgs() []interface{} {
	return c.cs
}

type JsonDelPath Completed

func (c JsonDelPath) GetArgs() []interface{} {
	return c.cs
}

type JsonForget Completed

func (b Builder) JsonForget() (c JsonForget) {
	c.cs = append(c.cs, "JSON.FORGET")
	return c
}

func (c JsonForget) Key(key interface{}) JsonForgetKey {
	c.cs = append(c.cs, key)
	return (JsonForgetKey)(c)
}

type JsonForgetKey Completed

func (c JsonForgetKey) Path(path interface{}) JsonForgetPath {
	c.cs = append(c.cs, path)
	return (JsonForgetPath)(c)
}

func (c JsonForgetKey) GetArgs() []interface{} {
	return c.cs
}

type JsonForgetPath Completed

func (c JsonForgetPath) GetArgs() []interface{} {
	return c.cs
}

type JsonGet Completed

func (b Builder) JsonGet() (c JsonGet) {
	c.cs = append(c.cs, "JSON.GET")
	return c
}

func (c JsonGet) Key(key interface{}) JsonGetKey {
	c.cs = append(c.cs, key)
	return (JsonGetKey)(c)
}

type JsonGetIndent Completed

func (c JsonGetIndent) Newline(newline interface{}) JsonGetNewline {
	c.cs = append(c.cs, "NEWLINE", newline)
	return (JsonGetNewline)(c)
}

func (c JsonGetIndent) Space(space interface{}) JsonGetSpace {
	c.cs = append(c.cs, "SPACE", space)
	return (JsonGetSpace)(c)
}

func (c JsonGetIndent) Path(path ...interface{}) JsonGetPath {
	c.cs = append(c.cs, path...)
	return (JsonGetPath)(c)
}

func (c JsonGetIndent) GetArgs() []interface{} {
	return c.cs
}

type JsonGetKey Completed

func (c JsonGetKey) Indent(indent interface{}) JsonGetIndent {
	c.cs = append(c.cs, "INDENT", indent)
	return (JsonGetIndent)(c)
}

func (c JsonGetKey) Newline(newline interface{}) JsonGetNewline {
	c.cs = append(c.cs, "NEWLINE", newline)
	return (JsonGetNewline)(c)
}

func (c JsonGetKey) Space(space interface{}) JsonGetSpace {
	c.cs = append(c.cs, "SPACE", space)
	return (JsonGetSpace)(c)
}

func (c JsonGetKey) Path(path ...interface{}) JsonGetPath {
	c.cs = append(c.cs, path...)
	return (JsonGetPath)(c)
}

func (c JsonGetKey) GetArgs() []interface{} {
	return c.cs
}

type JsonGetNewline Completed

func (c JsonGetNewline) Space(space interface{}) JsonGetSpace {
	c.cs = append(c.cs, "SPACE", space)
	return (JsonGetSpace)(c)
}

func (c JsonGetNewline) Path(path ...interface{}) JsonGetPath {
	c.cs = append(c.cs, path...)
	return (JsonGetPath)(c)
}

func (c JsonGetNewline) GetArgs() []interface{} {
	return c.cs
}

type JsonGetPath Completed

func (c JsonGetPath) Path(path ...interface{}) JsonGetPath {
	c.cs = append(c.cs, path...)
	return c
}

func (c JsonGetPath) GetArgs() []interface{} {
	return c.cs
}

type JsonGetSpace Completed

func (c JsonGetSpace) Path(path ...interface{}) JsonGetPath {
	c.cs = append(c.cs, path...)
	return (JsonGetPath)(c)
}

func (c JsonGetSpace) GetArgs() []interface{} {
	return c.cs
}

type JsonMerge Completed

func (b Builder) JsonMerge() (c JsonMerge) {
	c.cs = append(c.cs, "JSON.MERGE")
	return c
}

func (c JsonMerge) Key(key interface{}) JsonMergeKey {
	c.cs = append(c.cs, key)
	return (JsonMergeKey)(c)
}

type JsonMergeKey Completed

func (c JsonMergeKey) Path(path interface{}) JsonMergePath {
	c.cs = append(c.cs, path)
	return (JsonMergePath)(c)
}

type JsonMergePath Completed

func (c JsonMergePath) Value(value interface{}) JsonMergeValue {
	c.cs = append(c.cs, value)
	return (JsonMergeValue)(c)
}

type JsonMergeValue Completed

func (c JsonMergeValue) GetArgs() []interface{} {
	return c.cs
}

type JsonMget Completed

func (b Builder) JsonMget() (c JsonMget) {
	c.cs = append(c.cs, "JSON.MGET")
	return c
}

func (c JsonMget) Key(key ...interface{}) JsonMgetKey {
	c.cs = append(c.cs, key...)
	return (JsonMgetKey)(c)
}

type JsonMgetKey Completed

func (c JsonMgetKey) Key(key ...interface{}) JsonMgetKey {
	c.cs = append(c.cs, key...)
	return c
}

func (c JsonMgetKey) Path(path interface{}) JsonMgetPath {
	c.cs = append(c.cs, path)
	return (JsonMgetPath)(c)
}

type JsonMgetPath Completed

func (c JsonMgetPath) GetArgs() []interface{} {
	return c.cs
}

type JsonMset Completed

func (b Builder) JsonMset() (c JsonMset) {
	c.cs = append(c.cs, "JSON.MSET")
	return c
}

func (c JsonMset) Key(key interface{}) JsonMsetTripletKey {
	c.cs = append(c.cs, key)
	return (JsonMsetTripletKey)(c)
}

type JsonMsetTripletKey Completed

func (c JsonMsetTripletKey) Path(path interface{}) JsonMsetTripletPath {
	c.cs = append(c.cs, path)
	return (JsonMsetTripletPath)(c)
}

type JsonMsetTripletPath Completed

func (c JsonMsetTripletPath) Value(value interface{}) JsonMsetTripletValue {
	c.cs = append(c.cs, value)
	return (JsonMsetTripletValue)(c)
}

type JsonMsetTripletValue Completed

func (c JsonMsetTripletValue) Key(key interface{}) JsonMsetTripletKey {
	c.cs = append(c.cs, key)
	return (JsonMsetTripletKey)(c)
}

func (c JsonMsetTripletValue) GetArgs() []interface{} {
	return c.cs
}

type JsonNumincrby Completed

func (b Builder) JsonNumincrby() (c JsonNumincrby) {
	c.cs = append(c.cs, "JSON.NUMINCRBY")
	return c
}

func (c JsonNumincrby) Key(key interface{}) JsonNumincrbyKey {
	c.cs = append(c.cs, key)
	return (JsonNumincrbyKey)(c)
}

type JsonNumincrbyKey Completed

func (c JsonNumincrbyKey) Path(path interface{}) JsonNumincrbyPath {
	c.cs = append(c.cs, path)
	return (JsonNumincrbyPath)(c)
}

type JsonNumincrbyPath Completed

func (c JsonNumincrbyPath) Value(value float64) JsonNumincrbyValue {
	c.cs = append(c.cs, strconv.FormatFloat(value, 'f', -1, 64))
	return (JsonNumincrbyValue)(c)
}

type JsonNumincrbyValue Completed

func (c JsonNumincrbyValue) GetArgs() []interface{} {
	return c.cs
}

type JsonNummultby Completed

func (b Builder) JsonNummultby() (c JsonNummultby) {
	c.cs = append(c.cs, "JSON.NUMMULTBY")
	return c
}

func (c JsonNummultby) Key(key interface{}) JsonNummultbyKey {
	c.cs = append(c.cs, key)
	return (JsonNummultbyKey)(c)
}

type JsonNummultbyKey Completed

func (c JsonNummultbyKey) Path(path interface{}) JsonNummultbyPath {
	c.cs = append(c.cs, path)
	return (JsonNummultbyPath)(c)
}

type JsonNummultbyPath Completed

func (c JsonNummultbyPath) Value(value float64) JsonNummultbyValue {
	c.cs = append(c.cs, strconv.FormatFloat(value, 'f', -1, 64))
	return (JsonNummultbyValue)(c)
}

type JsonNummultbyValue Completed

func (c JsonNummultbyValue) GetArgs() []interface{} {
	return c.cs
}

type JsonObjkeys Completed

func (b Builder) JsonObjkeys() (c JsonObjkeys) {
	c.cs = append(c.cs, "JSON.OBJKEYS")
	return c
}

func (c JsonObjkeys) Key(key interface{}) JsonObjkeysKey {
	c.cs = append(c.cs, key)
	return (JsonObjkeysKey)(c)
}

type JsonObjkeysKey Completed

func (c JsonObjkeysKey) Path(path interface{}) JsonObjkeysPath {
	c.cs = append(c.cs, path)
	return (JsonObjkeysPath)(c)
}

func (c JsonObjkeysKey) GetArgs() []interface{} {
	return c.cs
}

type JsonObjkeysPath Completed

func (c JsonObjkeysPath) GetArgs() []interface{} {
	return c.cs
}

type JsonObjlen Completed

func (b Builder) JsonObjlen() (c JsonObjlen) {
	c.cs = append(c.cs, "JSON.OBJLEN")
	return c
}

func (c JsonObjlen) Key(key interface{}) JsonObjlenKey {
	c.cs = append(c.cs, key)
	return (JsonObjlenKey)(c)
}

type JsonObjlenKey Completed

func (c JsonObjlenKey) Path(path interface{}) JsonObjlenPath {
	c.cs = append(c.cs, path)
	return (JsonObjlenPath)(c)
}

func (c JsonObjlenKey) GetArgs() []interface{} {
	return c.cs
}

type JsonObjlenPath Completed

func (c JsonObjlenPath) GetArgs() []interface{} {
	return c.cs
}

type JsonResp Completed

func (b Builder) JsonResp() (c JsonResp) {
	c.cs = append(c.cs, "JSON.RESP")
	return c
}

func (c JsonResp) Key(key interface{}) JsonRespKey {
	c.cs = append(c.cs, key)
	return (JsonRespKey)(c)
}

type JsonRespKey Completed

func (c JsonRespKey) Path(path interface{}) JsonRespPath {
	c.cs = append(c.cs, path)
	return (JsonRespPath)(c)
}

func (c JsonRespKey) GetArgs() []interface{} {
	return c.cs
}

type JsonRespPath Completed

func (c JsonRespPath) GetArgs() []interface{} {
	return c.cs
}

type JsonSet Completed

func (b Builder) JsonSet() (c JsonSet) {
	c.cs = append(c.cs, "JSON.SET")
	return c
}

func (c JsonSet) Key(key interface{}) JsonSetKey {
	c.cs = append(c.cs, key)
	return (JsonSetKey)(c)
}

type JsonSetConditionNx Completed

func (c JsonSetConditionNx) GetArgs() []interface{} {
	return c.cs
}

type JsonSetConditionXx Completed

func (c JsonSetConditionXx) GetArgs() []interface{} {
	return c.cs
}

type JsonSetKey Completed

func (c JsonSetKey) Path(path interface{}) JsonSetPath {
	c.cs = append(c.cs, path)
	return (JsonSetPath)(c)
}

type JsonSetPath Completed

func (c JsonSetPath) Value(value interface{}) JsonSetValue {
	c.cs = append(c.cs, value)
	return (JsonSetValue)(c)
}

type JsonSetValue Completed

func (c JsonSetValue) Nx() JsonSetConditionNx {
	c.cs = append(c.cs, "NX")
	return (JsonSetConditionNx)(c)
}

func (c JsonSetValue) Xx() JsonSetConditionXx {
	c.cs = append(c.cs, "XX")
	return (JsonSetConditionXx)(c)
}

func (c JsonSetValue) GetArgs() []interface{} {
	return c.cs
}

type JsonStrappend Completed

func (b Builder) JsonStrappend() (c JsonStrappend) {
	c.cs = append(c.cs, "JSON.STRAPPEND")
	return c
}

func (c JsonStrappend) Key(key interface{}) JsonStrappendKey {
	c.cs = append(c.cs, key)
	return (JsonStrappendKey)(c)
}

type JsonStrappendKey Completed

func (c JsonStrappendKey) Path(path interface{}) JsonStrappendPath {
	c.cs = append(c.cs, path)
	return (JsonStrappendPath)(c)
}

func (c JsonStrappendKey) Value(value interface{}) JsonStrappendValue {
	c.cs = append(c.cs, value)
	return (JsonStrappendValue)(c)
}

type JsonStrappendPath Completed

func (c JsonStrappendPath) Value(value interface{}) JsonStrappendValue {
	c.cs = append(c.cs, value)
	return (JsonStrappendValue)(c)
}

type JsonStrappendValue Completed

func (c JsonStrappendValue) GetArgs() []interface{} {
	return c.cs
}

type JsonStrlen Completed

func (b Builder) JsonStrlen() (c JsonStrlen) {
	c.cs = append(c.cs, "JSON.STRLEN")
	return c
}

func (c JsonStrlen) Key(key interface{}) JsonStrlenKey {
	c.cs = append(c.cs, key)
	return (JsonStrlenKey)(c)
}

type JsonStrlenKey Completed

func (c JsonStrlenKey) Path(path interface{}) JsonStrlenPath {
	c.cs = append(c.cs, path)
	return (JsonStrlenPath)(c)
}

func (c JsonStrlenKey) GetArgs() []interface{} {
	return c.cs
}

type JsonStrlenPath Completed

func (c JsonStrlenPath) GetArgs() []interface{} {
	return c.cs
}

type JsonToggle Completed

func (b Builder) JsonToggle() (c JsonToggle) {
	c.cs = append(c.cs, "JSON.TOGGLE")
	return c
}

func (c JsonToggle) Key(key interface{}) JsonToggleKey {
	c.cs = append(c.cs, key)
	return (JsonToggleKey)(c)
}

type JsonToggleKey Completed

func (c JsonToggleKey) Path(path interface{}) JsonTogglePath {
	c.cs = append(c.cs, path)
	return (JsonTogglePath)(c)
}

type JsonTogglePath Completed

func (c JsonTogglePath) GetArgs() []interface{} {
	return c.cs
}

type JsonType Completed

func (b Builder) JsonType() (c JsonType) {
	c.cs = append(c.cs, "JSON.TYPE")
	return c
}

func (c JsonType) Key(key interface{}) JsonTypeKey {
	c.cs = append(c.cs, key)
	return (JsonTypeKey)(c)
}

type JsonTypeKey Completed

func (c JsonTypeKey) Path(path interface{}) JsonTypePath {
	c.cs = append(c.cs, path)
	return (JsonTypePath)(c)
}

func (c JsonTypeKey) GetArgs() []interface{} {
	return c.cs
}

type JsonTypePath Completed

func (c JsonTypePath) GetArgs() []interface{} {
	return c.cs
}

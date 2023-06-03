// Code generated DO NOT EDIT

package cmds

import "strconv"

type FtSugadd Completed

func (b Builder) FtSugadd() (c FtSugadd) {
	c.cs = append(c.cs, "FT.SUGADD")
	return c
}

func (c FtSugadd) Key(key interface{}) FtSugaddKey {
	c.cs = append(c.cs, key)
	return (FtSugaddKey)(c)
}

type FtSugaddIncrementScoreIncr Completed

func (c FtSugaddIncrementScoreIncr) Payload(payload interface{}) FtSugaddPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSugaddPayload)(c)
}

func (c FtSugaddIncrementScoreIncr) GetArgs() []interface{} {
	return c.cs
}

type FtSugaddKey Completed

func (c FtSugaddKey) String(string interface{}) FtSugaddString {
	c.cs = append(c.cs, string)
	return (FtSugaddString)(c)
}

type FtSugaddPayload Completed

func (c FtSugaddPayload) GetArgs() []interface{} {
	return c.cs
}

type FtSugaddScore Completed

func (c FtSugaddScore) Incr() FtSugaddIncrementScoreIncr {
	c.cs = append(c.cs, "INCR")
	return (FtSugaddIncrementScoreIncr)(c)
}

func (c FtSugaddScore) Payload(payload interface{}) FtSugaddPayload {
	c.cs = append(c.cs, "PAYLOAD", payload)
	return (FtSugaddPayload)(c)
}

func (c FtSugaddScore) GetArgs() []interface{} {
	return c.cs
}

type FtSugaddString Completed

func (c FtSugaddString) Score(score float64) FtSugaddScore {
	c.cs = append(c.cs, strconv.FormatFloat(score, 'f', -1, 64))
	return (FtSugaddScore)(c)
}

type FtSugdel Completed

func (b Builder) FtSugdel() (c FtSugdel) {
	c.cs = append(c.cs, "FT.SUGDEL")
	return c
}

func (c FtSugdel) Key(key interface{}) FtSugdelKey {
	c.cs = append(c.cs, key)
	return (FtSugdelKey)(c)
}

type FtSugdelKey Completed

func (c FtSugdelKey) String(string interface{}) FtSugdelString {
	c.cs = append(c.cs, string)
	return (FtSugdelString)(c)
}

type FtSugdelString Completed

func (c FtSugdelString) GetArgs() []interface{} {
	return c.cs
}

type FtSugget Completed

func (b Builder) FtSugget() (c FtSugget) {
	c.cs = append(c.cs, "FT.SUGGET")
	return c
}

func (c FtSugget) Key(key interface{}) FtSuggetKey {
	c.cs = append(c.cs, key)
	return (FtSuggetKey)(c)
}

type FtSuggetFuzzy Completed

func (c FtSuggetFuzzy) Withscores() FtSuggetWithscores {
	c.cs = append(c.cs, "WITHSCORES")
	return (FtSuggetWithscores)(c)
}

func (c FtSuggetFuzzy) Withpayloads() FtSuggetWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSuggetWithpayloads)(c)
}

func (c FtSuggetFuzzy) Max(max int64) FtSuggetMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(max, 10))
	return (FtSuggetMax)(c)
}

func (c FtSuggetFuzzy) GetArgs() []interface{} {
	return c.cs
}

type FtSuggetKey Completed

func (c FtSuggetKey) Prefix(prefix interface{}) FtSuggetPrefix {
	c.cs = append(c.cs, prefix)
	return (FtSuggetPrefix)(c)
}

type FtSuggetMax Completed

func (c FtSuggetMax) GetArgs() []interface{} {
	return c.cs
}

type FtSuggetPrefix Completed

func (c FtSuggetPrefix) Fuzzy() FtSuggetFuzzy {
	c.cs = append(c.cs, "FUZZY")
	return (FtSuggetFuzzy)(c)
}

func (c FtSuggetPrefix) Withscores() FtSuggetWithscores {
	c.cs = append(c.cs, "WITHSCORES")
	return (FtSuggetWithscores)(c)
}

func (c FtSuggetPrefix) Withpayloads() FtSuggetWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSuggetWithpayloads)(c)
}

func (c FtSuggetPrefix) Max(max int64) FtSuggetMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(max, 10))
	return (FtSuggetMax)(c)
}

func (c FtSuggetPrefix) GetArgs() []interface{} {
	return c.cs
}

type FtSuggetWithpayloads Completed

func (c FtSuggetWithpayloads) Max(max int64) FtSuggetMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(max, 10))
	return (FtSuggetMax)(c)
}

func (c FtSuggetWithpayloads) GetArgs() []interface{} {
	return c.cs
}

type FtSuggetWithscores Completed

func (c FtSuggetWithscores) Withpayloads() FtSuggetWithpayloads {
	c.cs = append(c.cs, "WITHPAYLOADS")
	return (FtSuggetWithpayloads)(c)
}

func (c FtSuggetWithscores) Max(max int64) FtSuggetMax {
	c.cs = append(c.cs, "MAX", strconv.FormatInt(max, 10))
	return (FtSuggetMax)(c)
}

func (c FtSuggetWithscores) GetArgs() []interface{} {
	return c.cs
}

type FtSuglen Completed

func (b Builder) FtSuglen() (c FtSuglen) {
	c.cs = append(c.cs, "FT.SUGLEN")
	return c
}

func (c FtSuglen) Key(key interface{}) FtSuglenKey {
	c.cs = append(c.cs, key)
	return (FtSuglenKey)(c)
}

type FtSuglenKey Completed

func (c FtSuglenKey) GetArgs() []interface{} {
	return c.cs
}

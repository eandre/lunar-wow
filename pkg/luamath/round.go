package luamath

import "github.com/eandre/lunar/lua"

var floor = lua.Raw(`math.floor`).(func(val float32) int32)
var ceil = lua.Raw(`math.ceil`).(func(val float32) int32)

func Floor(val float32) int32 {
	return floor(val)
}

func Ceil(val float32) int32 {
	return ceil(val)
}

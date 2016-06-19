package luamath

import "github.com/eandre/lunar/lua"

var floor = lua.Raw(`math.floor`).(func(val float32) int)
var ceil = lua.Raw(`math.ceil`).(func(val float32) int)

func Floor(val float32) int {
	return floor(val)
}

func Ceil(val float32) int {
	return ceil(val)
}

package luautil

import (
	"github.com/eandre/lunar/lua"
)

func IsNil(val interface{}) bool {
	return lua.Raw(`val == nil`).(bool)
}

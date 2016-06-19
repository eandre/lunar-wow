package wow

import "github.com/eandre/lunar/lua"

type Time float32

func GetTime() (t Time) {
	return lua.Raw(`GetTime()`).(Time)
}

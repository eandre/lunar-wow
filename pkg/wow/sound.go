package wow

import "github.com/eandre/lunar/lua"

func PlaySound(sound, channel string) {
	lua.Raw(`PlaySound(sound, channel)`)
}

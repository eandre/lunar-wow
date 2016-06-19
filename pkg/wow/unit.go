package wow

type UnitID string
type GUID string

func PlayerFacing() float32 {
	return 0
}

func UnitName(unit UnitID) (name, realm string) {
	return "", ""
}

func UnitGUID(unit UnitID) GUID {
	return ""
}

func UnitIsUnit(a, b UnitID) bool {
	return false
}

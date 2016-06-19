package wow

type AuraFilter string

const (
	AuraCancelable    = "CANCELABLE"
	AuraNotCancelable = "NOT_CANCELABLE"
	AuraPlayer        = "PLAYER"
	AuraRaid          = "RAID"
	AuraHarmful       = "HARMFUL"
	AuraHelpful       = "HELPFUL"
)

func UnitAura(unit UnitID, index int, filter AuraFilter) (name, rank, icon string, count int, dispelType string, duration float32, expires Time, caster UnitID, isStealable, shouldConsolidate bool, spellID int64, canApplyAura, isBossDebuff bool, value1, value2, value3 float32) {
	return "", "", "", 0, "", 0, 0, "", false, false, 0, false, false, 0, 0, 0
}

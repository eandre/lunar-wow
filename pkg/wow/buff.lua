-- Package declaration
local _wow = _G["github.com/eandre/lunar-wow/pkg/wow"] or {}
_G["github.com/eandre/lunar-wow/pkg/wow"] = _wow

_wow.AuraCancelable = "CANCELABLE"
_wow.AuraNotCancelable = "NOT_CANCELABLE"
_wow.AuraPlayer = "PLAYER"
_wow.AuraRaid = "RAID"
_wow.AuraHarmful = "HARMFUL"
_wow.AuraHelpful = "HELPFUL"

local UnitAura = UnitAura

_wow.UnitAura = function(unit, index, filter)
	local name, rank, icon, count, dispelType, duration, expires, caster, isStealable, shouldConsolidate, spellID, canApplyAura, isBossDebuff, value1, value2, value3 = UnitAura(unit, index, filter)
	if name == nil then
		return "", "", "", 0, "", 0, 0, "", false, false, 0, false, false, 0, 0, 0
	end
	return name, rank, icon, count, dispelType, duration, expires, caster or "", isStealable == 1, shouldConsolidate == 1, spellID, canApplyAura == 1, isBossDebuff == 1, value1, value2, value3
end
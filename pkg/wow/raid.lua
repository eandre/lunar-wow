-- Package declaration
local _wow = _G["github.com/eandre/lunar-wow/pkg/wow"] or {}
_G["github.com/eandre/lunar-wow/pkg/wow"] = _wow

_wow.RaidRankMember = 0
_wow.RaidRankAssistant = 1
_wow.RaidRankLeader = 2

_wow.GetNumGroupMembers = function()
    return GetNumGroupMembers()
end

_wow.GetRaidRosterInfo = function(idx)
	local name, rank, subgroup, level, class, fileName, zone, online, isDead, role, isML = GetRaidRosterInfo(idx)
	online = online == 1
	isDead = isDead == 1
	isML = isML == 1
	return name, rank, subgroup, level, class, fileName, zone, online, isDead, role, isML
end

-- Package declaration
local _wow = _G["github.com/eandre/lunar-wow/pkg/wow"] or {}
_G["github.com/eandre/lunar-wow/pkg/wow"] = _wow

_wow.PlayerFacing = function()
    return GetPlayerFacing()
end

_wow.UnitName = function(unit)
    local name, realm = UnitName(unit)
    if realm == nil then
        realm = GetRealmName()
    end
    return name, realm
end

_wow.UnitGUID = function(unit)
    return UnitGUID(unit)
end

_wow.UnitIsUnit = function(a, b)
    return UnitIsUnit(a, b)
end
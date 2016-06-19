-- Package declaration
local _wow = _G["github.com/eandre/lunar-wow/pkg/wow"] or {}
_G["github.com/eandre/lunar-wow/pkg/wow"] = _wow

_wow.ChatTypeRaid = "RAID"
_wow.AddonChatTypeRaid = "RAID"

_wow.SendChatMessage = function(msg, chatType, channel)
    SendChatMessage(msg, chatType, nil, channel)
end

_wow.SendAddonMessage = function(prefix, msg, chatType, target)
    SendAddonMessage(prefix, msg, chatType, target)
end

_wow.RegisterAddonMessagePrefix = function(prefix)
	return RegisterAddonMessagePrefix(prefix)
end
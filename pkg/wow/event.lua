-- Package declaration
local wow = _G["github.com/eandre/lunar-wow/pkg/wow"] or {}
_G["github.com/eandre/lunar-wow/pkg/wow"] = wow

local f = CreateFrame("Frame")

local eventMap = {}
local updates = {}

f:SetScript("OnEvent", function(self, event, ...)
    event = event:upper()
    local m = eventMap[event]
    if m == nil then
        return
    end
    local args = {...}
    for listener in pairs(m) do
        listener(event, args)
    end
end)

f:SetScript("OnUpdate", function(self, dt)
    for listener in pairs(updates) do
        listener(dt)
    end
end)

wow.RegisterEvent = function(event, listener)
    event = event:upper()
    local m = eventMap[event]
    if m == nil then
        m = {}
        eventMap[event] = m
    end
    m[listener] = true
    f:RegisterEvent(event)
end

wow.UnregisterEvent = function(event, listener)
    event = event:upper()
    local m = eventMap[event]
    if m ~= nil then
        m[listener] = nil
        
        -- Don't unregister if table is not empty
        for f in pairs(m) do
            return
        end
        f:UnregisterEvent(event)
    end
end

wow.RegisterUpdate = function(listener)
    updates[listener] = true
end

wow.UnregisterUpdate = function(listener)
    updates[listener] = nil
end
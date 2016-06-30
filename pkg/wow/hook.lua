-- Package declaration
local wow = _G["github.com/eandre/lunar-wow/pkg/wow"] or {}
_G["github.com/eandre/lunar-wow/pkg/wow"] = wow

wow.HookSecureFunc = function(name, f)
	hooksecurefunc(name, f)
end
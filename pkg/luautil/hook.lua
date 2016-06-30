-- Package declaration
local luautil = _G["github.com/eandre/lunar-wow/pkg/luautil"] or {}
_G["github.com/eandre/lunar-wow/pkg/luautil"] = luautil

local builtins = _G.lunar_go_builtins

local loadstring = loadstring

luautil.HookMethod = function(obj, method, f)
	local o = _G[obj]
	local orig = o[method]
	o[method] = function(self, ...)
		return f(orig, self, ...)
	end
end

luautil.HookFunc = function(name, f)
	local orig = _G[name]
	_G[name] = function(...)
		return f(orig, ...)
	end
end
-- Package declaration
local luastrings = _G["github.com/eandre/lunar-wow/pkg/luastrings"] or {}
_G["github.com/eandre/lunar-wow/pkg/luastrings"] = luastrings

local builtins = _G.lunar_go_builtins

local loadstring = loadstring
local strsplit = strsplit
local tonumber = tonumber
local floor = math.floor
local ceil = math.ceil

luastrings.HasPrefix = function(str, prefix)
	local l = #prefix
	return str:sub(1, l) == prefix
end

luastrings.ToString = function(i)
	return "" .. i
end

luastrings.ToInt = function(i)
	local n = tonumber(i)
	if n < 0 then
		return ceil(n)
	else
		return floor(n)
	end
end

luastrings.ToFloat = function(i)
	return tonumber(i)
end

luastrings.Split = function(sep, str, maxPieces)
	if maxPieces >= 0 then
		return {strsplit(sep, str, maxPieces)}
	end
	return {strsplit(sep, str)}
end

luastrings.Find = function(str, substr)
	return not not str:find(substr)
end
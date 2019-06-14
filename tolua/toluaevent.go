package tolua

import (
//"github.com/iakud/luago/lua"
)

/*
func class_index_event(L *lua.Lua_State) int {
	int type = lua.Lua_type(L, 1);
	if (type == LUA_TUSERDATA) {
		lua.Lua_getfield(L, LUA_REGISTRYINDEX, "tolua_peers");
		lua.Lua_pushvalue(L, 1);
		lua.Lua_rawget(L, -2);
		if (lua.Lua_istable(L, -1)) {
			lua.Lua_pushvalue(L, 2);
			lua.Lua_rawget(L, -2);
			if (!lua.Lua_isnil(L, -1)) {
				return 1;
			}
		}

		lua.Lua_settop(L, 2);
		// Try metatables
		lua.Lua_pushvalue(L, 1);
		while (lua.Lua_getmetatable(L, -1)) {
			lua.Lua_remove(L, -2);
			lua.Lua_pushvalue(L, 2);
			lua.Lua_rawget(L, -2);
			if (!lua.Lua_isnil(L, -1)) {
				return 1;
			}
			lua.Lua_pop(L, 1);
		}
		lua.Lua_pushnil(L);
		return 1;
	} else if (type == LUA_TTABLE) {
		lua.Lua_pushvalue(L, 1);
		while (lua.Lua_getmetatable(L, -1)) {
			lua.Lua_remove(L, -2);
			lua.Lua_pushvalue(L, 2);
			lua.Lua_rawget(L, -2);
			if (!lua.Lua_isnil(L, -1)) {
				return 1;
			}
			lua.Lua_pop(L, 1);
		}
		lua.Lua_pushnil(L);
		return 1;
	}
	lua.Lua_pushnil(L);
	return 1;
}

func class_newindex_event(L *lua.Lua_State) int {
	int type = lua.Lua_type(L,1);
	if (type == LUA_TUSERDATA) {
		lua.Lua_getfield(L, LUA_REGISTRYINDEX, "tolua_peers");
		lua.Lua_pushvalue(L, 1);
		lua.Lua_rawget(L, -2);
		if (!lua.Lua_istable(L, -1)) {
			lua.Lua_pop(L, 1);
			lua.Lua_newtable(L);
			lua.Lua_pushvalue(L, 1);
			lua.Lua_pushvalue(L, -2);
			lua.Lua_rawset(L, -4);
		}
		lua.Lua_replace(L, 1);
		lua.Lua_pop(L, 1);
		lua.Lua_rawset(L, 1);
	} else if (type == LUA_TTABLE) {
		lua.Lua_getmetatable(L, 1);
		lua.Lua_replace(L, 1);
		lua.Lua_rawset(L, 1);
	}
	return 0;
}

func tolua_classevents(L *lua.Lua_State) {
	lua.Lua_pushstring(L, "__index");
	lua.Lua_pushcfunction(L, class_index_event);
	lua.Lua_rawset(L, -3);
	lua.Lua_pushstring(L, "__newindex");
	lua.Lua_pushcfunction(L, class_newindex_event);
	lua.Lua_rawset(L, -3);
}

func tolua_collector(L *lua.Lua_State, col lua.Lua_CFunction) {
	lua.Lua_pushstring(L, "__gc");
	lua.Lua_pushcfunction(L, col);
	lua.Lua_rawset(L, -3);
}*/

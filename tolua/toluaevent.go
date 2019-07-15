package tolua

/*
#include <lua.h>

int class_index_event(lua_State *L) {
	int type = lua_type(L, 1);
	if (type == LUA_TUSERDATA) {
		lua_getfield(L, LUA_REGISTRYINDEX, "tolua_peers");
		lua_pushvalue(L, 1);
		lua_rawget(L, -2);
		if (lua_istable(L, -1)) {
			lua_pushvalue(L, 2);
			lua_rawget(L, -2);
			if (!lua_isnil(L, -1)) {
				return 1;
			}
		}

		lua_settop(L, 2);
		// Try metatables
		lua_pushvalue(L, 1);
		while (lua_getmetatable(L, -1)) {
			lua_remove(L, -2);
			lua_pushvalue(L, 2);
			lua_rawget(L, -2);
			if (!lua_isnil(L, -1)) {
				return 1;
			}
			lua_pop(L, 1);
		}
		lua_pushnil(L);
		return 1;
	} else if (type == LUA_TTABLE) {
		lua_pushvalue(L, 1);
		while (lua_getmetatable(L, -1)) {
			lua_remove(L, -2);
			lua_pushvalue(L, 2);
			lua_rawget(L, -2);
			if (!lua_isnil(L, -1)) {
				return 1;
			}
			lua_pop(L, 1);
		}
		lua_pushnil(L);
		return 1;
	}
	lua_pushnil(L);
	return 1;
}

int class_newindex_event(lua_State *L) {
	int type = lua_type(L,1);
	if (type == LUA_TUSERDATA) {
		lua_getfield(L, LUA_REGISTRYINDEX, "tolua_peers");
		lua_pushvalue(L, 1);
		lua_rawget(L, -2);
		if (!lua_istable(L, -1)) {
			lua_pop(L, 1);
			lua_newtable(L);
			lua_pushvalue(L, 1);
			lua_pushvalue(L, -2);
			lua_rawset(L, -4);
		}
		lua_replace(L, 1);
		lua_pop(L, 1);
		lua_rawset(L, 1);
	} else if (type == LUA_TTABLE) {
		lua_getmetatable(L, 1);
		lua_replace(L, 1);
		lua_rawset(L, 1);
	}
	return 0;
}
*/
import "C"

import (
	"github.com/iakud/luago/lua"
)

func classevents(L *lua.Lua_State) {
	lua.Lua_pushstring(L, "__index")
	lua.Lua_pushcfunction(L, (lua.Lua_CFunction)(C.class_index_event))
	lua.Lua_rawset(L, -3)
	lua.Lua_pushstring(L, "__newindex")
	lua.Lua_pushcfunction(L, (lua.Lua_CFunction)(C.class_newindex_event))
	lua.Lua_rawset(L, -3)
}

func collector(L *lua.Lua_State, col lua.Lua_CFunction) {
	lua.Lua_pushstring(L, "__gc")
	lua.Lua_pushcfunction(L, col)
	lua.Lua_rawset(L, -3)
}

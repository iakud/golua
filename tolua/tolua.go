package tolua

import (
	"github.com/iakud/luago/lua"
)

func Open(L *lua.Lua_State) {
	// usertype
	lua.Lua_newtable(L)
	lua.Lua_setfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertable")
	// super
	lua.Lua_newtable(L)
	lua.Lua_setfield(L, lua.LUA_REGISTRYINDEX, "tolua_super")
	// tolua_peers
	lua.Lua_newtable(L)
	lua.Lua_createtable(L, 0, 1)
	lua.Lua_pushliteral(L, "__mode")
	lua.Lua_pushliteral(L, "k")
	lua.Lua_rawset(L, -3)
	lua.Lua_setmetatable(L, -2)
	lua.Lua_setfield(L, lua.LUA_REGISTRYINDEX, "tolua_peers")
	// usertype mapping
	lua.Lua_newtable(L)
	lua.Lua_createtable(L, 0, 1)
	lua.Lua_pushliteral(L, "__mode")
	lua.Lua_pushliteral(L, "v")
	lua.Lua_rawset(L, -3)
	lua.Lua_setmetatable(L, -2)
	lua.Lua_setfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertype_mapping")
	// function mapping
	lua.Lua_newtable(L)
	lua.Lua_setfield(L, lua.LUA_REGISTRYINDEX, "tolua_function_ref")
}

func Module(L *lua.Lua_State, name string) {
	if len(name) > 0 {
		lua.Lua_pushstring(L, name)
		lua.Lua_rawget(L, -2)
		if !lua.Lua_istable(L, -1) {
			lua.Lua_pop(L, 1)
			lua.Lua_newtable(L)
			lua.Lua_pushstring(L, name)
			lua.Lua_pushvalue(L, -2)
			lua.Lua_rawset(L, -4)
		}
		lua.Lua_pop(L, 1)
	}
}

func BeginModule(L *lua.Lua_State, name string) {
	if len(name) > 0 {
		lua.Lua_pushstring(L, name)
		lua.Lua_rawget(L, -2)
		if lua.Lua_istable(L, -1) {
			return
		}
		lua.Lua_pop(L, 1)
	}
	lua.Lua_pushvalue(L, lua.LUA_GLOBALSINDEX)
}

func EndModule(L *lua.Lua_State) {
	lua.Lua_pop(L, 1)
}

func Function(L *lua.Lua_State, name string, f lua.Lua_CFunction) {
	lua.Lua_pushstring(L, name)
	lua.Lua_pushcfunction(L, f)
	lua.Lua_rawset(L, -3)
}

/*
func tolua_usertype(L *lua.Lua_State, name string, col lua.Lua_CFunction) {
	if lua.LuaL_newmetatable(L,name) > 0 {
		tolua_classevents(L);
		if col != nil {
			tolua_collector(L, col);
		}
	}
	lua.Lua_pop(L, 1);
}

func tolua_inheritance(L *lua.Lua_State, name, base string) {
	lua.LuaL_getmetatable(L, name);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		return;
	}
	if len(base)>0 {
		lua.LuaL_getmetatable(L, base);
	} else {
		lua.Lua_pushnil(L);
	}

	if (!lua.Lua_isnil(L, -1)) {
		lua.Lua_pushstring(L, "tolua_usertype_mapping");
		lua.Lua_rawget(L,-2);
	} else {
		lua.Lua_pushnil(L);
	}

	if (!lua.Lua_isnil(L, -1)) {
		lua.Lua_pushstring(L, "tolua_usertype_mapping");
		lua.Lua_insert(L, -2);
		lua.Lua_rawset(L,-4);
	} else {
		lua.Lua_pop(L, 1);
		lua.Lua_pushstring(L, "tolua_usertype_mapping");
		lua.Lua_newtable(L);
		lua.Lua_createtable(L, 0, 1);
		lua.Lua_pushliteral(L, "__mode");
		lua.Lua_pushliteral(L, "v");
		lua.Lua_rawset(L, -3);
		lua.Lua_setmetatable(L, -2);
		lua.Lua_rawset(L, -4);
	}
	lua.Lua_setmetatable(L, -2);
	lua.Lua_pop(L, 1);
}

func tolua_super(L *lua.Lua_State, name, base string) {
	if len(base) > 0 {
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_super");
		lua.LuaL_getmetatable(L, base);
		if (lua.Lua_isnil(L, -1)) {
			lua.Lua_pop(L, 2);
			return;
		}
	} else {
		return;
	}

	lua.Lua_rawget(L, -2);
	lua.LuaL_getmetatable(L, name);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 3);
		return;
	}
	lua.Lua_rawget(L, -3);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		lua.Lua_newtable(L);
		lua.LuaL_getmetatable(L, name);
		lua.Lua_pushvalue(L, -2);
		lua.Lua_rawset(L, -5);
	}
	lua.Lua_pushstring(L, base);
	lua.Lua_pushboolean(L, 1);
	lua.Lua_rawset(L, -3);
	lua.Lua_replace(L, -3);
	if (lua.Lua_istable(L, -1)) {
		lua.Lua_pushnil(L);
		while (lua.Lua_next(L, -2) != 0) {
			lua.Lua_pushvalue(L, -2);
			lua.Lua_insert(L, -2);
			lua.Lua_rawset(L, -5);
		}
	}
	lua.Lua_pop(L, 2);
}

func tolua_usertable(L *lua.Lua_State, lname, name string) {
	lua.Lua_newtable(L);
	lua.LuaL_getmetatable(L, name);
	lua.Lua_setmetatable(L, -2);
	lua.Lua_pushstring(L, lname);
	lua.Lua_pushvalue(L, -2);
	lua.Lua_rawset(L, -4);

	lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertable");
	lua.Lua_insert(L, -2);
	lua.Lua_pushboolean(L, 1);
	lua.Lua_rawset(L, -3);
	lua.Lua_pop(L, 1);
}

void tolua_class(L *lua.Lua_State, const char* lname, const char* name, const char* base) {
	tolua_inheritance(L, name, base);
	tolua_super(L, name, base);
	tolua_usertable(L, lname, name);
}

void tolua_beginusertype(L *lua.Lua_State, const char* name) {
	lua.LuaL_getmetatable(L, name);
}

void tolua_endusertype(L *lua.Lua_State) {
	lua.Lua_pop(L, 1);
}

int tolua_isusertable(L *lua.Lua_State, int index, const char* name) {
	lua.Lua_pushvalue(L, index);
	lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertable");
	lua.Lua_insert(L, -2);
	lua.Lua_rawget(L, -2);
	if (!lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 2);
		if (lua.Lua_getmetatable(L, index)) {
			lua.LuaL_getmetatable(L, name);
			if (lua.Lua_rawequal(L, -1, -2)) {
				lua.Lua_pop(L, 2);
				return 1;
			}
			lua.Lua_pop(L, 2);
		}
	} else {
		lua.Lua_pop(L, 2);
	}
	return 0;
}

void tolua_pushusertype(L *lua.Lua_State, void* p, const char* name) {
	if (p == NULL) {
		lua.Lua_pushnil(L);
		return;
	}

	lua.LuaL_getmetatable(L, name);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		return;
	}

	lua.Lua_pushstring(L, "tolua_usertype_mapping");
	lua.Lua_rawget(L,-2);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertype_mapping");
	}

	lua.Lua_pushlightuserdata(L, p);
	lua.Lua_rawget(L, -2);

	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		lua.Lua_pushlightuserdata(L, p);
		*reinterpret_cast<void**>(lua.Lua_newuserdata(L, sizeof(void*))) = p;
		lua.Lua_pushvalue(L, -1);
		lua.Lua_insert(L, -5);
		lua.Lua_rawset(L, -3);
		lua.Lua_pop(L, 1);

		lua.Lua_setmetatable(L, -2);
	} else {
		lua.Lua_insert(L, -3);
		lua.Lua_pop(L, 1);
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_super");
		lua.Lua_getmetatable(L, -3);
		lua.Lua_rawget(L, -2);
		if (lua.Lua_istable(L, -1)) {
			lua.Lua_pushstring(L, name);
			lua.Lua_rawget(L, -2);
			if (lua.Lua_toboolean(L, -1) == 1) {
				lua.Lua_pop(L, 4);
				return;
			}
			lua.Lua_pop(L, 1);
		}
		lua.Lua_pop(L, 2);

		lua.Lua_setmetatable(L, -2);
	}
}

void* tolua_tousertype(L *lua.Lua_State, int index, const char* name) {
	if (!lua.Lua_isuserdata(L, index)) {
		return NULL;
	}

	if (lua.Lua_getmetatable(L, index)) {
		lua.LuaL_getmetatable(L, name);
		if (lua.Lua_isnil(L, -1)) {
			lua.Lua_pop(L, 2);
			return NULL;
		}
		if (lua.Lua_rawequal(L, -1, -2)) {
			lua.Lua_pop(L, 2);
			return *reinterpret_cast<void**>(lua.Lua_touserdata(L, index));
		}
		lua.Lua_pop(L, 1);
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_super");
		lua.Lua_insert(L, -2);
		lua.Lua_rawget(L, -2);
		if (lua.Lua_istable(L,-1)) {
			lua.Lua_pushstring(L, name);
			lua.Lua_rawget(L, -2);
			if (lua.Lua_toboolean(L, -1)) {
				lua.Lua_pop(L, 3);
				return *reinterpret_cast<void**>(lua.Lua_touserdata(L, index));
			}
			lua.Lua_pop(L, 1);
		}
		lua.Lua_pop(L, 2);
	}
	return NULL;
}

void tolua_removeusertype(L *lua.Lua_State, void* p, const char* name) {
	if (p == NULL) {
		return;
	}

	lua.LuaL_getmetatable(L, name);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		return;
	}
	lua.Lua_pushstring(L, "tolua_usertype_mapping");
	lua.Lua_rawget(L, -2);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 1);
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertype_mapping");
	}

	lua.Lua_pushlightuserdata(L, p);
	lua.Lua_rawget(L, -2);
	if (lua.Lua_isnil(L, -1)) {
		lua.Lua_pop(L, 3);
		return;
	}

	lua.Lua_pop(L, 1);
	lua.Lua_pushlightuserdata(L, p);
	lua.Lua_pushnil(L);
	lua.Lua_rawset(L, -3);
	lua.Lua_pop(L, 2);
}*/

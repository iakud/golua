package tolua

import (
	"unsafe"

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
		if lua.Lua_istable(L, -1) {
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

func UserType(L *lua.Lua_State, name string, col lua.Lua_CFunction) {
	if lua.LuaL_newmetatable(L, name) > 0 {
		classevents(L)
		if col != nil {
			collector(L, col)
		}
	}
	lua.Lua_pop(L, 1)
}

func inheritance(L *lua.Lua_State, name, base string) {
	lua.LuaL_getmetatable(L, name)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		return
	}
	if len(base) > 0 {
		lua.LuaL_getmetatable(L, base)
	} else {
		lua.Lua_pushnil(L)
	}

	if !lua.Lua_isnil(L, -1) {
		lua.Lua_pushstring(L, "tolua_usertype_mapping")
		lua.Lua_rawget(L, -2)
	} else {
		lua.Lua_pushnil(L)
	}

	if !lua.Lua_isnil(L, -1) {
		lua.Lua_pushstring(L, "tolua_usertype_mapping")
		lua.Lua_insert(L, -2)
		lua.Lua_rawset(L, -4)
	} else {
		lua.Lua_pop(L, 1)
		lua.Lua_pushstring(L, "tolua_usertype_mapping")
		lua.Lua_newtable(L)
		lua.Lua_createtable(L, 0, 1)
		lua.Lua_pushliteral(L, "__mode")
		lua.Lua_pushliteral(L, "v")
		lua.Lua_rawset(L, -3)
		lua.Lua_setmetatable(L, -2)
		lua.Lua_rawset(L, -4)
	}
	lua.Lua_setmetatable(L, -2)
	lua.Lua_pop(L, 1)
}

func super(L *lua.Lua_State, name, base string) {
	if len(base) > 0 {
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_super")
		lua.LuaL_getmetatable(L, base)
		if lua.Lua_isnil(L, -1) {
			lua.Lua_pop(L, 2)
			return
		}
	} else {
		return
	}

	lua.Lua_rawget(L, -2)
	lua.LuaL_getmetatable(L, name)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 3)
		return
	}
	lua.Lua_rawget(L, -3)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		lua.Lua_newtable(L)
		lua.LuaL_getmetatable(L, name)
		lua.Lua_pushvalue(L, -2)
		lua.Lua_rawset(L, -5)
	}
	lua.Lua_pushstring(L, base)
	lua.Lua_pushboolean(L, true)
	lua.Lua_rawset(L, -3)
	lua.Lua_replace(L, -3)
	if lua.Lua_istable(L, -1) {
		lua.Lua_pushnil(L)
		for lua.Lua_next(L, -2) != 0 {
			lua.Lua_pushvalue(L, -2)
			lua.Lua_insert(L, -2)
			lua.Lua_rawset(L, -5)
		}
	}
	lua.Lua_pop(L, 2)
}

func usertable(L *lua.Lua_State, lname, name string) {
	lua.Lua_newtable(L)
	lua.LuaL_getmetatable(L, name)
	lua.Lua_setmetatable(L, -2)
	lua.Lua_pushstring(L, lname)
	lua.Lua_pushvalue(L, -2)
	lua.Lua_rawset(L, -4)

	lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertable")
	lua.Lua_insert(L, -2)
	lua.Lua_pushboolean(L, true)
	lua.Lua_rawset(L, -3)
	lua.Lua_pop(L, 1)
}

func Class(L *lua.Lua_State, lname, name, base string) {
	inheritance(L, name, base)
	super(L, name, base)
	usertable(L, lname, name)
}

func BeginUserType(L *lua.Lua_State, name string) {
	lua.LuaL_getmetatable(L, name)
}

func EndUserType(L *lua.Lua_State) {
	lua.Lua_pop(L, 1)
}

func IsUserTable(L *lua.Lua_State, index int, name string) bool {
	lua.Lua_pushvalue(L, index)
	lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertable")
	lua.Lua_insert(L, -2)
	lua.Lua_rawget(L, -2)
	if !lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 2)
		if lua.Lua_getmetatable(L, index) != 0 {
			lua.LuaL_getmetatable(L, name)
			if lua.Lua_rawequal(L, -1, -2) != 0 {
				lua.Lua_pop(L, 2)
				return true
			}
			lua.Lua_pop(L, 2)
		}
	} else {
		lua.Lua_pop(L, 2)
	}
	return false
}

func PushUserType(L *lua.Lua_State, p unsafe.Pointer, name string) {
	if p == nil {
		lua.Lua_pushnil(L)
		return
	}

	lua.LuaL_getmetatable(L, name)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		return
	}

	lua.Lua_pushstring(L, "tolua_usertype_mapping")
	lua.Lua_rawget(L, -2)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertype_mapping")
	}

	lua.Lua_pushlightuserdata(L, p)
	lua.Lua_rawget(L, -2)

	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		lua.Lua_pushlightuserdata(L, p)
		*((*unsafe.Pointer)(lua.Lua_newuserdata(L, uint(unsafe.Sizeof(p))))) = p
		lua.Lua_pushvalue(L, -1)
		lua.Lua_insert(L, -5)
		lua.Lua_rawset(L, -3)
		lua.Lua_pop(L, 1)

		lua.Lua_setmetatable(L, -2)
	} else {
		lua.Lua_insert(L, -3)
		lua.Lua_pop(L, 1)
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_super")
		lua.Lua_getmetatable(L, -3)
		lua.Lua_rawget(L, -2)
		if lua.Lua_istable(L, -1) {
			lua.Lua_pushstring(L, name)
			lua.Lua_rawget(L, -2)
			if lua.Lua_toboolean(L, -1) {
				lua.Lua_pop(L, 4)
				return
			}
			lua.Lua_pop(L, 1)
		}
		lua.Lua_pop(L, 2)

		lua.Lua_setmetatable(L, -2)
	}
}

func ToUserType(L *lua.Lua_State, index int, name string) unsafe.Pointer {
	if !lua.Lua_isuserdata(L, index) {
		return nil
	}

	if lua.Lua_getmetatable(L, index) != 0 {
		lua.LuaL_getmetatable(L, name)
		if lua.Lua_isnil(L, -1) {
			lua.Lua_pop(L, 2)
			return nil
		}
		if lua.Lua_rawequal(L, -1, -2) != 0 {
			lua.Lua_pop(L, 2)
			return *((*unsafe.Pointer)(lua.Lua_touserdata(L, index)))
		}
		lua.Lua_pop(L, 1)
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_super")
		lua.Lua_insert(L, -2)
		lua.Lua_rawget(L, -2)
		if lua.Lua_istable(L, -1) {
			lua.Lua_pushstring(L, name)
			lua.Lua_rawget(L, -2)
			if lua.Lua_toboolean(L, -1) {
				lua.Lua_pop(L, 3)
				return *((*unsafe.Pointer)(lua.Lua_touserdata(L, index)))
			}
			lua.Lua_pop(L, 1)
		}
		lua.Lua_pop(L, 2)
	}
	return nil
}

func RemoveUserType(L *lua.Lua_State, p unsafe.Pointer, name string) {
	if p == nil {
		return
	}

	lua.LuaL_getmetatable(L, name)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		return
	}
	lua.Lua_pushstring(L, "tolua_usertype_mapping")
	lua.Lua_rawget(L, -2)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 1)
		lua.Lua_getfield(L, lua.LUA_REGISTRYINDEX, "tolua_usertype_mapping")
	}

	lua.Lua_pushlightuserdata(L, p)
	lua.Lua_rawget(L, -2)
	if lua.Lua_isnil(L, -1) {
		lua.Lua_pop(L, 3)
		return
	}

	lua.Lua_pop(L, 1)
	lua.Lua_pushlightuserdata(L, p)
	lua.Lua_pushnil(L)
	lua.Lua_rawset(L, -3)
	lua.Lua_pop(L, 2)
}

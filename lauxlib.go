package lua

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lluajit -lmingwex

#include <lauxlib.h>
*/
import "C"

func LuaL_newstate() *Lua_State {
	return (*Lua_State)(C.luaL_newstate())
}

func LuaL_loadfile(L *Lua_State, filename string) int {
	return int(C.luaL_loadfile(L, C.CString(filename)))
}

func LuaL_dofile(L *Lua_State, fn string) int {
	if LuaL_loadfile(L, fn) != 0 {
		return 1
	}
	if Lua_pcall(L, 0, C.LUA_MULTRET, 0) != 0 {
		return 1
	}
	return 0
}

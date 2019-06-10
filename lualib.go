package lua

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lluajit -lmingwex

#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>
*/
import "C"

func LuaL_openlibs(L *Lua_State) {
	C.luaL_openlibs(L)
}

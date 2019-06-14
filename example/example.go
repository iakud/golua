package main

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -lluajit -lmingwex

#include "lua.h"
#include <stdlib.h>
int goAdd_cgo(lua_State* L);
*/
import "C"

import (
	"errors"
	"fmt"

	"github.com/iakud/luago/lua"
)

//export goAdd
func goAdd(L *C.lua_State) int {
	fmt.Printf("goAdd_go(): called\n")
	a := lua.Lua_tointeger((*lua.Lua_State)(L), 1)
	b := lua.Lua_tointeger((*lua.Lua_State)(L), 2)
	lua.Lua_pushinteger((*lua.Lua_State)(L), a+b)
	return 1
}

func main() {
	L := lua.LuaL_newstate()
	defer lua.Lua_close(L)
	lua.LuaL_openlibs(L)
	lua.Lua_register(L, "goAdd", (lua.Lua_CFunction)(C.goAdd_cgo))
	lua.LuaL_dofile(L, "test.lua")
	lua.Lua_getglobal(L, "add")

	a, b := 11, 7
	lua.Lua_pushinteger(L, lua.Lua_Integer(a))
	lua.Lua_pushinteger(L, lua.Lua_Integer(b))
	if ret := lua.Lua_pcall(L, 2, 1, 0); ret != 0 {
		luaerr := lua.Lua_tostring(L, -1)
		lua.Lua_pop(L, 1)
		panic(errors.New(luaerr))
	}
	result := lua.Lua_tonumber(L, -1)
	lua.Lua_settop(L, 0)
	fmt.Printf("%v+%v=%v\n", a, b, result)
}

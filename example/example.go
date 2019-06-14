package main

/*
#include <lua.h>
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
	a := lua.Lua_tointeger((*lua.Lua_State)(L), 1)
	b := lua.Lua_tointeger((*lua.Lua_State)(L), 2)
	fmt.Printf("goAdd_go(%v, %v): called\n", int(a), int(b))
	lua.Lua_pushinteger((*lua.Lua_State)(L), a+b)
	return 1
}

func Add(L *lua.Lua_State, a, b int) int {
	lua.Lua_getglobal(L, "add")
	lua.Lua_pushinteger(L, lua.Lua_Integer(a))
	lua.Lua_pushinteger(L, lua.Lua_Integer(b))
	if ret := lua.Lua_pcall(L, 2, 1, 0); ret != 0 {
		luaerr := lua.Lua_tostring(L, -1)
		lua.Lua_pop(L, 1)
		panic(errors.New(luaerr))
	}
	result := int(lua.Lua_tonumber(L, -1))
	lua.Lua_settop(L, 0)
	return result
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	L := lua.LuaL_newstate()
	defer lua.Lua_close(L)
	lua.LuaL_openlibs(L)
	lua.Lua_register(L, "goAdd", (lua.Lua_CFunction)(C.goAdd_cgo))
	lua.LuaL_dofile(L, "test.lua")
	result := Add(L, 11, 7) // result = 18
	fmt.Printf("result=%v\n", result)
}

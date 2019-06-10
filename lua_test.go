package lua

import (
	"fmt"
	"syscall"
	"testing"
)

func GoTest(L *Lua_State) int {
	a := int(Lua_tointeger(L, 1))
	b := int(Lua_tointeger(L, 2))
	Lua_pushinteger(L, a+b)
	return 1
}

var GoTest_CFunc = syscall.NewCallback(GoTest)

func TestLua(t *testing.T) {
	L := LuaL_newstate()
	defer Lua_close(L)
	LuaL_openlibs(L)
	Lua_register(L, "goTest", Lua_CFunction(GoTest_CFunc))
	LuaL_dofile(L, "test.lua")
	Lua_getglobal(L, "test")
	Lua_pushinteger(L, 1)
	Lua_pushinteger(L, 2)
	if ret := Lua_pcall(L, 2, 1, 0); ret != 0 {
		fmt.Println(Lua_tostring(L, -1))
		return
	}
	result := Lua_tonumber(L, -1)
	fmt.Println(result)
}

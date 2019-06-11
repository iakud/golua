package lua

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lluajit -lmingwex

#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type LuaStack C.lua_State

func NewLuaStack() *LuaStack {
	return (*LuaStack)(C.luaL_newstate())
}

func (L *LuaStack) OpenLibs() {
	C.luaL_openlibs(L)
}

func (L *LuaStack) AddPackagePath(path string) {
	L.GetGlobal(C.LUA_LOADLIBNAME)
	L.GetField(-1, "path")
	L.PushString(fmt.Sprintf("%s;%s/?.lua", L.ToString(-1), path))
	L.SetField(-3, "path")
	L.Pop(2)
}

func (L *LuaStack) Load(modname string) {
	if len(modname) == 0 {
		return
	}
	require := fmt.Sprintf("require '%v'", modname)
	L.ExecuteString(require)
}

func (L *LuaStack) Unload(modname string) {
	if len(modname) == 0 {
		return
	}
	L.GetGlobal(C.LUA_LOADLIBNAME)
	L.GetField(-1, "loaded")
	L.PushString(modname)
	C.lua_gettable(L, -2)
	if C.lua_type(L, -1) != C.LUA_TNIL {
		L.PushString(modname)
		L.PushNil()
		C.lua_settable(L, -4)
	}
	L.Pop(3)
}

func (L *LuaStack) Reload(modname string) {
	L.Unload(modname)
	L.Load(modname)
}

//
// push value
//
func (L *LuaStack) PushNil() {
	C.lua_pushnil(L)
}

func (L *LuaStack) PushBool(value bool) {
	if value {
		C.lua_pushboolean(L, 1)
	} else {
		C.lua_pushboolean(L, 0)
	}
}

func (L *LuaStack) PushInt32(value int32) {
	C.lua_pushinteger(L, C.lua_Integer(value))
}

func (L *LuaStack) PushInt64(value int64) {
	C.lua_pushnumber(L, C.lua_Number(value))
}

func (L *LuaStack) PushFloat32(value float32) {
	C.lua_pushnumber(L, C.lua_Number(value))
}

func (L *LuaStack) PushFloat64(value float64) {
	C.lua_pushnumber(L, C.lua_Number(value))
}

func (L *LuaStack) PushString(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))
	C.lua_pushstring(L, c_value)
}

//
// to value
//

func (L *LuaStack) ToBool(index int) bool {
	if C.lua_toboolean(L, C.int(index)) != 0 {
		return true
	} else {
		return false
	}
}

func (L *LuaStack) ToInt32(index int) int32 {
	return int32(C.lua_tointeger(L, C.int(index)))
}

func (L *LuaStack) ToInt64(index int) int64 {
	return int64(C.lua_tonumber(L, C.int(index)))
}

func (L *LuaStack) ToFloat32(index int) float32 {
	return float32(C.lua_tonumber(L, C.int(index)))
}

func (L *LuaStack) ToFloat64(index int) float64 {
	return float64(C.lua_tonumber(L, C.int(index)))
}

func (L *LuaStack) ToString(index int) string {
	return C.GoString(C.lua_tolstring(L, C.int(index), nil))
}

//
func (L *LuaStack) GetTop() int {
	return int(C.lua_gettop(L))
}

func (L *LuaStack) Clean() {
	C.lua_settop(L, 0)
}

func (L *LuaStack) FormatIndex(index int) int {
	if index < 0 {
		return int(C.lua_gettop(L)) + 1 + index
	} else {
		return index
	}
}

//
// excute
//
const MULTRET int = C.LUA_MULTRET

func (L *LuaStack) ExecuteGlobalFunction(funcname string, nargs, nresults int) {
	L.GetGlobal(funcname)
	if nargs > 0 {
		C.lua_insert(L, C.int(-(nargs + 1)))
	}
	L.execute(nargs, nresults)
}

func (L *LuaStack) ExecuteString(codes string) {
	c_codes := C.CString(codes)
	defer C.free(unsafe.Pointer(c_codes))
	C.luaL_loadstring(L, c_codes)
	L.execute(0, 0)
}

func (L *LuaStack) execute(nargs, nresults int) { // LUA_MULTRET
	functionIndex := L.FormatIndex(-(nargs + 1))
	traceback := 0
	L.GetGlobal("__TRACKBACK__")
	if C.lua_type(L, -1) == C.LUA_TFUNCTION {
		C.lua_insert(L, C.int(functionIndex))
		traceback = functionIndex
	} else {
		L.Pop(1)
	}
	if C.lua_pcall(L, C.int(nargs), C.int(nresults), C.int(traceback)) != 0 {
		err := L.ToString(-1)
		if traceback != 0 {
			L.Pop(2)
		} else {
			L.Pop(1)
		}
		panic(err)
	} else if traceback != 0 {
		C.lua_remove(L, C.int(traceback))
	}
}

//
//
//
func (L *LuaStack) GetField(index int, key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))
	C.lua_getfield(L, C.int(index), c_key)
}

func (L *LuaStack) SetField(index int, key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))
	C.lua_setfield(L, C.int(index), c_key)
}

func (L *LuaStack) GetGlobal(key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))
	C.lua_getfield(L, C.LUA_GLOBALSINDEX, c_key)
}

func (L *LuaStack) SetGlobal(key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))
	C.lua_setfield(L, C.LUA_GLOBALSINDEX, c_key)
}

func (L *LuaStack) Pop(n int) {
	C.lua_settop(L, C.int(-n-1))
}

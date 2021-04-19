package tolua

/*
#cgo pkg-config: tolua luajit
#include <tolua/tolua.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/iakud/golua/lua"

	"unsafe"
)

func Open(L *lua.Lua_State) {
	C.tolua_open((*C.lua_State)(L))
}

func Module(L *lua.Lua_State, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_module((*C.lua_State)(L), c_name)
}

func BeginModule(L *lua.Lua_State, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_beginmodule((*C.lua_State)(L), c_name)
}

func EndModule(L *lua.Lua_State) {
	C.tolua_endmodule((*C.lua_State)(L))
}

func Function(L *lua.Lua_State, name string, f lua.Lua_CFunction) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_function((*C.lua_State)(L), c_name, f)
}

func UserType(L *lua.Lua_State, name string, col lua.Lua_CFunction) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_usertype((*C.lua_State)(L), c_name, col)
}

func Class(L *lua.Lua_State, lname, name, base string) {
	c_lname := C.CString(lname)
	defer C.free(unsafe.Pointer(c_lname))
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_base := C.CString(base)
	defer C.free(unsafe.Pointer(c_base))
	C.tolua_class((*C.lua_State)(L), c_lname, c_name, c_base)
}

func BeginUserType(L *lua.Lua_State, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_beginusertype((*C.lua_State)(L), c_name)
}

func EndUserType(L *lua.Lua_State) {
	C.tolua_endusertype((*C.lua_State)(L))
}

func IsUserTable(L *lua.Lua_State, index int, name string) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.tolua_isusertable((*C.lua_State)(L), C.int(index), c_name))
}

func PushUserType(L *lua.Lua_State, p unsafe.Pointer, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_pushusertype((*C.lua_State)(L), p, c_name)
}

func ToUserType(L *lua.Lua_State, index int, name string) unsafe.Pointer {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return C.tolua_tousertype((*C.lua_State)(L), C.int(index), c_name)
}

func RemoveUserType(L *lua.Lua_State, p unsafe.Pointer, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.tolua_removeusertype((*C.lua_State)(L), p, c_name)
}

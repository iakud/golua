package tolua

/*
#include <tolua/tolua_function.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/iakud/luago/lua"
)

type Tolua_FunctionRef C.tolua_FunctionRef

func RefFunction(L *lua.Lua_State, index int) *Tolua_FunctionRef {
	return (*Tolua_FunctionRef)(C.tolua_function_ref((*C.lua_State)(L), C.int(index)))
}

func PushFunctionByRef(L *lua.Lua_State, f *Tolua_FunctionRef) {
	C.tolua_push_function_by_ref((*C.lua_State)(L), (*C.tolua_FunctionRef)(f))
}

func RemoveFunctionByRef(L *lua.Lua_State, f *Tolua_FunctionRef) {
	C.tolua_remove_function_by_ref((*C.lua_State)(L), (*C.tolua_FunctionRef)(f))
}

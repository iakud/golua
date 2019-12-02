package tolua

/*
#include <tolua/tolua_function.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/iakud/luago/lua"
)

type Tolua_Function_Ref C.tolua_function_ref

func RefFunction(L *lua.Lua_State, index int) *Tolua_Function_Ref {
	return (*Tolua_Function_Ref)(C.tolua_ref_function((*C.lua_State)(L), C.int(index)))
}

func PushFunctionByRef(L *lua.Lua_State, f *Tolua_Function_Ref) {
	C.tolua_push_function_by_ref((*C.lua_State)(L), (*C.tolua_function_ref)(f))
}

func RemoveFunctionByRef(L *lua.Lua_State, f *Tolua_Function_Ref) {
	C.tolua_remove_function_by_ref((*C.lua_State)(L), (*C.tolua_function_ref)(f))
}

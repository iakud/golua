package tolua

/*
#include <tolua/tolua_function.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/iakud/golua/lua"
)

type Tolua_FunctionRef C.tolua_FunctionRef

func ToFunctionRef(L *lua.Lua_State, index int) *Tolua_FunctionRef {
	return (*Tolua_FunctionRef)(C.tolua_tofunction_ref((*C.lua_State)(L), C.int(index)))
}

func PushFunctionRef(L *lua.Lua_State, f *Tolua_FunctionRef) {
	C.tolua_pushfunction_ref((*C.lua_State)(L), (*C.tolua_FunctionRef)(f))
}

func RemoveFunctionRef(L *lua.Lua_State, f *Tolua_FunctionRef) {
	C.tolua_removefunction_ref((*C.lua_State)(L), (*C.tolua_FunctionRef)(f))
}

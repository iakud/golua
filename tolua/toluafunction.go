package tolua

/*
#include <tolua/tolua_function.h>
*/
import "C"

import (
	"github.com/iakud/golua/lua"
)

type FunctionRef = C.tolua_FunctionRef

func ToFunctionRef(L *lua.Lua_State, index int) *FunctionRef {
	return C.tolua_tofunction_ref((*C.lua_State)(L), C.int(index))
}

func PushFunctionRef(L *lua.Lua_State, f *FunctionRef) {
	C.tolua_pushfunction_ref((*C.lua_State)(L), f)
}

func RemoveFunctionRef(L *lua.Lua_State, f *FunctionRef) {
	C.tolua_removefunction_ref((*C.lua_State)(L), f)
}

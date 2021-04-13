package tolua

/*
#include <tolua/tolua_call.h>
*/
import "C"

import (
	"github.com/iakud/golua/lua"
)

func DoCall(L *lua.Lua_State, nargs, nresults int) int {
	return int(C.tolua_docall((*C.lua_State)(L), C.int(nargs), C.int(nresults)))
}

package main

/*
typedef struct lua_State lua_State;
int lua_ClassA_getMessage_cgo(lua_State* L);
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/iakud/luago"
	"github.com/iakud/luago/lua"
	"github.com/iakud/luago/tolua"
)

type ClassA struct {
	message string
}

func (this *ClassA) getMessage() string {
	return this.message
}

//export lua_ClassA_getMessage
func lua_ClassA_getMessage(l *C.lua_State) int {
	L := (*lua.Lua_State)(l)
	a := (*ClassA)(tolua.ToUserType(L, 1, "ClassA"))
	if a == nil {
		lua.LuaL_error(L, "invalid 'obj' in function '%s'", "lua_ClassA_getMessage")
		return 0
	}
	argc := lua.Lua_gettop(L) - 1
	if argc == 0 {
		message := a.getMessage()
		lua.Lua_pushstring(L, message)
		return 1
	}
	lua.LuaL_error(L, "'%s' has wrong number of arguments: %d, was expecting %d \n", "lua_ClassA_getMessage", argc, 1)
	return 0
}

func lua_register_class(L *lua.Lua_State) {
	tolua.BeginModule(L, "")
	tolua.UserType(L, "ClassA", nil)
	tolua.Class(L, "ClassA", "ClassA", "")
	tolua.BeginUserType(L, "ClassA")
	{
		tolua.Function(L, "getMessage", (lua.Lua_CFunction)(C.lua_ClassA_getMessage_cgo))
	}
	tolua.EndUserType(L)
	tolua.EndModule(L)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	stack := luago.NewLuaStack()
	defer stack.Close()
	lua_register_class(stack.LuaState())

	stack.AddPackagePath("script")

	a := &ClassA{"hello world!"}
	stack.Load("example")

	stack.PushUserType(unsafe.Pointer(a), "ClassA")
	stack.ExecuteGlobalFunction("setfunc", 1, 0)
	fmt.Println("call setfunc")
	stack.Clean()

	stack.PushUserType(unsafe.Pointer(a), "ClassA")
	stack.ExecuteGlobalFunction("getfunc", 1, 2)
	luanumber := stack.ToInt(-2)
	luastring := stack.ToString(-1)
	fmt.Println(luanumber, luastring)
	stack.Clean()

	stack.PushUserType(unsafe.Pointer(a), "ClassA")
	stack.ExecuteGlobalFunction("showmessage", 1, 0)
	stack.Clean()
}

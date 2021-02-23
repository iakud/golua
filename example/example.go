package main

/*
typedef struct lua_State lua_State;
int lua_ball_getName(lua_State* L);
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/iakud/golua"
	"github.com/iakud/golua/lua"
	"github.com/iakud/golua/tolua"
)

type ball struct {
	name string
}

func (b *ball) getName() string {
	return b.name
}

//export lua_ball_getName
func lua_ball_getName(l *C.lua_State) C.int {
	L := (*lua.Lua_State)(l)
	a := (*ball)(tolua.ToUserType(L, 1, "Ball"))
	if a == nil {
		lua.LuaL_error(L, "invalid 'obj' in function '%s'", "lua_ball_getName")
		return 0
	}
	argc := lua.Lua_gettop(L) - 1
	if argc == 0 {
		name := a.getName()
		lua.Lua_pushstring(L, name)
		return 1
	}
	lua.LuaL_error(L, "'%s' has wrong number of arguments: %d, was expecting %d \n", "lua_ball_getName", argc, 1)
	return 0
}

func lua_register_class(L *lua.Lua_State) {
	tolua.BeginModule(L, "")
	tolua.UserType(L, "Ball", nil)
	tolua.Class(L, "Ball", "Ball", "")
	tolua.BeginUserType(L, "Ball")
	{
		tolua.Function(L, "getName", (lua.Lua_CFunction)(C.lua_ball_getName))
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

	stack := golua.NewLuaStack()
	defer stack.Close()
	lua_register_class(stack.LuaState())

	stack.AddPackagePath("script")

	a := &ball{"hello world!"}
	stack.Load("example")

	stack.PushUserType(unsafe.Pointer(a), "Ball")
	stack.ExecuteGlobalFunction("setfunc", 1, 0)
	fmt.Println("call setfunc")
	stack.Clean()

	stack.PushUserType(unsafe.Pointer(a), "Ball")
	stack.ExecuteGlobalFunction("getfunc", 1, 2)
	luanumber := stack.ToInt(-2)
	luastring := stack.ToString(-1)
	fmt.Println(luanumber, luastring)
	stack.Clean()

	stack.PushUserType(unsafe.Pointer(a), "Ball")
	stack.ExecuteGlobalFunction("showname", 1, 0)
	stack.Clean()
}

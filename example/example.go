package main

import (
	"log"
	"unsafe"

	"github.com/iakud/golua"
	"github.com/iakud/golua/lua"
	"github.com/iakud/golua/tolua"
)

/*
typedef struct lua_State lua_State;
int book_Name(lua_State* L);
int book_SetSellCallback(lua_State* L);
*/
import "C"

type book struct {
	name         string
	sellCallback *tolua.FunctionRef
}

func (b *book) Name() string {
	return b.name
}

func (b *book) SetSellCallback(callback *tolua.FunctionRef) {
	b.sellCallback = callback
}

func (b *book) GetSellCallback() *tolua.FunctionRef {
	return b.sellCallback
}

//export book_Name
func book_Name(l *C.lua_State) C.int {
	L := (*lua.Lua_State)(l)
	b := (*book)(tolua.ToUserType(L, 1, "Book"))
	if b == nil {
		lua.LuaL_error(L, "invalid 'obj' in function '%s'", "book_Name")
		return 0
	}
	argc := lua.Lua_gettop(L) - 1
	if argc == 0 {
		lua.Lua_pushstring(L, b.Name())
		return 1
	}
	lua.LuaL_error(L, "'%s' has wrong number of arguments: %d, was expecting %d \n", "book_Name", argc, 0)
	return 0
}

//export book_SetSellCallback
func book_SetSellCallback(l *C.lua_State) C.int {
	L := (*lua.Lua_State)(l)
	b := (*book)(tolua.ToUserType(L, 1, "Book"))
	if b == nil {
		lua.LuaL_error(L, "invalid 'obj' in function '%s'", "book_Name")
		return 0
	}
	argc := lua.Lua_gettop(L) - 1
	if argc == 1 {
		b.SetSellCallback(tolua.ToFunctionRef(L, 2))
		return 0
	}
	lua.LuaL_error(L, "'%s' has wrong number of arguments: %d, was expecting %d \n", "book_Name", argc, 1)
	return 0
}

func register_book(L *lua.Lua_State) {
	tolua.BeginModule(L, "")
	tolua.UserType(L, "Book", nil)
	tolua.Class(L, "Book", "Book", "")
	tolua.BeginUserType(L, "Book")
	{
		tolua.Function(L, "Name", (lua.Lua_CFunction)(C.book_Name))
		tolua.Function(L, "SetSellCallback", (lua.Lua_CFunction)(C.book_SetSellCallback))
	}
	tolua.EndUserType(L)
	tolua.EndModule(L)
}

func main() {
	stack := golua.NewLuaStack()
	defer stack.Close()

	register_book(stack.L)
	stack.AddPackagePath(".")
	stack.Load("test")

	b := &book{name: "Programming in Lua"}
	stack.PushUserType(unsafe.Pointer(b), "Book")
	stack.ExecuteGlobalFunction("init_book", 1, 0)
	stack.Clean()
	log.Println("call init_book")
	// store author
	stack.PushUserType(unsafe.Pointer(b), "Book")
	stack.PushString("Roberto Ierusalimschy")
	stack.ExecuteGlobalFunction("store_author", 2, 0)
	stack.Clean()
	log.Println("call store_author")
	// load author
	stack.PushUserType(unsafe.Pointer(b), "Book")
	stack.ExecuteGlobalFunction("load_author", 1, 1)
	author := stack.ToString(-1)
	stack.Clean()
	log.Println("call load_author, author:", author)
	// callback
	stack.PushInt(6)
	stack.ExecuteFunction(b.GetSellCallback(), 1, 0)
	stack.Clean()
	log.Println("callback sell")
	// test error
	func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		log.Println("call test_error before")
		stack.ExecuteGlobalFunction("test_error", 0, 0)
		stack.Clean()
	}()
	log.Println("finish")
}

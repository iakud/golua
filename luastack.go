package golua

import (
	"github.com/iakud/golua/lua"
	"github.com/iakud/golua/tolua"

	"fmt"
	"unsafe"
)

type LuaStack struct {
	L *lua.Lua_State
}

func NewLuaStack() *LuaStack {
	L := lua.LuaL_newstate()
	lua.LuaL_openlibs(L)
	tolua.Open(L)
	stack := &LuaStack{L}
	return stack
}

func (stack *LuaStack) Close() {
	lua.Lua_close(stack.L)
}

func (stack *LuaStack) AddPackagePath(path string) {
	lua.Lua_getglobal(stack.L, lua.LUA_LOADLIBNAME)
	lua.Lua_getfield(stack.L, -1, "path")
	lua.Lua_pushfstring(stack.L, "%s;%s/?.lua", lua.Lua_tostring(stack.L, -1), path)
	lua.Lua_setfield(stack.L, -3, "path")
	lua.Lua_pop(stack.L, 2)
}

func (stack *LuaStack) Load(modname string) {
	if len(modname) == 0 {
		return
	}
	require := fmt.Sprintf("require '%v'", modname)
	stack.ExecuteString(require)
}

func (stack *LuaStack) Unload(modname string) {
	if len(modname) == 0 {
		return
	}
	lua.Lua_getglobal(stack.L, lua.LUA_LOADLIBNAME)
	lua.Lua_getfield(stack.L, -1, "loaded")
	lua.Lua_pushstring(stack.L, modname)
	lua.Lua_gettable(stack.L, -2)
	if !lua.Lua_isnil(stack.L, -1) {
		lua.Lua_pushstring(stack.L, modname)
		lua.Lua_pushnil(stack.L)
		lua.Lua_settable(stack.L, -4)
	}
	lua.Lua_pop(stack.L, 3)
}

func (L *LuaStack) Reload(modname string) {
	L.Unload(modname)
	L.Load(modname)
}

//
// push value
//
func (stack *LuaStack) PushNil() {
	lua.Lua_pushnil(stack.L)
}

func (stack *LuaStack) PushBool(value bool) {
	lua.Lua_pushboolean(stack.L, value)
}

func (stack *LuaStack) PushInt(value int) {
	lua.Lua_pushinteger(stack.L, lua.Lua_Integer(value))
}

func (stack *LuaStack) PushInt32(value int32) {
	lua.Lua_pushinteger(stack.L, lua.Lua_Integer(value))
}

func (stack *LuaStack) PushInt64(value int64) {
	lua.Lua_pushnumber(stack.L, lua.Lua_Number(value))
}

func (stack *LuaStack) PushFloat32(value float32) {
	lua.Lua_pushnumber(stack.L, lua.Lua_Number(value))
}

func (stack *LuaStack) PushFloat64(value float64) {
	lua.Lua_pushnumber(stack.L, lua.Lua_Number(value))
}

func (stack *LuaStack) PushString(value string) {
	lua.Lua_pushstring(stack.L, value)
}

func (stack *LuaStack) PushLString(value string) {
	lua.Lua_pushlstring(stack.L, value)
}

func (stack *LuaStack) PushUserType(p unsafe.Pointer, name string) {
	tolua.PushUserType(stack.L, p, name)
}

func (stack *LuaStack) PushFunctionRef(f *tolua.FunctionRef) {
	tolua.PushFunctionRef(stack.L, f)
}

//
// to value
//
func (stack *LuaStack) ToBool(index int) bool {
	return lua.Lua_toboolean(stack.L, index)
}

func (stack *LuaStack) ToInt(index int) int {
	return int(lua.Lua_tointeger(stack.L, index))
}

func (stack *LuaStack) ToInt32(index int) int32 {
	return int32(lua.Lua_tointeger(stack.L, index))
}

func (stack *LuaStack) ToInt64(index int) int64 {
	return int64(lua.Lua_tonumber(stack.L, index))
}

func (stack *LuaStack) ToFloat32(index int) float32 {
	return float32(lua.Lua_tonumber(stack.L, index))
}

func (stack *LuaStack) ToFloat64(index int) float64 {
	return float64(lua.Lua_tonumber(stack.L, index))
}

func (stack *LuaStack) ToString(index int) string {
	return lua.Lua_tostring(stack.L, index)
}

func (stack *LuaStack) ToLString(index int) string {
	return lua.Lua_tolstring(stack.L, index)
}

func (stack *LuaStack) ToUserType(index int, name string) unsafe.Pointer {
	return tolua.ToUserType(stack.L, index, name)
}

func (stack *LuaStack) ToFunctionRef(index int) *tolua.FunctionRef {
	return tolua.ToFunctionRef(stack.L, index)
}

//
// remove
//
func (stack *LuaStack) RemoveFunctionRef(f *tolua.FunctionRef) {
	tolua.RemoveFunctionRef(stack.L, f)
}

//
// stack
//
func (stack *LuaStack) GetTop() int {
	return lua.Lua_gettop(stack.L)
}

func (stack *LuaStack) Clean() {
	lua.Lua_settop(stack.L, 0)
}

func (stack *LuaStack) FormatIndex(index int) int {
	if index < 0 {
		return stack.GetTop() + 1 + index
	} else {
		return index
	}
}

//
// excute
//
func (stack *LuaStack) ExecuteGlobalFunction(funcname string, nargs, nresults int) {
	lua.Lua_getglobal(stack.L, funcname)
	if nargs > 0 {
		lua.Lua_insert(stack.L, -(nargs + 1))
	}
	stack.execute(nargs, nresults)
}

func (stack *LuaStack) ExecuteFunction(f *tolua.FunctionRef, nargs, nresults int) {
	tolua.PushFunctionRef(stack.L, f)
	if nargs > 0 {
		lua.Lua_insert(stack.L, -(nargs + 1))
	}
	stack.execute(nargs, nresults)
}

func (stack *LuaStack) ExecuteString(codes string) {
	lua.LuaL_loadstring(stack.L, codes)
	stack.execute(0, 0)
}

func (stack *LuaStack) execute(nargs, nresults int) {
	if tolua.DoCall(stack.L, nargs, nresults) != 0 && !lua.Lua_isnil(stack.L, -1) {
		err := lua.Lua_tostring(stack.L, -1)
		lua.Lua_pop(stack.L, 1)
		panic(err)
	}
}

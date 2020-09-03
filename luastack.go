package golua

import (
	"golua/lua"
	"golua/tolua"

	"fmt"
	"runtime"
	"unsafe"
)

type LuaStack struct {
	l *lua.Lua_State
}

func NewLuaStack() *LuaStack {
	L := lua.LuaL_newstate()
	lua.LuaL_openlibs(L)
	tolua.Open(L)
	stack := &LuaStack{
		l: L,
	}
	runtime.SetFinalizer(stack, (*LuaStack).Close)
	return stack
}

func (this *LuaStack) Close() {
	lua.Lua_close(this.l)
	// no need for a finalizer anymore
	runtime.SetFinalizer(this, nil)
}

func (this *LuaStack) LuaState() *lua.Lua_State {
	return this.l
}

func (this *LuaStack) AddPackagePath(path string) {
	lua.Lua_getglobal(this.l, lua.LUA_LOADLIBNAME)
	lua.Lua_getfield(this.l, -1, "path")
	lua.Lua_pushfstring(this.l, "%s;%s/?.lua", lua.Lua_tostring(this.l, -1), path)
	lua.Lua_setfield(this.l, -3, "path")
	lua.Lua_pop(this.l, 2)
}

func (this *LuaStack) Load(modname string) {
	if len(modname) == 0 {
		return
	}
	require := fmt.Sprintf("require '%v'", modname)
	this.ExecuteString(require)
}

func (this *LuaStack) Unload(modname string) {
	if len(modname) == 0 {
		return
	}
	lua.Lua_getglobal(this.l, lua.LUA_LOADLIBNAME)
	lua.Lua_getfield(this.l, -1, "loaded")
	lua.Lua_pushstring(this.l, modname)
	lua.Lua_gettable(this.l, -2)
	if !lua.Lua_isnil(this.l, -1) {
		lua.Lua_pushstring(this.l, modname)
		lua.Lua_pushnil(this.l)
		lua.Lua_settable(this.l, -4)
	}
	lua.Lua_pop(this.l, 3)
}

func (L *LuaStack) Reload(modname string) {
	L.Unload(modname)
	L.Load(modname)
}

//
// push value
//
func (this *LuaStack) PushNil() {
	lua.Lua_pushnil(this.l)
}

func (this *LuaStack) PushBool(value bool) {
	lua.Lua_pushboolean(this.l, value)
}

func (this *LuaStack) PushInt(value int) {
	lua.Lua_pushinteger(this.l, lua.Lua_Integer(value))
}

func (this *LuaStack) PushInt32(value int32) {
	lua.Lua_pushinteger(this.l, lua.Lua_Integer(value))
}

func (this *LuaStack) PushInt64(value int64) {
	lua.Lua_pushnumber(this.l, lua.Lua_Number(value))
}

func (this *LuaStack) PushFloat32(value float32) {
	lua.Lua_pushnumber(this.l, lua.Lua_Number(value))
}

func (this *LuaStack) PushFloat64(value float64) {
	lua.Lua_pushnumber(this.l, lua.Lua_Number(value))
}

func (this *LuaStack) PushString(value string) {
	lua.Lua_pushstring(this.l, value)
}

func (this *LuaStack) PushLString(value string) {
	lua.Lua_pushlstring(this.l, value)
}

func (this *LuaStack) PushUserType(p unsafe.Pointer, name string) {
	tolua.PushUserType(this.l, p, name)
}

func (this *LuaStack) PushFunctionRef(f *tolua.Tolua_FunctionRef) {
	tolua.PushFunctionRef(this.l, f)
}

//
// to value
//
func (this *LuaStack) ToBool(index int) bool {
	return lua.Lua_toboolean(this.l, index)
}

func (this *LuaStack) ToInt(index int) int {
	return int(lua.Lua_tointeger(this.l, index))
}

func (this *LuaStack) ToInt32(index int) int32 {
	return int32(lua.Lua_tointeger(this.l, index))
}

func (this *LuaStack) ToInt64(index int) int64 {
	return int64(lua.Lua_tonumber(this.l, index))
}

func (this *LuaStack) ToFloat32(index int) float32 {
	return float32(lua.Lua_tonumber(this.l, index))
}

func (this *LuaStack) ToFloat64(index int) float64 {
	return float64(lua.Lua_tonumber(this.l, index))
}

func (this *LuaStack) ToString(index int) string {
	return lua.Lua_tostring(this.l, index)
}

func (this *LuaStack) ToLString(index int) string {
	return lua.Lua_tolstring(this.l, index)
}

func (this *LuaStack) ToUserType(index int, name string) unsafe.Pointer {
	return tolua.ToUserType(this.l, index, name)
}

func (this *LuaStack) ToFunctionRef(index int) *tolua.Tolua_FunctionRef {
	return tolua.ToFunctionRef(this.l, index)
}

//
// remove
//
func (this *LuaStack) RemoveFunctionRef(f *tolua.Tolua_FunctionRef) {
	tolua.RemoveFunctionRef(this.l, f)
}

//
// stack
//
func (this *LuaStack) GetTop() int {
	return lua.Lua_gettop(this.l)
}

func (this *LuaStack) Clean() {
	lua.Lua_settop(this.l, 0)
}

func (this *LuaStack) FormatIndex(index int) int {
	if index < 0 {
		return this.GetTop() + 1 + index
	} else {
		return index
	}
}

//
// excute
//
func (this *LuaStack) ExecuteGlobalFunction(funcname string, nargs, nresults int) {
	lua.Lua_getglobal(this.l, funcname)
	if nargs > 0 {
		lua.Lua_insert(this.l, -(nargs + 1))
	}
	this.execute(nargs, nresults)
}

func (this *LuaStack) ExecuteString(codes string) {
	lua.LuaL_loadstring(this.l, codes)
	this.execute(0, 0)
}

func (this *LuaStack) execute(nargs, nresults int) {
	functionIndex := this.FormatIndex(-(nargs + 1))
	traceback := 0
	lua.Lua_getglobal(this.l, "__TRACKBACK__")
	if lua.Lua_isfunction(this.l, -1) {
		lua.Lua_insert(this.l, functionIndex)
		traceback = functionIndex
	} else {
		lua.Lua_pop(this.l, 1)
	}
	if lua.Lua_pcall(this.l, nargs, nresults, traceback) != 0 {
		err := lua.Lua_tostring(this.l, -1)
		if traceback != 0 {
			lua.Lua_pop(this.l, 2)
		} else {
			lua.Lua_pop(this.l, 1)
		}
		panic(err)
	} else if traceback != 0 {
		lua.Lua_remove(this.l, traceback)
	}
}

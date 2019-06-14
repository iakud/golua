package lua

/*
#include <lauxlib.h>
#include <stdlib.h>

#if defined(luaL_getn)
#undef luaL_getn
LUALIB_API int luaL_getn(lua_State *L, int i) { return (int)lua_objlen(L, i); }
#endif

#if defined(luaL_setn)
#undef luaL_setn
LUALIB_API void luaL_setn(lua_State *L, int i, int j) { (void)0; }
#endif

LUALIB_API int (LuaL_error) (lua_State *L, const char *s) { return luaL_error(L, s); }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func LuaL_getn(L *Lua_State, i int) int { return int(C.luaL_getn(L, C.int(i))) }
func LuaL_setn(L *Lua_State, i, j int)  { C.luaL_setn(L, C.int(i), C.int(j)) }

/* extra error code for `luaL_load' */
const LUA_ERRFILE int = LUA_ERRERR + 1

type LuaL_Reg C.luaL_Reg

func LuaI_openlib(L *Lua_State, libname string, l *LuaL_Reg, nup int) {
	c_libname := C.CString(libname)
	defer C.free(unsafe.Pointer(c_libname))
	C.luaI_openlib(L, c_libname, l, C.int(nup))
}
func LuaL_register(L *Lua_State, libname string, l *LuaL_Reg) {
	c_libname := C.CString(libname)
	defer C.free(unsafe.Pointer(c_libname))
	C.luaL_register(L, c_libname, l)
}
func LuaL_getmetafield(L *Lua_State, obj int, e string) int {
	c_e := C.CString(e)
	defer C.free(unsafe.Pointer(c_e))
	return int(C.luaL_getmetafield(L, C.int(obj), c_e))
}
func LuaL_callmeta(L *Lua_State, obj int, e string) int {
	c_e := C.CString(e)
	defer C.free(unsafe.Pointer(c_e))
	return int(C.luaL_callmeta(L, C.int(obj), c_e))
}
func LuaL_typerror(L *Lua_State, narg int, tname string) int {
	c_tname := C.CString(tname)
	defer C.free(unsafe.Pointer(c_tname))
	return int(C.luaL_typerror(L, C.int(narg), c_tname))
}
func LuaL_argerror(L *Lua_State, numarg int, extramsg string) int {
	c_extramsg := C.CString(extramsg)
	defer C.free(unsafe.Pointer(c_extramsg))
	return int(C.luaL_argerror(L, C.int(numarg), c_extramsg))
}
func LuaL_checklstring(L *Lua_State, numArg int) string {
	var l C.size_t
	c_s := C.luaL_checklstring(L, C.int(numArg), &l)
	return C.GoStringN(c_s, C.int(l))
}
func LuaL_optlstring(L *Lua_State, numArg int, def string) string {
	c_def := C.CString(def)
	defer C.free(unsafe.Pointer(c_def))
	var l C.size_t
	c_s := C.luaL_optlstring(L, C.int(numArg), c_def, &l)
	return C.GoStringN(c_s, C.int(l))
}
func LuaL_checknumber(L *Lua_State, numArg int) Lua_Number {
	return Lua_Number(C.luaL_checknumber(L, C.int(numArg)))
}
func LuaL_optnumber(L *Lua_State, numArg int, def Lua_Number) Lua_Number {
	return Lua_Number(C.luaL_optnumber(L, C.int(numArg), C.lua_Number(def)))
}

func LuaL_checkinteger(L *Lua_State, numArg int) Lua_Integer {
	return Lua_Integer(C.luaL_checkinteger(L, C.int(numArg)))
}
func LuaL_optinteger(L *Lua_State, numArg int, def Lua_Integer) Lua_Integer {
	return Lua_Integer(C.luaL_optinteger(L, C.int(numArg), C.lua_Integer(def)))
}

func LuaL_checkstack(L *Lua_State, sz int, msg string) {
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))
	C.luaL_checkstack(L, C.int(sz), c_msg)
}
func LuaL_checktype(L *Lua_State, narg, t int) { C.luaL_checktype(L, C.int(narg), C.int(t)) }
func LuaL_checkany(L *Lua_State, narg int)     { C.luaL_checkany(L, C.int(narg)) }

func LuaL_newmetatable(L *Lua_State, tname string) int {
	c_tname := C.CString(tname)
	defer C.free(unsafe.Pointer(c_tname))
	return int(C.luaL_newmetatable(L, c_tname))
}
func LuaL_checkudata(L *Lua_State, ud int, tname string) unsafe.Pointer {
	c_tname := C.CString(tname)
	defer C.free(unsafe.Pointer(c_tname))
	return C.luaL_checkudata(L, C.int(ud), c_tname)
}

func LuaL_where(L *Lua_State, lvl int) { C.luaL_where(L, C.int(lvl)) }
func LuaL_error(L *Lua_State, format string, a ...interface{}) int {
	c_fmt := C.CString(fmt.Sprintf(format, a...))
	defer C.free(unsafe.Pointer(c_fmt))
	return int(C.LuaL_error(L, c_fmt))
}

// LUALIB_API int (luaL_checkoption) (lua_State *L, int narg, const char *def, const char *const lst[]);

func LuaL_ref(L *Lua_State, t int) int    { return int(C.luaL_ref(L, C.int(t))) }
func LuaL_unref(L *Lua_State, t, ref int) { C.luaL_unref(L, C.int(t), C.int(ref)) }

func LuaL_loadfile(L *Lua_State, filename string) int {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return int(C.luaL_loadfile(L, c_filename))
}
func LuaL_loadbuffer(L *Lua_State, buff string, sz uint, name string) int {
	c_buff := C.CString(buff)
	defer C.free(unsafe.Pointer(c_buff))
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.luaL_loadbuffer(L, c_buff, C.size_t(sz), c_name))
}
func LuaL_loadstring(L *Lua_State, s string) int {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	return int(C.luaL_loadstring(L, c_s))
}

func LuaL_newstate() *Lua_State { return (*Lua_State)(C.luaL_newstate()) }

func LuaL_gsub(L *Lua_State, s, p, r string) string {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))
	c_r := C.CString(r)
	defer C.free(unsafe.Pointer(c_r))
	return C.GoString(C.luaL_gsub(L, c_s, c_p, c_r))
}

func LuaL_findtable(L *Lua_State, idx int, fname string, szhint int) string {
	c_fname := C.CString(fname)
	defer C.free(unsafe.Pointer(c_fname))
	return C.GoString(C.luaL_findtable(L, C.int(idx), c_fname, C.int(szhint)))
}

/*
** ===============================================================
** some useful macros
** ===============================================================
 */

func LuaL_dofile(L *Lua_State, fn string) int {
	if LuaL_loadfile(L, fn) != 0 {
		return 1
	}
	if Lua_pcall(L, 0, C.LUA_MULTRET, 0) != 0 {
		return 1
	}
	return 0
}

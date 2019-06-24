package lua

/*
#include <lauxlib.h>
#include <stdlib.h>

LUALIB_API int (LuaL_error) (lua_State *L, const char *s) { return luaL_error(L, s); }
LUALIB_API void (LuaL_addchar) (luaL_Buffer *B, char c) { luaL_addchar(B, c); }
LUALIB_API void (LuaL_addsize) (luaL_Buffer *B, int n) { luaL_addsize(B, n); }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

/* extra error code for `luaL_load' */
const LUA_ERRFILE int = LUA_ERRERR + 1

type LuaL_Reg C.luaL_Reg

func luaL_openlib(L *Lua_State, libname string, l *LuaL_Reg, nup int) {
	c_libname := C.CString(libname)
	defer C.free(unsafe.Pointer(c_libname))
	C.luaL_openlib((*C.lua_State)(L), c_libname, (*C.luaL_Reg)(l), C.int(nup))
}
func LuaL_register(L *Lua_State, libname string, l *LuaL_Reg) {
	c_libname := C.CString(libname)
	defer C.free(unsafe.Pointer(c_libname))
	C.luaL_register((*C.lua_State)(L), c_libname, (*C.luaL_Reg)(l))
}
func LuaL_getmetafield(L *Lua_State, obj int, e string) int {
	c_e := C.CString(e)
	defer C.free(unsafe.Pointer(c_e))
	return int(C.luaL_getmetafield((*C.lua_State)(L), C.int(obj), c_e))
}
func LuaL_callmeta(L *Lua_State, obj int, e string) int {
	c_e := C.CString(e)
	defer C.free(unsafe.Pointer(c_e))
	return int(C.luaL_callmeta((*C.lua_State)(L), C.int(obj), c_e))
}
func LuaL_typerror(L *Lua_State, narg int, tname string) int {
	c_tname := C.CString(tname)
	defer C.free(unsafe.Pointer(c_tname))
	return int(C.luaL_typerror((*C.lua_State)(L), C.int(narg), c_tname))
}
func LuaL_argerror(L *Lua_State, numarg int, extramsg string) int {
	c_extramsg := C.CString(extramsg)
	defer C.free(unsafe.Pointer(c_extramsg))
	return int(C.luaL_argerror((*C.lua_State)(L), C.int(numarg), c_extramsg))
}
func LuaL_checklstring(L *Lua_State, numArg int) string {
	var l C.size_t
	c_s := C.luaL_checklstring((*C.lua_State)(L), C.int(numArg), &l)
	return C.GoStringN(c_s, C.int(l))
}
func LuaL_optlstring(L *Lua_State, numArg int, def string) string {
	c_def := C.CString(def)
	defer C.free(unsafe.Pointer(c_def))
	var l C.size_t
	c_s := C.luaL_optlstring((*C.lua_State)(L), C.int(numArg), c_def, &l)
	return C.GoStringN(c_s, C.int(l))
}
func LuaL_checknumber(L *Lua_State, numArg int) Lua_Number {
	return Lua_Number(C.luaL_checknumber((*C.lua_State)(L), C.int(numArg)))
}
func LuaL_optnumber(L *Lua_State, numArg int, def Lua_Number) Lua_Number {
	return Lua_Number(C.luaL_optnumber((*C.lua_State)(L), C.int(numArg), C.lua_Number(def)))
}

func LuaL_checkinteger(L *Lua_State, numArg int) Lua_Integer {
	return Lua_Integer(C.luaL_checkinteger((*C.lua_State)(L), C.int(numArg)))
}
func LuaL_optinteger(L *Lua_State, numArg int, def Lua_Integer) Lua_Integer {
	return Lua_Integer(C.luaL_optinteger((*C.lua_State)(L), C.int(numArg), C.lua_Integer(def)))
}

func LuaL_checkstack(L *Lua_State, sz int, msg string) {
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))
	C.luaL_checkstack((*C.lua_State)(L), C.int(sz), c_msg)
}
func LuaL_checktype(L *Lua_State, narg, t int) {
	C.luaL_checktype((*C.lua_State)(L), C.int(narg), C.int(t))
}
func LuaL_checkany(L *Lua_State, narg int) { C.luaL_checkany((*C.lua_State)(L), C.int(narg)) }

func LuaL_newmetatable(L *Lua_State, tname string) int {
	c_tname := C.CString(tname)
	defer C.free(unsafe.Pointer(c_tname))
	return int(C.luaL_newmetatable((*C.lua_State)(L), c_tname))
}
func LuaL_checkudata(L *Lua_State, ud int, tname string) unsafe.Pointer {
	c_tname := C.CString(tname)
	defer C.free(unsafe.Pointer(c_tname))
	return C.luaL_checkudata((*C.lua_State)(L), C.int(ud), c_tname)
}

func LuaL_where(L *Lua_State, lvl int) { C.luaL_where((*C.lua_State)(L), C.int(lvl)) }
func LuaL_error(L *Lua_State, format string, a ...interface{}) int {
	c_fmt := C.CString(fmt.Sprintf(format, a...))
	defer C.free(unsafe.Pointer(c_fmt))
	return int(C.LuaL_error((*C.lua_State)(L), c_fmt))
}

// LUALIB_API int (luaL_checkoption) (lua_State *L, int narg, const char *def, const char *const lst[]);

func LuaL_ref(L *Lua_State, t int) int    { return int(C.luaL_ref((*C.lua_State)(L), C.int(t))) }
func LuaL_unref(L *Lua_State, t, ref int) { C.luaL_unref((*C.lua_State)(L), C.int(t), C.int(ref)) }

func LuaL_loadfile(L *Lua_State, filename string) int {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return int(C.luaL_loadfile((*C.lua_State)(L), c_filename))
}
func LuaL_loadbuffer(L *Lua_State, buff string, sz uint, name string) int {
	c_buff := C.CString(buff)
	defer C.free(unsafe.Pointer(c_buff))
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.luaL_loadbuffer((*C.lua_State)(L), c_buff, C.size_t(sz), c_name))
}
func LuaL_loadstring(L *Lua_State, s string) int {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	return int(C.luaL_loadstring((*C.lua_State)(L), c_s))
}

func LuaL_newstate() *Lua_State { return (*Lua_State)(C.luaL_newstate()) }

func LuaL_gsub(L *Lua_State, s, p, r string) string {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))
	c_r := C.CString(r)
	defer C.free(unsafe.Pointer(c_r))
	return C.GoString(C.luaL_gsub((*C.lua_State)(L), c_s, c_p, c_r))
}

func LuaL_findtable(L *Lua_State, idx int, fname string, szhint int) string {
	c_fname := C.CString(fname)
	defer C.free(unsafe.Pointer(c_fname))
	return C.GoString(C.luaL_findtable((*C.lua_State)(L), C.int(idx), c_fname, C.int(szhint)))
}

/*
** ===============================================================
** some useful macros
** ===============================================================
 */

func LuaL_argcheck(L *Lua_State, cond bool, numarg int, extramsg string) {
	if !cond {
		LuaL_argerror(L, numarg, extramsg)
	}
}
func LuaL_checkstring(L *Lua_State, n int) string {
	return C.GoString(C.luaL_checklstring((*C.lua_State)(L), C.int(n), nil))
}
func LuaL_optstring(L *Lua_State, n int, d string) string {
	c_d := C.CString(d)
	defer C.free(unsafe.Pointer(c_d))
	return C.GoString(C.luaL_optlstring((*C.lua_State)(L), C.int(n), c_d, nil))
}
func LuaL_checkint(L *Lua_State, n int) int   { return int(LuaL_checkinteger(L, n)) }
func LuaL_optint(L *Lua_State, n, d int) int  { return int(LuaL_optinteger(L, n, Lua_Integer(d))) }
func LuaL_checklong(L *Lua_State, n int) int  { return int(LuaL_checkinteger(L, n)) }
func LuaL_optlong(L *Lua_State, n, d int) int { return int(LuaL_optinteger(L, n, Lua_Integer(d))) }

func LuaL_typename(L *Lua_State, i int) string { return Lua_typename(L, Lua_type(L, i)) }

func LuaL_dofile(L *Lua_State, fn string) int {
	if ret := LuaL_loadfile(L, fn); ret != 0 {
		return ret
	}
	return Lua_pcall(L, 0, C.LUA_MULTRET, 0)
}
func LuaL_dostring(L *Lua_State, s string) int {
	if ret := LuaL_loadstring(L, s); ret != 0 {
		return ret
	}
	return Lua_pcall(L, 0, C.LUA_MULTRET, 0)
}

func LuaL_getmetatable(L *Lua_State, n string) { Lua_getfield(L, LUA_REGISTRYINDEX, n) }

// #define luaL_opt(L,f,n,d)	(lua_isnoneornil(L,(n)) ? (d) : f(L,(n)))

/*
** {======================================================
** Generic Buffer manipulation
** =======================================================
 */

type LuaL_Buffer C.luaL_Buffer

func LuaL_addchar(B *LuaL_Buffer, c byte) { C.LuaL_addchar((*C.luaL_Buffer)(B), C.char(c)) }

/* compatibility only */
func LuaL_putchar(B *LuaL_Buffer, c byte) { LuaL_addchar(B, c) }

func LuaL_addsize(B *LuaL_Buffer, n int) { C.LuaL_addsize((*C.luaL_Buffer)(B), C.int(n)) }

func LuaL_buffinit(L *Lua_State, B *LuaL_Buffer) {
	C.luaL_buffinit((*C.lua_State)(L), (*C.luaL_Buffer)(B))
}
func LuaL_prepbuffer(B *LuaL_Buffer) string { return C.GoString(C.luaL_prepbuffer((*C.luaL_Buffer)(B))) }
func LuaL_addlstring(B *LuaL_Buffer, s string) {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	C.luaL_addlstring((*C.luaL_Buffer)(B), c_s, C.size_t(len(s)))
}
func LuaL_addstring(B *LuaL_Buffer, s string) {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	C.luaL_addstring((*C.luaL_Buffer)(B), c_s)
}
func LuaL_addvalue(B *LuaL_Buffer)   { C.luaL_addvalue((*C.luaL_Buffer)(B)) }
func LuaL_pushresult(B *LuaL_Buffer) { C.luaL_pushresult((*C.luaL_Buffer)(B)) }

/* }====================================================== */

/* compatibility with ref system */

/* pre-defined references */
const LUA_NOREF int = C.LUA_NOREF
const LUA_REFNIL int = C.LUA_REFNIL

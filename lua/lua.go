package lua

/*
#include <lua.h>
#include <stdlib.h>

LUA_API const char *(Lua_pushfstring) (lua_State *L, const char *s) { return lua_pushfstring(L, s); }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	LUA_VERSION     string = C.LUA_VERSION
	LUA_RELEASE     string = C.LUA_RELEASE
	LUA_VERSION_NUM int    = C.LUA_VERSION_NUM
	LUA_COPYRIGHT   string = C.LUA_COPYRIGHT
	LUA_AUTHORS     string = C.LUA_AUTHORS
)

/* mark for precompiled code (`<esc>Lua') */
const LUA_SIGNATURE string = C.LUA_SIGNATURE

/* option for multiple returns in `lua_pcall' and `lua_call' */
const LUA_MULTRET int = C.LUA_MULTRET

/*
** pseudo-indices
 */
const (
	LUA_REGISTRYINDEX int = C.LUA_REGISTRYINDEX
	LUA_ENVIRONINDEX  int = C.LUA_ENVIRONINDEX
	LUA_GLOBALSINDEX  int = C.LUA_GLOBALSINDEX
)

func Lua_upvalueindex(i int) int { return LUA_GLOBALSINDEX - i }

/* thread status; 0 is OK */
const (
	LUA_YIELD     int = C.LUA_YIELD
	LUA_ERRRUN    int = C.LUA_ERRRUN
	LUA_ERRSYNTAX int = C.LUA_ERRSYNTAX
	LUA_ERRMEM    int = C.LUA_ERRMEM
	LUA_ERRERR    int = C.LUA_ERRERR
)

type Lua_State C.lua_State

type Lua_CFunction C.lua_CFunction

/*
** functions that read/write blocks when loading/dumping Lua chunks
 */
type Lua_Reader C.lua_Reader

type Lua_Writer C.lua_Writer

/*
** prototype for memory-allocation functions
 */
type Lua_Alloc C.lua_Alloc

/*
** basic types
 */
const (
	LUA_TNONE int = C.LUA_TNONE

	LUA_TNIL           int = C.LUA_TNIL
	LUA_TBOOLEAN       int = C.LUA_TBOOLEAN
	LUA_TLIGHTUSERDATA int = C.LUA_TLIGHTUSERDATA
	LUA_TNUMBER        int = C.LUA_TNUMBER
	LUA_TSTRING        int = C.LUA_TSTRING
	LUA_TTABLE         int = C.LUA_TTABLE
	LUA_TFUNCTION      int = C.LUA_TFUNCTION
	LUA_TUSERDATA      int = C.LUA_TUSERDATA
	LUA_TTHREAD        int = C.LUA_TTHREAD
)

/* minimum Lua stack available to a C function */
const LUA_MINSTACK int = C.LUA_MINSTACK

/* type of numbers in Lua */
type Lua_Number float64

/* type for integer functions */
type Lua_Integer int

/*
** state manipulation
 */
func Lua_newstate(f Lua_Alloc, ud unsafe.Pointer) *Lua_State {
	return (*Lua_State)(C.lua_newstate(f, ud))
}
func Lua_close(L *Lua_State)                { C.lua_close((*C.lua_State)(L)) }
func Lua_newthread(L *Lua_State) *Lua_State { return (*Lua_State)(C.lua_newthread((*C.lua_State)(L))) }

func Lua_atpanic(L *Lua_State, panicf Lua_CFunction) Lua_CFunction {
	return Lua_CFunction(C.lua_atpanic((*C.lua_State)(L), panicf))
}

/*
** basic stack manipulation
 */
func Lua_gettop(L *Lua_State) int         { return int(C.lua_gettop((*C.lua_State)(L))) }
func Lua_settop(L *Lua_State, idx int)    { C.lua_settop((*C.lua_State)(L), C.int(idx)) }
func Lua_pushvalue(L *Lua_State, idx int) { C.lua_pushvalue((*C.lua_State)(L), C.int(idx)) }
func Lua_remove(L *Lua_State, idx int)    { C.lua_remove((*C.lua_State)(L), C.int(idx)) }
func Lua_insert(L *Lua_State, idx int)    { C.lua_insert((*C.lua_State)(L), C.int(idx)) }
func Lua_replace(L *Lua_State, idx int)   { C.lua_replace((*C.lua_State)(L), C.int(idx)) }
func Lua_checkstack(L *Lua_State, sz int) int {
	return int(C.lua_checkstack((*C.lua_State)(L), C.int(sz)))
}

func Lua_xmove(from, to *Lua_State, n int) {
	C.lua_xmove((*C.lua_State)(from), (*C.lua_State)(to), C.int(n))
}

/*
** access functions (stack -> C)
 */
func Lua_isnumber(L *Lua_State, idx int) bool {
	return C.lua_isnumber((*C.lua_State)(L), C.int(idx)) != 0
}
func Lua_isstring(L *Lua_State, idx int) bool {
	return C.lua_isstring((*C.lua_State)(L), C.int(idx)) != 0
}
func Lua_iscfunction(L *Lua_State, idx int) bool {
	return C.lua_iscfunction((*C.lua_State)(L), C.int(idx)) != 0
}
func Lua_isuserdata(L *Lua_State, idx int) bool {
	return C.lua_isuserdata((*C.lua_State)(L), C.int(idx)) != 0
}
func Lua_type(L *Lua_State, idx int) int { return int(C.lua_type((*C.lua_State)(L), C.int(idx))) }
func Lua_typename(L *Lua_State, tp int) string {
	return C.GoString(C.lua_typename((*C.lua_State)(L), C.int(tp)))
}

func Lua_equal(L *Lua_State, idx1, idx2 int) int {
	return int(C.lua_equal((*C.lua_State)(L), C.int(idx1), C.int(idx2)))
}
func Lua_rawequal(L *Lua_State, idx1, idx2 int) int {
	return int(C.lua_rawequal((*C.lua_State)(L), C.int(idx1), C.int(idx2)))
}
func Lua_lessthan(L *Lua_State, idx1, idx2 int) int {
	return int(C.lua_lessthan((*C.lua_State)(L), C.int(idx1), C.int(idx2)))
}

func Lua_tonumber(L *Lua_State, idx int) Lua_Number {
	return Lua_Number(C.lua_tonumber((*C.lua_State)(L), C.int(idx)))
}
func Lua_tointeger(L *Lua_State, idx int) Lua_Integer {
	return Lua_Integer(C.lua_tointeger((*C.lua_State)(L), C.int(idx)))
}
func Lua_toboolean(L *Lua_State, idx int) bool {
	return C.lua_toboolean((*C.lua_State)(L), C.int(idx)) != 0
}
func Lua_tolstring(L *Lua_State, idx int) string {
	var l C.size_t
	c_s := C.lua_tolstring((*C.lua_State)(L), C.int(idx), &l)
	return C.GoStringN(c_s, C.int(l))
}
func Lua_objlen(L *Lua_State, idx int) uint { return uint(C.lua_objlen((*C.lua_State)(L), C.int(idx))) }
func Lua_tocfunction(L *Lua_State, idx int) Lua_CFunction {
	return Lua_CFunction(C.lua_tocfunction((*C.lua_State)(L), C.int(idx)))
}
func Lua_touserdata(L *Lua_State, idx int) unsafe.Pointer {
	return unsafe.Pointer(C.lua_touserdata((*C.lua_State)(L), C.int(idx)))
}
func Lua_tothread(L *Lua_State, idx int) *Lua_State {
	return (*Lua_State)(C.lua_tothread((*C.lua_State)(L), C.int(idx)))
}
func Lua_topointer(L *Lua_State, idx int) unsafe.Pointer {
	return unsafe.Pointer(C.lua_topointer((*C.lua_State)(L), C.int(idx)))
}

/*
** push functions (C -> stack)
 */
func Lua_pushnil(L *Lua_State)                  { C.lua_pushnil((*C.lua_State)(L)) }
func Lua_pushnumber(L *Lua_State, n Lua_Number) { C.lua_pushnumber((*C.lua_State)(L), C.lua_Number(n)) }
func Lua_pushinteger(L *Lua_State, n Lua_Integer) {
	C.lua_pushinteger((*C.lua_State)(L), C.lua_Integer(n))
}
func Lua_pushlstring(L *Lua_State, s string) {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	C.lua_pushlstring((*C.lua_State)(L), c_s, C.size_t(len(s)))
}
func Lua_pushstring(L *Lua_State, s string) {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))
	C.lua_pushstring((*C.lua_State)(L), c_s)
}
func Lua_pushfstring(L *Lua_State, format string, a ...interface{}) string {
	c_s := C.CString(fmt.Sprintf(format, a...))
	defer C.free(unsafe.Pointer(c_s))
	return C.GoString(C.Lua_pushfstring((*C.lua_State)(L), c_s))
}
func Lua_pushcclosure(L *Lua_State, f Lua_CFunction, n int) {
	C.lua_pushcclosure((*C.lua_State)(L), f, C.int(n))
}
func Lua_pushboolean(L *Lua_State, b bool) {
	if b {
		C.lua_pushboolean((*C.lua_State)(L), 1)
	} else {
		C.lua_pushboolean((*C.lua_State)(L), 0)
	}
}
func Lua_pushlightuserdata(L *Lua_State, p unsafe.Pointer) {
	C.lua_pushlightuserdata((*C.lua_State)(L), p)
}
func Lua_pushthread(L *Lua_State) int { return int(C.lua_pushthread((*C.lua_State)(L))) }

/*
** get functions (Lua -> stack)
 */
func Lua_gettable(L *Lua_State, idx int) { C.lua_gettable((*C.lua_State)(L), C.int(idx)) }
func Lua_getfield(L *Lua_State, idx int, k string) {
	c_k := C.CString(k)
	defer C.free(unsafe.Pointer(c_k))
	C.lua_getfield((*C.lua_State)(L), C.int(idx), c_k)
}
func Lua_rawget(L *Lua_State, idx int)     { C.lua_rawget((*C.lua_State)(L), C.int(idx)) }
func Lua_rawgeti(L *Lua_State, idx, n int) { C.lua_rawgeti((*C.lua_State)(L), C.int(idx), C.int(n)) }
func Lua_createtable(L *Lua_State, idx, nrec int) {
	C.lua_createtable((*C.lua_State)(L), C.int(idx), C.int(nrec))
}
func Lua_newuserdata(L *Lua_State, sz uint) unsafe.Pointer {
	return C.lua_newuserdata((*C.lua_State)(L), C.size_t(sz))
}
func Lua_getmetatable(L *Lua_State, objindex int) int {
	return int(C.lua_getmetatable((*C.lua_State)(L), C.int(objindex)))
}
func Lua_getfenv(L *Lua_State, idx int) { C.lua_getfenv((*C.lua_State)(L), C.int(idx)) }

/*
** set functions (stack -> Lua)
 */
func Lua_settable(L *Lua_State, idx int) { C.lua_settable((*C.lua_State)(L), C.int(idx)) }
func Lua_setfield(L *Lua_State, idx int, k string) {
	c_k := C.CString(k)
	defer C.free(unsafe.Pointer(c_k))
	C.lua_setfield((*C.lua_State)(L), C.int(idx), c_k)
}
func Lua_rawset(L *Lua_State, idx int)     { C.lua_rawset((*C.lua_State)(L), C.int(idx)) }
func Lua_rawseti(L *Lua_State, idx, n int) { C.lua_rawseti((*C.lua_State)(L), C.int(idx), C.int(n)) }
func Lua_setmetatable(L *Lua_State, objindex int) int {
	return int(C.lua_setmetatable((*C.lua_State)(L), C.int(objindex)))
}
func Lua_setfenv(L *Lua_State, idx int) int { return int(C.lua_setfenv((*C.lua_State)(L), C.int(idx))) }

/*
** `load' and `call' functions (load and run Lua code)
 */
func Lua_call(L *Lua_State, nargs int, nresults int) {
	C.lua_call((*C.lua_State)(L), C.int(nargs), C.int(nresults))
}
func Lua_pcall(L *Lua_State, nargs int, nresults int, errfunc int) int {
	return int(C.lua_pcall((*C.lua_State)(L), C.int(nargs), C.int(nresults), C.int(errfunc)))
}
func Lua_cpcall(L *Lua_State, f Lua_CFunction, ud unsafe.Pointer) int {
	return int(C.lua_cpcall((*C.lua_State)(L), f, ud))
}
func Lua_load(L *Lua_State, reader Lua_Reader, dt unsafe.Pointer, chunkname string) int {
	c_chunkname := C.CString(chunkname)
	defer C.free(unsafe.Pointer(c_chunkname))
	return int(C.lua_load((*C.lua_State)(L), reader, dt, c_chunkname))
}

func Lua_dump(L *Lua_State, writer Lua_Writer, data unsafe.Pointer) int {
	return int(C.lua_dump((*C.lua_State)(L), writer, data))
}

/*
** coroutine functions
 */
func Lua_yield(L *Lua_State, nresults int) int {
	return int(C.lua_yield((*C.lua_State)(L), C.int(nresults)))
}
func Lua_resume(L *Lua_State, narg int) int { return int(C.lua_resume((*C.lua_State)(L), C.int(narg))) }
func Lua_status(L *Lua_State) int           { return int(C.lua_status((*C.lua_State)(L))) }

/*
** garbage-collection function and options
 */
const (
	LUA_GCSTOP       int = C.LUA_GCSTOP
	LUA_GCRESTART    int = C.LUA_GCRESTART
	LUA_GCCOLLECT    int = C.LUA_GCCOLLECT
	LUA_GCCOUNT      int = C.LUA_GCCOUNT
	LUA_GCCOUNTB     int = C.LUA_GCCOUNTB
	LUA_GCSTEP       int = C.LUA_GCSTEP
	LUA_GCSETPAUSE   int = C.LUA_GCSETPAUSE
	LUA_GCSETSTEPMUL int = C.LUA_GCSETSTEPMUL
)

func Lua_gc(L *Lua_State, what, data int) int {
	return int(C.lua_gc((*C.lua_State)(L), C.int(what), C.int(data)))
}

/*
** miscellaneous functions
 */
func Lua_error(L *Lua_State) int { return int(C.lua_error((*C.lua_State)(L))) }

func Lua_next(L *Lua_State, idx int) int { return int(C.lua_next((*C.lua_State)(L), C.int(idx))) }

func Lua_concat(L *Lua_State, n int) { C.lua_concat((*C.lua_State)(L), C.int(n)) }

func Lua_getallocf(L *Lua_State, ud *unsafe.Pointer) Lua_Alloc {
	return Lua_Alloc(C.lua_getallocf((*C.lua_State)(L), ud))
}
func Lua_setallocf(L *Lua_State, f Lua_Alloc, ud unsafe.Pointer) {
	C.lua_setallocf((*C.lua_State)(L), f, ud)
}

/*
** ===============================================================
** some useful macros
** ===============================================================
 */
func Lua_pop(L *Lua_State, n int) { Lua_settop(L, -(n)-1) }

func Lua_newtable(L *Lua_State) { Lua_createtable(L, 0, 0) }

func Lua_register(L *Lua_State, n string, f Lua_CFunction) {
	Lua_pushcfunction(L, f)
	Lua_setglobal(L, n)
}

func Lua_pushcfunction(L *Lua_State, f Lua_CFunction) { Lua_pushcclosure(L, f, 0) }

func Lua_strlen(L *Lua_State, i int) uint { return Lua_objlen(L, i) }

func Lua_isfunction(L *Lua_State, n int) bool      { return Lua_type(L, n) == LUA_TFUNCTION }
func Lua_istable(L *Lua_State, n int) bool         { return Lua_type(L, n) == LUA_TTABLE }
func Lua_islightuserdata(L *Lua_State, n int) bool { return Lua_type(L, n) == LUA_TLIGHTUSERDATA }
func Lua_isnil(L *Lua_State, n int) bool           { return Lua_type(L, n) == LUA_TNIL }
func Lua_isboolean(L *Lua_State, n int) bool       { return Lua_type(L, n) == LUA_TBOOLEAN }
func Lua_isthread(L *Lua_State, n int) bool        { return Lua_type(L, n) == LUA_TTHREAD }
func Lua_isnone(L *Lua_State, n int) bool          { return Lua_type(L, n) == LUA_TNONE }
func Lua_isnoneornil(L *Lua_State, n int) bool     { return Lua_type(L, n) <= 0 }

func Lua_pushliteral(L *Lua_State, s string) { Lua_pushlstring(L, s) }

func Lua_setglobal(L *Lua_State, s string) { Lua_setfield(L, LUA_GLOBALSINDEX, s) }
func Lua_getglobal(L *Lua_State, s string) { Lua_getfield(L, LUA_GLOBALSINDEX, s) }

func Lua_tostring(L *Lua_State, i int) string {
	return C.GoString(C.lua_tolstring((*C.lua_State)(L), C.int(i), nil))
}

/*
** compatibility macros and functions
 */

func Lua_open() *Lua_State { return LuaL_newstate() }

func Lua_getregistry(L *Lua_State) { Lua_pushvalue(L, LUA_REGISTRYINDEX) }

func Lua_getgccount(L *Lua_State) int { return Lua_gc(L, LUA_GCCOUNT, 0) }

type Lua_Chunkreader C.lua_Reader
type Lua_Chunkwriter C.lua_Writer

/* hack */
// func Lua_setlevel(from, to *Lua_State) { C.lua_setlevel(from, to) }

/*
** {======================================================================
** Debug API
** =======================================================================
 */

/*
** Event codes
 */
const (
	LUA_HOOKCALL    int = C.LUA_HOOKCALL
	LUA_HOOKRET     int = C.LUA_HOOKRET
	LUA_HOOKLINE    int = C.LUA_HOOKLINE
	LUA_HOOKCOUNT   int = C.LUA_HOOKCOUNT
	LUA_HOOKTAILRET int = C.LUA_HOOKTAILRET
)

/*
** Event masks
 */
const (
	LUA_MASKCALL  int = C.LUA_MASKCALL
	LUA_MASKRET   int = C.LUA_MASKRET
	LUA_MASKLINE  int = C.LUA_MASKLINE
	LUA_MASKCOUNT int = C.LUA_MASKCOUNT
)

type Lua_Debug C.lua_Debug /* activation record */

/* Functions to be called by the debuger in specific events */
type Lua_Hook C.lua_Hook

func Lua_getstack(L *Lua_State, level int, ar *Lua_Debug) int {
	return int(C.lua_getstack((*C.lua_State)(L), C.int(level), (*C.lua_Debug)(ar)))
}
func Lua_getinfo(L *Lua_State, what string, ar *Lua_Debug) int {
	c_what := C.CString(what)
	defer C.free(unsafe.Pointer(c_what))
	return int(C.lua_getinfo((*C.lua_State)(L), c_what, (*C.lua_Debug)(ar)))
}
func Lua_getlocal(L *Lua_State, ar *Lua_Debug, n int) string {
	return C.GoString(C.lua_getlocal((*C.lua_State)(L), (*C.lua_Debug)(ar), C.int(n)))
}
func Lua_setlocal(L *Lua_State, ar *Lua_Debug, n int) string {
	return C.GoString(C.lua_setlocal((*C.lua_State)(L), (*C.lua_Debug)(ar), C.int(n)))
}
func Lua_getupvalue(L *Lua_State, funcindex int, n int) string {
	return C.GoString(C.lua_getupvalue((*C.lua_State)(L), C.int(funcindex), C.int(n)))
}
func Lua_setupvalue(L *Lua_State, funcindex int, n int) string {
	return C.GoString(C.lua_setupvalue((*C.lua_State)(L), C.int(funcindex), C.int(n)))
}

func Lua_sethook(L *Lua_State, f Lua_Hook, mask int, count int) int {
	return int(C.lua_sethook((*C.lua_State)(L), C.lua_Hook(f), C.int(mask), C.int(count)))
}
func Lua_gethook(L *Lua_State) Lua_Hook { return Lua_Hook(C.lua_gethook((*C.lua_State)(L))) }
func lua_gethookmask(L *Lua_State) int  { return int(C.lua_gethookmask((*C.lua_State)(L))) }
func lua_gethookcount(L *Lua_State) int { return int(C.lua_gethookcount((*C.lua_State)(L))) }

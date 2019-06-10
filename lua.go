package lua

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lluajit -lmingwex

#include <lua.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

var (
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

type Lua_CFunction uintptr // func(L *Lua_State) int

/*
** functions that read/write blocks when loading/dumping Lua chunks
 */
// typedef const char * (*lua_Reader) (lua_State *L, void *ud, size_t *sz);

// typedef int (*lua_Writer) (lua_State *L, const void* p, size_t sz, void* ud);

/*
** prototype for memory-allocation functions
 */
// typedef void * (*lua_Alloc) (void *ud, void *ptr, size_t osize, size_t nsize);

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
type LUA_NUMBER Lua_Number

/* type for integer functions */
type LUA_INTEGER Lua_Integer

/*
** state manipulation
 */
// LUA_API lua_State *(lua_newstate) (lua_Alloc f, void *ud);
func Lua_close(L *Lua_State)                { C.lua_close(L) }
func Lua_newthread(L *Lua_State) *Lua_State { return (*Lua_State)(C.lua_newthread(L)) }

func Lua_atpanic(L *Lua_State, panicf Lua_CFunction) Lua_CFunction {
	return Lua_CFunction(unsafe.Pointer(C.lua_atpanic(L, (C.lua_CFunction)(unsafe.Pointer(panicf)))))
}

/*
** basic stack manipulation
 */
func Lua_gettop(L *Lua_State) int             { return int(C.lua_gettop(L)) }
func Lua_settop(L *Lua_State, idx int)        { C.lua_settop(L, C.int(idx)) }
func Lua_pushvalue(L *Lua_State, idx int)     { C.lua_pushvalue(L, C.int(idx)) }
func Lua_remove(L *Lua_State, idx int)        { C.lua_remove(L, C.int(idx)) }
func Lua_insert(L *Lua_State, idx int)        { C.lua_insert(L, C.int(idx)) }
func Lua_replace(L *Lua_State, idx int)       { C.lua_replace(L, C.int(idx)) }
func Lua_checkstack(L *Lua_State, sz int) int { return int(C.lua_checkstack(L, C.int(sz))) }

func Lua_xmove(from, to *Lua_State, n int) { C.lua_xmove(from, to, C.int(n)) }

/*
** access functions (stack -> C)
 */
func Lua_isnumber(L *Lua_State, idx int) int    { return int(C.lua_isnumber(L, C.int(idx))) }
func Lua_isstring(L *Lua_State, idx int) int    { return int(C.lua_isstring(L, C.int(idx))) }
func Lua_iscfunction(L *Lua_State, idx int) int { return int(C.lua_iscfunction(L, C.int(idx))) }
func Lua_isuserdata(L *Lua_State, idx int) int  { return int(C.lua_isuserdata(L, C.int(idx))) }
func Lua_type(L *Lua_State, idx int) int        { return int(C.lua_type(L, C.int(idx))) }
func Lua_typename(L *Lua_State, tp int) string  { return C.GoString(C.lua_typename(L, C.int(tp))) }

func Lua_equal(L *Lua_State, idx1, idx2 int) int { return int(C.lua_equal(L, C.int(idx1), C.int(idx2))) }
func Lua_rawequal(L *Lua_State, idx1, idx2 int) int {
	return int(C.lua_rawequal(L, C.int(idx1), C.int(idx2)))
}
func Lua_lessthan(L *Lua_State, idx1, idx2 int) int {
	return int(C.lua_lessthan(L, C.int(idx1), C.int(idx2)))
}

func Lua_tonumber(L *Lua_State, idx int) Lua_Number { return Lua_Number(C.lua_tonumber(L, C.int(idx))) }
func Lua_tointeger(L *Lua_State, idx int) Lua_Integer {
	return Lua_Integer(C.lua_tointeger(L, C.int(idx)))
}
func Lua_toboolean(L *Lua_State, idx int) bool { return C.lua_toboolean(L, C.int(idx)) != 0 }
func Lua_tolstring(L *Lua_State, idx int) string {
	var len C.size_t
	s := C.lua_tolstring(L, C.int(idx), &len)
	return C.GoStringN(s, C.int(len))
}
func Lua_objlen(L *Lua_State, idx int) uint { return uint(C.lua_objlen(L, C.int(idx))) }
func Lua_tocfunction(L *Lua_State, idx int) Lua_CFunction {
	return Lua_CFunction(unsafe.Pointer(C.lua_tocfunction(L, C.int(idx))))
}
func Lua_touserdata(L *Lua_State, idx int) unsafe.Pointer {
	return unsafe.Pointer(C.lua_touserdata(L, C.int(idx)))
}
func Lua_tothread(L *Lua_State, idx int) *Lua_State {
	return (*Lua_State)(C.lua_tothread(L, C.int(idx)))
}
func Lua_topointer(L *Lua_State, idx int) unsafe.Pointer {
	return unsafe.Pointer(C.lua_topointer(L, C.int(idx)))
}

/*
** push functions (C -> stack)
 */
func Lua_pushnil(L *Lua_State)               { C.lua_pushnil(L) }
func Lua_pushnumber(L *Lua_State, n float64) { C.lua_pushnumber(L, (C.lua_Number)(n)) }
func Lua_pushinteger(L *Lua_State, n int)    { C.lua_pushinteger(L, (C.lua_Integer)(n)) }
func Lua_pushlstring(L *Lua_State, s string, l uint) {
	cs := C.CString(s)
	defer C.free(cs)
	C.lua_pushlstring(L, cs, C.size_t(l))
}
func Lua_pushstring(L *Lua_State, s string) {
	cs := C.CString(s)
	defer C.free(cs)
	C.lua_pushstring(L, cs)
}

func Lua_pushcclosure(L *Lua_State, fn Lua_CFunction, n int) {
	C.lua_pushcclosure(L, C.lua_CFunction(unsafe.Pointer(fn)), C.int(n))
}
func Lua_pushboolean(L *Lua_State, b bool) {
	if b {
		C.lua_pushboolean(L, 1)
	} else {
		C.lua_pushboolean(L, 0)
	}
}
func Lua_pushlightuserdata(L *Lua_State, p unsafe.Pointer) {
	C.lua_pushlightuserdata(L, p)
}
func Lua_pushthread(L *Lua_State) int {
	return int(C.lua_pushthread(L))
}

/*
** get functions (Lua -> stack)
 */
func Lua_gettable(L *Lua_State, idx int)           { C.lua_gettable(L, C.int(idx)) }
func Lua_getfield(L *Lua_State, idx int, k string) { C.lua_getfield(L, C.int(idx), C.CString(k)) }
func Lua_rawget(L *Lua_State, idx int)             { C.lua_rawget(L, C.int(idx)) }
func Lua_rawgeti(L *Lua_State, idx, n int)         { C.lua_rawgeti(L, C.int(idx), C.int(n)) }
func Lua_createtable(L *Lua_State, idx, nrec int)  { C.lua_createtable(L, C.int(idx), C.int(nrec)) }
func Lua_newuserdata(L *Lua_State, sz uint) unsafe.Pointer {
	return unsafe.Pointer(C.lua_newuserdata(L, C.size_t(sz)))
}
func Lua_getmetatable(L *Lua_State, objindex int) int {
	return int(C.lua_getmetatable(L, C.int(objindex)))
}
func Lua_getfenv(L *Lua_State, idx int) { C.lua_getfenv(L, C.int(idx)) }

/*
** set functions (stack -> Lua)
 */
func Lua_settable(L *Lua_State, idx int) { C.lua_settable(L, C.int(idx)) }
func Lua_setfield(L *Lua_State, idx int, k string) {
	ck := C.CString(k)
	defer C.free(ck)
	C.lua_setfield(L, C.int(idx), ck)
}
func Lua_rawset(L *Lua_State, idx int)     { C.lua_rawset(L, C.int(idx)) }
func Lua_rawseti(L *Lua_State, idx, n int) { C.lua_rawseti(L, C.int(idx), C.int(n)) }
func Lua_setmetatable(L *Lua_State, objindex int) int {
	return int(C.lua_setmetatable(L, C.int(objindex)))
}
func Lua_setfenv(L *Lua_State, idx int) int { return int(C.lua_setfenv(L, C.int(idx))) }

/*
** `load' and `call' functions (load and run Lua code)
 */
func Lua_call(L *Lua_State, nargs int, nresults int) { C.lua_call(L, C.int(nargs), C.int(nresults)) }
func Lua_pcall(L *Lua_State, nargs int, nresults int, errfunc int) int {
	return int(C.lua_pcall(L, C.int(nargs), C.int(nresults), C.int(errfunc)))
}
func Lua_cpcall(L *Lua_State, f Lua_CFunction, ud unsafe.Pointer) int {
	return int(C.lua_cpcall(L, C.lua_CFunction(unsafe.Pointer(f)), ud))
}

// LUA_API int   (lua_load) (lua_State *L, lua_Reader reader, void *dt, const char *chunkname);
// LUA_API int (lua_dump) (lua_State *L, lua_Writer writer, void *data);

/*
** coroutine functions
 */
func Lua_yield(L *Lua_State, nresults int) int { return int(C.lua_yield(L, C.int(nresults))) }
func Lua_resume(L *Lua_State, narg int) int    { return int(C.lua_resume(L, C.int(narg))) }
func Lua_status(L *Lua_State) int              { return int(C.lua_status(L)) }

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

func Lua_gc(L *Lua_State, what, data int) int { return int(C.lua_gc(L, C.int(what), C.int(data))) }

/*
** miscellaneous functions
 */
func Lua_error(L *Lua_State) int { return int(C.lua_error(L)) }

func Lua_next(L *Lua_State, idx int) int { return int(C.lua_next(L, C.int(idx))) }

func Lua_concat(L *Lua_State, n int) { C.lua_concat(L, C.int(n)) }

// LUA_API lua_Alloc (lua_getallocf) (lua_State *L, void **ud);
// LUA_API void lua_setallocf (lua_State *L, lua_Alloc f, void *ud);

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

func Lua_pushliteral(L *Lua_State, s string) { Lua_pushlstring(L, s, uint(len(s))) }

func Lua_setglobal(L *Lua_State, s string) { Lua_setfield(L, LUA_GLOBALSINDEX, s) }
func Lua_getglobal(L *Lua_State, s string) { Lua_getfield(L, LUA_GLOBALSINDEX, s) }

func Lua_tostring(L *Lua_State, idx int) string {
	return C.GoString(C.lua_tolstring(L, C.int(idx), nil))
}

/*
** compatibility macros and functions
 */

func Lua_open() *Lua_State { return LuaL_newstate() }

func Lua_getregistry(L *Lua_State) { Lua_pushvalue(L, LUA_REGISTRYINDEX) }

func Lua_getgccount(L *Lua_State) int { return Lua_gc(L, LUA_GCCOUNT, 0) }

// #define lua_Chunkreader		lua_Reader
// #define lua_Chunkwriter		lua_Writer

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

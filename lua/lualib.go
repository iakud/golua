package lua

/*
#include <lualib.h>

void Lua_assert(int x) { lua_assert(x); }
*/
import "C"

/* Key to file-handle type */
const LUA_FILEHANDLE string = C.LUA_FILEHANDLE

const LUA_COLIBNAME string = C.LUA_COLIBNAME

func Luaopen_base(L *Lua_State) int { return int(C.luaopen_base((*C.lua_State)(L))) }

const LUA_TABLIBNAME string = C.LUA_TABLIBNAME

func Luaopen_table(L *Lua_State) int { return int(C.luaopen_table((*C.lua_State)(L))) }

const LUA_IOLIBNAME string = C.LUA_IOLIBNAME

func Luaopen_io(L *Lua_State) int { return int(C.luaopen_io((*C.lua_State)(L))) }

const LUA_OSLIBNAME string = C.LUA_OSLIBNAME

func Luaopen_os(L *Lua_State) int { return int(C.luaopen_os((*C.lua_State)(L))) }

const LUA_STRLIBNAME string = C.LUA_STRLIBNAME

func Luaopen_string(L *Lua_State) int { return int(C.luaopen_string((*C.lua_State)(L))) }

const LUA_MATHLIBNAME string = C.LUA_MATHLIBNAME

func Luaopen_math(L *Lua_State) int { return int(C.luaopen_math((*C.lua_State)(L))) }

const LUA_DBLIBNAME string = C.LUA_DBLIBNAME

func Luaopen_debug(L *Lua_State) int { return int(C.luaopen_debug((*C.lua_State)(L))) }

const LUA_LOADLIBNAME string = C.LUA_LOADLIBNAME

func Luaopen_package(L *Lua_State) int { return int(C.luaopen_package((*C.lua_State)(L))) }

/* open all previous libraries */
func LuaL_openlibs(L *Lua_State) { C.luaL_openlibs((*C.lua_State)(L)) }

func Lua_assert(x int) { C.Lua_assert(C.int(x)) }

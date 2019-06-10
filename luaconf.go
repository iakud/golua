package lua

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lluajit -lmingwex

#include <lua.h>
*/
type Lua_Number float64
type Lua_Integer int

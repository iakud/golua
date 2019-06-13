package main

/*
#include "lua.h"
#include <stdio.h>

// The gateway function
int goAdd_cgo(lua_State* L) {
	printf("C.goAdd_cgo(): called\n");
	int goAdd(lua_State* L);
	return goAdd(L);
}
*/
import "C"

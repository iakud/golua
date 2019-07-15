function setfunc(a)
	-- 支持userdata直接get/set
	a.number = 1993
	a.string = "tolua string"
end

function getfunc(a)
	return a.number, a.string
end

function showmessage(a)
	print(a:getMessage())
end

-- 调用堆栈
__TRACKBACK__ = function(msg)
    return debug.traceback(msg, 3)
end
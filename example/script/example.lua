function setfunc(a)
	-- 支持userdata直接get/set
	a.number = 1993
	a.string = "tolua string"
end

function getfunc(a)
	return a.number, a.string
end

function showname(a)
	print(a:getName())
end

-- 调用堆栈
__TRACKBACK__ = function(msg)
    return debug.traceback(msg, 3)
end
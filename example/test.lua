function store_author(book, author)
	book.author = author
end

function load_author(book)
	return book.author
end

function print_name(book)
	print(book:Name())
end

function test_error()
	error("in test_error()")
end
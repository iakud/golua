function init_book(b)
	b:SetSellCallback(function(n)
		print(b:Name().." sell: "..n)
	end)
end

function store_author(book, author)
	book.author = author
end

function load_author(book)
	return book.author
end

function test_error()
	error("in test_error()")
end
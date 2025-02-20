module github.com/brend-n/monkey-interpreter/repl

replace github.com/brend-n/monkey-interpreter/lexer => ../lexer

replace github.com/brend-n/monkey-interpreter/token => ../token

go 1.23.1

require (
	github.com/brend-n/monkey-interpreter/lexer v0.0.0-00010101000000-000000000000
	github.com/brend-n/monkey-interpreter/token v0.0.0-00010101000000-000000000000
)

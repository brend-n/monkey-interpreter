package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brend-n/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	type expectedTokens struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	tests := []struct {
		input          string
		expectedTokens []expectedTokens
	}{
		{
			"=+(){},;",
			[]expectedTokens{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""}},
		},
		{
			`let five = 5;`,
			[]expectedTokens{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
			},
		},
		{
			`let five = 5;
			let ten = 10;

			let add = fn(x, y) {
				x + y;
			}

			let result = add(five, ten);
			`,
			[]expectedTokens{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)
		for _, expect := range tt.expectedTokens {
			tok := l.NextToken()
			assert.Equal(t, expect.expectedType, tok.Type)
			assert.Equal(t, expect.expectedLiteral, tok.Literal)
		}
	}
}

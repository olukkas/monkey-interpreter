package lexer

import (
	"github.com/stretchr/testify/assert"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x, y) {
    x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
    return true;
	} else {
    return false;
	}

	10 == 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];
	{"foo": "bar"}
	`

	testes := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assing, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assing, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assing, "="},
		{token.Function, "fn"},
		{token.Lparen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assing, "="},
		{token.Ident, "add"},
		{token.Lparen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.Rparen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.Gt, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.Lparen, "("},
		{token.Int, "5"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Else, "else"},
		{token.Lbrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Int, "10"},
		{token.Eq, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.NotEq, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		{token.String, "foobar"},
		{token.String, "foo bar"},
		{token.Lbracket, "["},
		{token.Int, "1"},
		{token.Comma, ","},
		{token.Int, "2"},
		{token.Rbracket, "]"},
		{token.Semicolon, ";"},
		{token.Lbrace, "{"},
		{token.String, "foo"},
		{token.Collon, ":"},
		{token.String, "bar"},
		{token.Rbrace, "}"},
		{token.Eof, ""},
	}

	asserts := assert.New(t)

	l := New(input)

	for _, tt := range testes {
		tok := l.NextToken()

		asserts.Equal(tok.Type, tt.expectedType, "token type wrong")
		asserts.Equal(tok.Literal, tt.expectedLiteral, "literal wrong")
	}
}

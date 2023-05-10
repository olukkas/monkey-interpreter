package token

type TokenType string

const (
	Illegal = "ILLEGAL"
	Eof     = "EOF"

	// identfiers + literal
	Ident = "IDENT"
	Int   = "INT"

	// Operators
	Assing   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"
	Lt       = "<"
	Gt       = ">"
	Eq       = "=="
	NotEq    = "!="

	// Delimiters
	Comma     = ","
	Semicolon = ";"
	Lparen    = "("
	Rparen    = ")"
	Lbrace    = "{"
	Rbrace    = "}"

	// Keywords
	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "False"
	If       = "IF"
	Else     = "Else"
	Return   = "Return"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return Ident
}

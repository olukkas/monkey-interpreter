package token

type TokenType string

const (
	Illegal = "ILLEGAL"
	Eof     = "EOF"

	// identfiers + literal
	Ident  = "IDENT"
	Int    = "INT"
	String = "STRING"

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
	Collon   = ":"

	// Delimiters
	Comma     = ","
	Semicolon = ";"
	Lparen    = "("
	Rparen    = ")"
	Lbrace    = "{"
	Rbrace    = "}"
	Lbracket  = "["
	Rbracket  = "]"

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

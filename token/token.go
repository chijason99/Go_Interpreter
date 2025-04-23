package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	// OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	// DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"return": RETURN,
}

func LookUpIdentifier(ident string) TokenType {
	if tokType, ok := keywords[ident]; ok {
		return tokType
	}

	return IDENT
}
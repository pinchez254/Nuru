package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

const (
	ILLEGAL = "HARAMU"
	EOF     = "MWISHO"

	// Identifiers + literals
	IDENT  = "KITAMBULISHI"
	INT    = "NAMBA"
	STRING = "NENO"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	POW      = "**"
	SLASH    = "/"
	MODULUS  = "%"
	LT       = "<"
	LTE      = "<="
	GT       = ">"
	GTE      = ">="
	EQ       = "=="
	NOT_EQ   = "!="
	AND      = "&&"
	OR       = "||"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "ACHA"
	TRUE     = "KWELI"
	FALSE    = "SIKWELI"
	IF       = "KAMA"
	ELSE     = "SIVYO"
	RETURN   = "RUDISHA"
	WHILE    = "WAKATI"
	NULL     = "TUPU"
	BREAK    = "SUSA"
	CONTINUE = "ENDELEA"
)

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"unda":    FUNCTION,
	"acha":    LET,
	"fanya":   LET,
	"kweli":   TRUE,
	"sikweli": FALSE,
	"kama":    IF,
	"au":      ELSE,
	"sivyo":   ELSE,
	"wakati":  WHILE,
	"rudisha": RETURN,
	"susa":    BREAK,
	"endelea": CONTINUE,
	"tupu":    NULL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

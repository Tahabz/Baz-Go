package token

type Token struct {
	Type string
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"
	IDENT	= "IDENT"
	INT		= "INT"
	ASSIGN	= "="
	PLUS	= "+"
	COMMA	= ","
	SEMICOLON	= ";"
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{" 
	RBRACE		= "}" 
	FUNCTION	= "FUNCTION"
	LET			= "LET"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT = "<"
	GT = ">"
	TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
)

var t = []string{
	"let",
	"fn",
	"true",
    "false",
    "if",
    "else",
    "return",
	}

func LookupIdent(ident string) string {
	for i := 0; i < len(t); i++ {
		if ident == t[i] {
			return getKeyType(ident)
		}
	}
	return IDENT
}

func getKeyType(ident string) string {
	var key string
	switch ident {
	case "let":
		key = LET
	case "fn":
		key = FUNCTION
	case "true":
		key = TRUE
	case "false":
		key = FALSE
	case "if":
		key = IF
	case "else":
		key = ELSE
	case "return":
		key = RETURN
	}
	return key
}
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
)

var t = []string{"let", "fn"}

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
	}
	return key
}
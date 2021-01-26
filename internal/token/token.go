package token

type Token struct {
	Type  string
	Value string
}

const (
	NUMBER   = "number"
	OPERATOR = "operator"

	PLUS     = "plus"
	MINUS    = "minus"
	DIVIDE   = "divide"
	MULTIPLY = "multiply"

	TERMINATION_SIGN = "?"
	QUESTION         = "what is"
)

var Keywords = map[string]string{
	"plus":          PLUS,
	"minus":         MINUS,
	"divided by":    DIVIDE,
	"multiplied by": MULTIPLY,
}

func NewNumericToken(v string) Token {
	t := Token{
		Type:  NUMBER,
		Value: v,
	}

	return t
}

func NewOperatorToken(v string) Token {
	t := Token{
		Type:  OPERATOR,
		Value: v,
	}

	return t
}

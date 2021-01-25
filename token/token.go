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

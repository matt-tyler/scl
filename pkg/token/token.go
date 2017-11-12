package token

type Token struct {
	Typ TokenType
	Val string
}

type TokenType int

const (
	String = iota

	keywordStart
	On
	Between

	opStart
	Or
	And
	opEnd

	LeftParen
	RightParen
	Comma
	keywordEnd
)

func IsOp(typ TokenType) bool {
	return opStart < typ && typ < opEnd
}

func (t Token) String() string {
	switch t.Typ {
	case String:
		return t.Val
	case Comma:
		return ","
	case On:
		return "ON"
	case Between:
		return "BETWEEN"
	case And:
		return "AND"
	case Or:
		return "OR"
	case LeftParen:
		return "("
	case RightParen:
		return ")"
	}

	return "UNKNOWN"
}

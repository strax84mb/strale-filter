package filter

import "fmt"

type ErrorType byte

const (
	NoError ErrorType = iota
	ErrUnexpectedToken
	ErrMalformedName
	ErrUnknownOperator
	ErrNotExhausted
)

type ErrParse struct {
	text    string
	errType ErrorType
}

func newParseError(errType ErrorType, runes []rune, index int) *ErrParse {
	if index >= len(runes) {
		return &ErrParse{
			text:    "failed to show error",
			errType: errType,
		}
	}
	return &ErrParse{
		text:    string(runes[index:]),
		errType: errType,
	}
}

func (ep *ErrParse) Error() string {
	var errText string
	switch ep.errType {
	case ErrUnexpectedToken:
		errText = "unexpected token encountered"
	case ErrMalformedName:
		errText = "character not allowed in name"
	case ErrUnknownOperator:
		errText = "unknown operator"
	case ErrNotExhausted:
		errText = "not all characted processed"
	default:
		errText = "no error"
	}
	return fmt.Sprintf("%s, unprocessed text: %s", errText, ep.text)
}

package filter

import "fmt"

type ErrorType string

const (
	NoError            ErrorType = "no error"
	ErrUnexpectedToken ErrorType = "unexpected token encountered"
	ErrMalformedName   ErrorType = "character not allowed in name"
	ErrUnknownOperator ErrorType = "unknown operator"
	ErrNotExhausted    ErrorType = "not all characted processed"
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
	return fmt.Sprintf("%s, unprocessed text: %s", string(ep.errType), ep.text)
}

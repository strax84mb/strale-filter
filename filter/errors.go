package filter

type ErrorType string

const (
	ErrUnexpectedToken ErrorType = "unexpected token encountered"
	ErrMalformedName   ErrorType = "character not allowed in name"
	ErrUnknownOperator ErrorType = "unknown operator"
	ErrNotExhausted    ErrorType = "not all characted processed"
)

func (et ErrorType) Error() string {
	return string(et)
}

type ErrParse struct {
	cause string
}

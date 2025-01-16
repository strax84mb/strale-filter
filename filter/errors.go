package filter

type ErrorType string

const (
	ErrUnexpectedToken ErrorType = "Unexpected token encountered"
	ErrMalformedName   ErrorType = "Character not allowed in name"
	ErrUnknownOperator ErrorType = "Unknown operator"
)

type ErrParse struct {
	cause string
}

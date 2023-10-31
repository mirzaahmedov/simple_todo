package errors

type ValidationErrors struct {
	code   int
	errors []ValidationError
}
type ValidationError struct {
	field   string
	kind    string
	message string
}

func (e *ValidationError) Error() string {
	return e.message
}

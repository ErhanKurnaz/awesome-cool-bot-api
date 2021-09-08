package errors

type NoBodyError struct {
}

func (_ *NoBodyError) Error() string {
	return "No body provided"
}

func NewNoBodyError() error {
	return &NoBodyError{}
}

package api_errors

type ApiError interface {
	Code() int
	Error() string
}

type BaseApiError struct {
	message string
	code    int
}

func (err *BaseApiError) Code() int {
	return err.code
}

func (err *BaseApiError) Error() string {
	return err.message
}

func New(code int, message string) ApiError {
	return &BaseApiError{
		message: message,
		code:    code,
	}
}

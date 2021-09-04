package models

type Response struct {
	Status string `json:"status"`
}

type ResponseSuccess struct {
	Response
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Response
	Error string `json:"error"`
}

func NewSuccessResponse(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{
		Response{"success"},
		data,
	}
}

func NewErrorResponse(error error) *ResponseError {
	return &ResponseError{
		Response{"error"},
		error.Error(),
	}
}

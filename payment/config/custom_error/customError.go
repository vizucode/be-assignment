package customerror

type CustomError struct {
	Message  string `json:"message"`
	HttpCode int    `json:"http_code"`
}

func NewCustomError(message string, httpCode int) *CustomError {
	return &CustomError{
		Message:  message,
		HttpCode: httpCode,
	}
}

func (err *CustomError) Error() string {
	return err.Message
}

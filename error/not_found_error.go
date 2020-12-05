package myerror

type NotFoundError struct {
	code          ErrorCode
	originalError error
	message       string
}

func NewNotFoundError(originalError error, message string) NotFoundError {
	return NotFoundError{
		code:          NotFound,
		originalError: originalError,
		message:       message,
	}
}

func (err NotFoundError) Error() string {
	return err.originalError.Error()
}

func (err NotFoundError) Code() ErrorCode {
	return err.code
}

func (err NotFoundError) Messages() []string {
	return []string{err.message}
}

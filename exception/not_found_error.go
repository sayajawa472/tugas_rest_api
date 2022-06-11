package exception

type NotFoundError struct {
	error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{error: error}

}

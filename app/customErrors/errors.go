package customErrors

type TypedError struct {
	Msg string
}

func (e TypedError) Error() string {
	return e.Msg
}

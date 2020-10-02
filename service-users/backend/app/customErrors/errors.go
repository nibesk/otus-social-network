package customErrors

type TypedError struct {
	Msg string
}

func (e TypedError) Error() string {
	return e.Msg
}

type TypedStatusError struct {
	Msg    string
	Status int
}

func (e TypedStatusError) Error() string {
	return e.Msg
}

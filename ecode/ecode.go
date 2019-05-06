package ecode

type ecodeError struct {
	error

	code    Code
	message string
	details []interface{}
}

func (e *ecodeError) Code() Code {
	if e == nil {
		return 0
	}
	return e.code
}

// upstream
type paramsError struct {
	error
}

type downstreamError struct {
	error
}

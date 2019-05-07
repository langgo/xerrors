package errors

type withMessage struct {
	err error
	s   string
}

func (w *withMessage) Error() string {
	panic("not implemented")
	// error link
}

func (w *withMessage) Unwrap() error {
	return w.err
}

type withStack struct {
	err error
	//
}

func (w *withStack) Error() string {
	panic("not implemented")
	// error link
}

func (w *withStack) Unwrap() error {
	panic("not implemented")
}

func a(err error) {

}

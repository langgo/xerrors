package main

import (
	"fmt"
	"runtime"
	"strings"

	_ "github.com/pkg/errors"
	_ "golang.org/x/xerrors"
)

func main() {
	a()
}

func a() {
	b()
}

func b() {
	c()
}

func c() {
	var frames [16]uintptr
	n := runtime.Callers(1, frames[:])

	st := frames[0:n]

	f := runtime.CallersFrames(st)

	for {
		fr, more := f.Next()

		fmt.Printf("%s\n\t%s:%d\n", fr.Function, fr.File, fr.Line)

		if !more {
			break
		}
	}

	fmt.Println("===")
}

type Wrapper interface {
	Unwrap() error
}

type Holder interface {
	GetFns() []func(err error)
	Format(p Printer) // 只是把公共后移到了Format中去做，并没有做到更好
}

type Hold struct {
	msgs   []string
	codes  []int
	stacks []string

	Fns []func(err error)
}

func NewHold() Holder {
	h := &Hold{}
	h.Fns = append(h.Fns, h.msg, h.code, h.stack)
	return h
}

func (h *Hold) GetFns() []func(err error) {
	return h.Fns
}

func (h *Hold) Format(p Printer) {
	if len(h.codes) > 0 {
		p.Printf("%d\n", h.codes[0])
	}
	p.Printf("%s\n", strings.Join(h.msgs, ":"))

	if len(h.stacks) > 0 {
		p.Printf("%s\n", h.stacks[0])
	}
}

func Wrap(err error, msg string) error {
	panic("not implemented")
}

func Wrapf(err error, fmt string, args ...interface{}) error {
	panic("not implemented")
}

func WrapCode(err error, code int, msg string) error {
	panic("not implemented")
}

func WrapCodef(err error, code int, fmt string, args ...interface{}) {
	panic("not implemented")
}

//

func WithCode(err error, code int) error {
	panic("not implemented")
}

func WithStack(err error) error {
	panic("not implemented")
}

func WithMessage(err error, msg string) error {
	panic("not implemented")
}

func (h *Hold) msg(err error) {
	if v, ok := err.(interface{ Message() string }); ok {
		h.msgs = append(h.msgs, v.Message())
	}
}

func (h *Hold) code(err error) {
	if v, ok := err.(interface{ Code() int }); ok {
		h.codes = append(h.codes, v.Code())
	}
}

func (h *Hold) stack(err error) {
	if v, ok := err.(interface{ Stack() string }); ok {
		h.stacks = append(h.stacks, v.Stack())
	}
}

func Error(hold func() Holder, err error, p Printer) {
	h := hold()
	for err != nil {
		next, ok := err.(Wrapper)
		if !ok {
			break
		}

		err = next.Unwrap()
		for _, fn := range h.GetFns() {
			fn(err)
		}
	}

	h.Format(p)
}

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

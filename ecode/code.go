package ecode

import (
	"context"
)

type Code int

func (c Code) Value() int {
	return int(c)
}

var (
	ErrParams Code = 400000 // 二级错误码
)

// 前两位表示错误类型
// 后四位表示错误类型下的错误码
// 错误类型+0000 表示默认错误
// 0 表示成功

var (
	Success Code = 0
)

var (
	ErrInternal Code = 500000
)

var (
	ErrDownstream Code = 520000
)

type HandleFunc func(ctx context.Context, req interface{}) (resp interface{}, err error)

type Middleware func(next HandleFunc) HandleFunc

// Code == 0 表示成功
type Resp struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	NodeId  string        `json:"node_id"` // or cluster+ip 在分布式的情况下，比较容易处理错误
	TraceId string        `json:"trace_id"`
	Details []interface{} `json:"details,omitempty"`
	Data    interface{}   `json:"data,omitempty"`
}

func demo(next HandleFunc) HandleFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// ensure trace_id

		resp, err := next(ctx, req)
		if err == nil {
			return &Resp{
				Data: resp,
			}, nil
		}

		var e = &Resp{}
		if _, ok := is("e is codeError"); ok {
			e.Code = 0
			e.Message = ""
		} else if _, ok := is("e is paramsError"); ok {
			e.Code = ErrParams.Value()
			e.Message = ""
			// 如果通过错误码区别，具体的错误。多个接口的错误码很难保证统一。在一个地方统一定义错误码，这样比较容易保证唯一，例如通过code显示文案
			//
			// 验证码错误
			// 身份证号与实名不符
			//
			// 上面这种，错误是要依据实际情况判断的，如果客户端不需要进行处理（或者想，格式不合法，这种前端可以直接禁止提交），那么一个通用的文案也是可以的。
		} else {
			if _, ok := is("e is downstream"); ok {
				e.Code = ErrDownstream.Value()
				e.Message = ""
			} else {
				e.Code = ErrInternal.Value()
				e.Message = "internal error" // 话术
			}
		}

		e.TraceId = ""  // parse from ctx
		e.Details = nil // if dev then marshal err.Unwrap()...

		return e, nil
	}
}

// just for doc
func is(v string) (interface{}, bool) {
	return v, v != ""
}

type fundamental struct {
	msg string
	// stack string // for doc
}

type withStack struct {
	error
	stack string // for doc
}

type withMessage struct {
	error
	message string
}

type withDetails struct {
	error
	// message string
	details []interface{}
}

// 通过接口判断

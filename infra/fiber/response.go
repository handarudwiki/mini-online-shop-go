package infrafiber

import (
	response "handarudwiki/mini-online-shop-go"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	HttpCode int         `json:"-"`
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Payload  interface{} `json:"payload,omitempty"`
	Query    interface{} `json:"query,omitempty"`
	Error    string      `json:"error,omitempty"`
	ErrCode  string      `json:"error_code,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var res = Response{
		Success: true,
	}
	for _, param := range params {
		param(&res)
	}
	return res
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithMessage(message string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Message = message
		return r
	}
}

func WithPayload(payload interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}

func WithQuery(query interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Query = query
		return r
	}
}

func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false
		myError, ok := err.(response.Error)
		if !ok {
			myError = response.ErrorGeneral
		}
		r.Error = myError.Code
		r.ErrCode = myError.Code
		r.HttpCode = myError.HttpCode

		return r
	}
}

func (r Response) Send(ctx *fiber.Ctx) error {
	return ctx.Status(r.HttpCode).JSON(r)
}

package infrafiber

import "github.com/aripkur/go-learn-shop/infra/response"

type Response struct {
	HttpCode  int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var resp = Response{
		Success: true,
	}

	for _, param := range params {
		param(&resp)
	}

	return resp
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithHttpMessage(message string) func(*Response) *Response {
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

func WithErrorMessage(errorMessage string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Error = errorMessage
		return r
	}
}
func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		myErr, ok := err.(response.Error)

		if !ok {
			myErr = response.ErrorGeneral
		}

		r.Error = myErr.Message
		r.ErrorCode = myErr.Code
		return r
	}
}

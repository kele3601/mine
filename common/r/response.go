package r

import "github.com/gin-gonic/gin"

type Response struct {
	Flag bool   `json:"flag"`
	Code int    `json:"code"`
	Mes  string `json:"mes"`
	Data any    `json:"data"`
}

func Return(ctx *gin.Context, r *Response) {
	ctx.JSON(r.Code, r)
}

func Fail() *Response {
	return &Response{Flag: false, Code: 999, Mes: "fail"}
}

func OK() *Response {
	return &Response{Flag: true, Code: 200, Mes: "ok"}
}

func (r *Response) SetData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) SetMes(mes string) *Response {
	r.Mes = mes
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

package bean

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ctx *gin.Context
}

type ResultData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultData{}
	res.Code = "200"
	res.Msg = "处理成功"
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *Result) Error(code string, msg string) {
	res := ResultData{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
	r.Ctx.Abort()
}

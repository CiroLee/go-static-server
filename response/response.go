package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int
	Msg  string
	Data any
}

var SavedError = Res{
	Code: -1002,
	Msg:  "fail to save",
}

var UnAuthorization = Res{
	Code: -3000,
	Msg:  "UnAuthorization",
}

var EmptyList = Res{
	Code: -1003,
	Msg:  "no data",
	Data: []string{},
}

func Success(ctx *gin.Context, data any, status int) {
	var s = http.StatusOK
	if status != 0 {
		s = status
	}
	ctx.JSON(s, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func Fail(ctx *gin.Context, res Res, status int) {
	var s = http.StatusOK
	if status != 0 {
		s = status
	}
	ctx.JSON(s, gin.H{
		"code": res.Code,
		"msg":  res.Msg,
		"data": res.Data,
	})
}

func WrongFormat(ctx *gin.Context, data string) {
	Fail(ctx, Res{
		Code: -1001,
		Msg:  "unaccepted file format",
		Data: data,
	}, 0)
}

func UnAuth(ctx *gin.Context) {
	Fail(ctx, UnAuthorization, http.StatusUnauthorized)
}

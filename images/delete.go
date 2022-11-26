package images

import (
	"github.com/CiroLee/go-static-server/response"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type deleteProps struct {
	Name string `form:"name" binding:"required"`
}

func ImageDeleteHandler(ctx *gin.Context) {
	var form deleteProps
	if err := ctx.ShouldBind(&form); err != nil {
		response.WrongParams(ctx, err)
		return
	}

	removeErr := os.Remove(path.Join(".", BasePath, form.Name))
	if removeErr != nil {
		response.Fail(ctx, response.Res{
			Code: response.DeleteFail.Code,
			Msg:  response.DeleteFail.Msg,
			Data: "no such file",
		}, 0)
		return
	}

	response.Success(ctx, nil, 0)

}

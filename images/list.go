package images

import (
	"github.com/CiroLee/go-static-server/utils"
	"log"
	"os"
	"path"

	"github.com/CiroLee/go-static-server/response"
	"github.com/gin-gonic/gin"
)

type List struct {
	Folder string `form:"folder" binding:"omitempty"`        // 目标目录
	Type   string `form:"type" binding:"oneof=folder image"` // 查询类型：目录 or  图片
}

func noDirectory(ctx *gin.Context, err error) {
	response.Fail(ctx, response.EmptyList, 0)
	log.Println(err)
	ctx.Abort()
}

func ImageListHandler(ctx *gin.Context) {
	var body List
	if err := ctx.ShouldBind(&body); err != nil {
		response.WrongParams(ctx, err)
		return
	}
	file, err := os.Open(path.Join(".", BasePath, body.Folder))

	if err != nil {
		noDirectory(ctx, err)
		return
	}
	files, readErr := file.ReadDir(-1)
	closeErr := file.Close()
	if closeErr != nil {
		noDirectory(ctx, err)
		return
	}
	if readErr != nil {
		noDirectory(ctx, err)
		return
	}
	var list = make([]string, 0)
	for _, f := range files {
		if body.Type == "image" {
			if !f.IsDir() {
				list = append(list, utils.GetUrlByEnv(path.Join(BasePath, body.Folder), f.Name()))
			}
		} else {
			if f.IsDir() {
				list = append(list, f.Name())
			}
		}

	}

	response.Success(ctx, list, 0)
}

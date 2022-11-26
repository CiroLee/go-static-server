package images

import (
	"log"
	"os"
	"path"

	"github.com/CiroLee/go-static-server/response"
	"github.com/CiroLee/go-static-server/utils"
	"github.com/gin-gonic/gin"
)

func noDirectory(ctx *gin.Context, err error) {
	response.Fail(ctx, response.EmptyList, 0)
	log.Println(err)
	ctx.Abort()
}

func ImageListHandler(ctx *gin.Context) {
	file, err := os.Open(path.Join(".", BasePath))
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
		list = append(list, utils.GetUrlByEnv(BasePath, f.Name()))
	}

	response.Success(ctx, list, 0)
}

package images

import (
	"fmt"
	"github.com/CiroLee/go-static-server/utils"
	"github.com/jaevor/go-nanoid"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/CiroLee/go-static-server/response"
	"github.com/gin-gonic/gin"
)

func ImageUploadHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	header := file.Header
	fileType := header["Content-Type"][0]
	fileExt := strings.ToLower(path.Ext(file.Filename))
	isImage, _ := regexp.Match("^image", []byte(fileType))
	// 不是图片格式，返回错误提示
	if !isImage {
		response.WrongFormat(ctx, fmt.Sprintf("file type is %v", fileType))
		return
	}
	// 检查存储目录是否存在，不存在则创建
	savedPath := path.Join(".", BasePath)
	if ok, _ := utils.PathExists(savedPath); !ok {
		err := os.MkdirAll(savedPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	// 创建唯一文件名
	createNanoId, _ := nanoid.Standard(21)
	// timestamp + nanoid + 图片后缀
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + createNanoId() + fmt.Sprintf("%v", fileExt)
	dst := path.Join(savedPath, filename)
	// 保存图片
	saveErr := ctx.SaveUploadedFile(file, dst)
	if saveErr != nil {
		response.Fail(ctx, response.SavedError, 0)
		return
	}

	response.Success(ctx, map[string]string{
		"url": utils.GetUrlByEnv(BasePath, filename),
	}, 0)
}

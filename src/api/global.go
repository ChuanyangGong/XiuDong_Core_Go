package api

import (
	"XDCore/src/global/response"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	root      string   = "storage"
	validPath []string = []string{"others", "cover"}
)

// 上传文件
func UploadFiles(ctx *gin.Context) {
	folder := GetValiedFolder(ctx.Param("folder"))
	wd, _ := os.Getwd()
	folderPath := path.Join(wd, root, folder)

	form, err := ctx.MultipartForm()
	if err != nil {
		zap.S().Errorf("上传文件请求出错：%v", err)
		ctx.JSON(http.StatusBadRequest, response.BaseRsp{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	f := form.File["file"]
	filePath := []string{}
	for _, file := range f {
		filename := file.Filename
		zap.S().Debugf("正在上传文件 %s", filename)
		ctx.SaveUploadedFile(file, path.Join(folderPath, filename))
		filePath = append(filePath, path.Join(folder, filename))
	}

	ctx.JSON(http.StatusOK, response.BaseRsp{
		Success: true,
		Data:    filePath,
	})
}

// 初始化目录
func GetValiedFolder(input string) string {
	output := validPath[0]
	for _, v := range validPath {
		if v == input {
			output = input
			break
		}
	}
	return output
}

// 创建文件夹
func TestAndCreateFolder(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			zap.S().Errorf("创建文件夹 %s 失败: %v", path, err)
		}
	} else {
		zap.S().Errorf("创建文件夹 %s 失败: %v", path, err)
	}
}

func init() {
	wd, _ := os.Getwd()
	rootPath := path.Join(wd, root)
	// 创建根目录文件夹
	TestAndCreateFolder(rootPath)

	// 创建文件目录
	for _, fold := range validPath {
		folder := path.Join(rootPath, fold)
		TestAndCreateFolder(folder)
	}
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			fmt.Printf("upload file failed ,err:%v\n", err)
		} else {
			fmt.Printf("file name :%v, size:%d, header:%v\n", file.Filename, file.Size, file.Header)
			// 将读取到的文件保存到本地(服务器本地)
			// dst := fmt.Sprintf("./%s", file.Filename)  // 路径
			dst := path.Join("./", file.Filename) // 路径
			fmt.Println(dst)
			context.SaveUploadedFile(file, dst)
		}
	})

	r.Run()
}

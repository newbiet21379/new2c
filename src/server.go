package main

import (
	"aweSomeGin/controller"
	"aweSomeGin/middleware"
	"aweSomeGin/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"os"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutPut () {
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f , os.Stdout)
}

func main()  {
	server := gin.New()

	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
		gindump.Dump())

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	server.Run(":8080")
}

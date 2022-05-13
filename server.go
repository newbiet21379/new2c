package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/newbiet21379/new2c/controller"
	"github.com/newbiet21379/new2c/middleware"
	"github.com/newbiet21379/new2c/service"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
	port                                       = goDotEnvVariable("server.port")
)

func setupLogOutPut() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
		gindump.Dump())
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(context *gin.Context) {
			err := videoController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})
		apiRoutes.DELETE("/videos/:id", func(context *gin.Context) {
			err := videoController.DeleteOne(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Delete Successfully!!"})
			}
		})
		apiRoutes.PUT("/video/url", func(context *gin.Context) {
			err := videoController.UpdateUrl(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Update URL Successfully!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(port)
}

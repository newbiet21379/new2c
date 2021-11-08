package controller

import (
	"aweSomeGin/entity"
	"aweSomeGin/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func (controller *controller) FindAll() []entity.Video {
	return  controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return entity.Video{}
	}
	controller.service.Save(video)
	return video
}

func New(service service.VideoService) VideoController{
	return &controller{
		service: service,
	}
}

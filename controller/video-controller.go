package controller

import (
	"aweSomeGin/entity"
	"aweSomeGin/service"
	"aweSomeGin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController{
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll() []entity.Video {
	return  controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}

func (controller *controller) ShowAll(ctx *gin.Context){
	videos := controller.service.FindAll()
	data := gin.H{
		"title" : "Video Page",
		"videos" : videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}


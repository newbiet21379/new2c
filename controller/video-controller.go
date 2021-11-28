package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/newbiet21379/new2c/entity"
	"github.com/newbiet21379/new2c/service"
	"github.com/newbiet21379/new2c/validators"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	UpdateUrl(ctx *gin.Context) error
	DeleteOne(ctx *gin.Context) error
}

type URLRequest struct {
	id string `json:"id"`
	url string `json:"url"`
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
	var (
		video entity.Video
		err error
	)
	err = ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	video,err = controller.service.Save(video)
	if err != nil {
		return err
	}
	return nil
}

func (controller *controller) DeleteOne(ctx *gin.Context) error{
	var id string
	id = ctx.Param("id")
	err := controller.service.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

func (controller *controller) UpdateUrl(ctx *gin.Context) error{
	var urlRequest URLRequest
	err := ctx.ShouldBindJSON(&urlRequest)
	if err != nil {
		return err
	}
	err = controller.service.UpdateUrl(urlRequest.id,urlRequest.url)
	if err != nil {
		return err
	}
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


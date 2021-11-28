package service

import (
	"github.com/newbiet21379/new2c/dao"
	"github.com/newbiet21379/new2c/entity"
)

type VideoService interface {
	Save(video entity.Video) (entity.Video,error)
	FindAll() []entity.Video
	DeleteOne(id string) error
	UpdateUrl(id string,url string) error
}

type videoService struct {
	videos []entity.Video
}

func (service *videoService) Save(video entity.Video) (entity.Video,error) {
	var result entity.Video
	err := dao.CreateVideo(video)
	if err != nil{
		return result,err
	}
	return video,nil
}

func (service *videoService) FindAll() []entity.Video {
	list ,err := dao.GetAllVideos()
	if err != nil {
		return []entity.Video{}
	}
	return list
}

func (service *videoService) DeleteOne(id string) error {
	err := dao.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *videoService) UpdateUrl(id string, url string) error{
	err := dao.UpdateURL(id,url)
	if err != nil {
		return err
	}
	return nil
}

func New() VideoService{
	return &videoService{}
}


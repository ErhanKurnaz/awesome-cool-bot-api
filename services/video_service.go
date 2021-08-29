package services

import (
	"github.com/ErhanKurnaz/awesome-cool-bot-api/entities"
	"github.com/ErhanKurnaz/awesome-cool-bot-api/repositories"
)

type VideoService interface {
	Save(entities.Video) entities.Video
	Update(video entities.Video) entities.Video
	Delete(video entities.Video)
	FindAll() []entities.Video
}

type videoService struct {
	repository repositories.VideoRepository
}

func (service *videoService) Save(video entities.Video) entities.Video {
	service.repository.Save(video)
	return video
}

func (service *videoService) FindAll() []entities.Video {
	return *service.repository.FindAll()
}

func (service *videoService) Update(video entities.Video) entities.Video {
	service.repository.Update(video)

	return video
}

func (service *videoService) Delete(video entities.Video) {
	service.repository.Delete(video)
}

func NewVideoService(repository repositories.VideoRepository) VideoService {
	return &videoService{
		repository: repository,
	}
}

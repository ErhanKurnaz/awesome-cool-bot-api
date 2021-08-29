package repositories

import (
	"github.com/ErhanKurnaz/awesome-cool-bot-api/entities"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entities.Video)
	Update(video entities.Video)
	Delete(video entities.Video)
	FindAll() *[]entities.Video
}

type videoRepository struct {
	connection *gorm.DB
}


func (repository *videoRepository) Save(video entities.Video) {
	repository.connection.Create(&video)
}

func (repository *videoRepository) Update(video entities.Video) {
	repository.connection.Save(&video)
}

func (repository *videoRepository) Delete(video entities.Video) {
	repository.connection.Delete(&video)
}

func (repository *videoRepository) FindAll() *[]entities.Video {
	var videos []entities.Video
	repository.connection.Set("gorm:auto_preload", true).Find(&videos)
	return &videos
}

func NewVideoRepository(db *gorm.DB) VideoRepository {
	err := db.AutoMigrate(&entities.Video{}, &entities.Person{})

	if err != nil {
		panic("Failed to migrate models: Video and Person, with error: " + err.Error())
	}

	return &videoRepository{
		connection: db,
	}
}

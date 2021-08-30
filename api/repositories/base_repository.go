package repositories

// We can only uncomment this code once generics are actually part of the golang language
// this should happen at the end of December 2021

/**
import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[T interface{}] interface {
	Save(entity T)
	Update(entity T)
	Delete(entity T)
	FindAll() *[]T
}


type DefaultRepository[T interface{}] struct {
	db *gorm.DB
}


func (repository *DefaultRepository[T]) Save(entity T) {
	repository.db.Create(&entity)
}

func (repository *DefaultRepository[T]) Update(entity T) {
	repository.db.Save(&entity)
}

func (repository *DefaultRepository[T]) Delete(entity T) {
	repository.db.Delete(&entity)
}

func (repository *DefaultRepository[T]) FindAll() *[]T {
	var entities []T
	repository.db.Preload(clause.Associations).Find(&entities)
	return &entities
}
**/

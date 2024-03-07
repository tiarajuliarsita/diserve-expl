package repository

import (
	"diserve-expl/cache"
	"log"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type RepoInterface interface {
	CreatePost(cache.Post) error
	FindByID(id string) (cache.Post, error)
	FindByName(name string) (cache.Post, error)
}

func NewRepo(db *gorm.DB) RepoInterface {
	return &repo{
		db,
	}
}

// CreatePost implements RepoInterface.
func (r *repo) CreatePost(v cache.Post) error {
	err := r.db.Create(&v).Error
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements RepoInterface.
func (r *repo) FindByID(id string) (cache.Post, error) {
	result := cache.Post{}

	err := r.db.Where("id = ?", id).Find(&result).Error
	if err != nil {
		return cache.Post{}, err
	}
	log.Println("result", result)
	return result, nil
}

// FindByName implements RepoInterface.
func (r *repo) FindByName(name string) (cache.Post, error) {
	result := cache.Post{}

	err := r.db.Where("name = ?", name).Find(&result).Error
	if err != nil {
		return cache.Post{}, err
	}
	log.Println("result", result)
	return result, nil
}

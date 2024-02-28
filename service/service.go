package service

import (
	"diserve-expl/cache"
	"diserve-expl/repository"
	"errors"

	"time"
)

type (
	svc struct {
		rp repository.RepoInterface
		// rd cache.PostChace
	}

	SvcInterface interface {
		CreatePost(cache.Post) error
		FindByID(id string) (cache.Post, error)
	}
)

func NewSvc(rp repository.RepoInterface) SvcInterface {
	return &svc{
		rp,
	}
}

func (s *svc) FindByID(id string) (cache.Post, error) {

	v, err := s.rp.FindByID(id)

	time.Sleep(11)
	if err != nil {
		return v, errors.New("error disini")
	}

	return v, nil

}

func (s *svc) CreatePost(v cache.Post) error {
	err := s.rp.CreatePost(v)
	if err != nil {
		return err
	}
	return nil
}

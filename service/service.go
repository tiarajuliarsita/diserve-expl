package service

import (
	"diserve-expl/auth"
	"diserve-expl/cache"
	"diserve-expl/repository"
	"errors"
	"os"

	"time"

	"github.com/joho/godotenv"
)

type (
	svc struct {
		rp repository.RepoInterface
		// rd cache.PostChace

	}

	SvcInterface interface {
		CreatePost(cache.Post) error
		FindByID(id string) (cache.Post, error)
		FindByName(name string) (cache.Post, auth.Token, error)
		RefreshToken(refreshToken string) (auth.Token, error)
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

// FindByName implements SvcInterface.
func (s *svc) FindByName(name string) (cache.Post, auth.Token, error) {
	t := auth.Token{}
	v, err := s.rp.FindByName(name)
	if err != nil {
		return cache.Post{}, t, err
	}

	t.AccessToken = auth.AccessToken(v.ID)
	t.RefreshToken = auth.RefreshToken(v.ID)

	return v, t, nil
}

func (s *svc) RefreshToken(refreshToken string) (auth.Token, error) {
	godotenv.Load()
	t := auth.Token{}
	id, err := auth.Parse(refreshToken, os.Getenv("ACCESS_TOKEN_SECRET"))
	if err != nil {
		return auth.Token{}, err
	}
	idS := id.(string)
	t.AccessToken = auth.AccessToken(idS)
	t.RefreshToken = auth.RefreshToken(idS)

	return t, nil

}

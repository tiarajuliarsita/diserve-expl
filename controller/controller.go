package controller

import (
	"diserve-expl/auth"
	"diserve-expl/cache"
	"diserve-expl/service"
	"log"

	"github.com/labstack/echo/v4"
)

type ctrl struct {
	svc service.SvcInterface
	rd  cache.PostChace
}

func NewController(s service.SvcInterface, rd cache.PostChace) *ctrl {
	return &ctrl{
		svc: s,
		rd:  rd,
	}
}

func (c *ctrl) CreatePost(e echo.Context) error {
	req := cache.Post{}
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, err)
	}
	err = c.svc.CreatePost(req)
	if err != nil {
		return e.JSON(400, err)
	}
	return e.JSON(201, "success")
}
func (c *ctrl) FindByID(e echo.Context) error {
	id := e.Param("id")
	log.Println("id : ", id)
	post, _ := c.rd.Get(id)
	// log.Println(" post : ", post)
	if post == nil {
		v, err := c.svc.FindByID(id)
		if err != nil {
			return e.JSON(400, err)
		}
		c.rd.Set(id, v)
		return e.JSON(200, v)

	} else {
		return e.JSON(200, post)
	}
}

func (c *ctrl) Login(e echo.Context) error {

	req := cache.Post{}
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(400, err)
	}

	v, at, err := c.svc.FindByName(req.Name)
	if err != nil {
		return e.JSON(500, err)
	}
	return e.JSON(200, echo.Map{
		"data":  v,
		"token": at,
	})

}

func (c *ctrl) RefreshToken(e echo.Context) error {
	ref := auth.ReqRefToken{}
	err := e.Bind(&ref)
	if err != nil {
		return e.JSON(400, err)
	}

	t, err := c.svc.RefreshToken(ref.RefreshToken)
	if err != nil {
		return e.JSON(400, err)
	}
	return e.JSON(200, t)

}

package handler

import (
	"github.com/EgorMizerov/testProject/repository"
	"github.com/fasthttp/router"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Router() *router.Router {
	r := router.New()

	r.GET("/json/hackers", h.GetSortedHackers)

	return r
}

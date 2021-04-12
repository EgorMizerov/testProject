package repository

import (
	"github.com/EgorMizerov/testProject/domain"
	"github.com/go-redis/redis/v8"
)

type Hacker interface {
	TestData() error
	GetHackers() ([]domain.Hacker, error)
}

type Repository struct {
	Hacker
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{
		Hacker: NewHackerRedis(rdb),
	}
}

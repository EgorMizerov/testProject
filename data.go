package main

import (
	"github.com/EgorMizerov/testProject/config"
	"github.com/EgorMizerov/testProject/repository"
)

func main() {
	redisCfg := repository.Config{
		Host:     config.Redis["host"],
		Port:     config.Redis["port"],
		Password: config.Redis["password"],
	}

	rdb := repository.ConnectRedis(redisCfg)
	repo := repository.NewRepository(rdb)

	err := repo.Hacker.TestData()
	if err != nil {
		return
	}
}

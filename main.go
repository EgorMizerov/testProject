package main

import (
	"fmt"
	"github.com/EgorMizerov/testProject/config"
	"github.com/EgorMizerov/testProject/handler"
	"github.com/EgorMizerov/testProject/repository"
	"github.com/valyala/fasthttp"
)

var count = 0

func main() {
	Run()
}

func Run() {
	forever := make(chan int)

	// redis config
	redisCfg := repository.Config{
		Host:     config.Redis["host"],
		Port:     config.Redis["port"],
		Password: config.Redis["password"],
	}

	rdb := repository.ConnectRedis(redisCfg)
	repo := repository.NewRepository(rdb)
	h := handler.NewHandler(repo)

	router := h.Router().Handler
	addr := fmt.Sprintf("%s:%s", config.Server["host"], config.Server["port"])

	go func() {
		fmt.Println("Starting listen to server...")
		fmt.Printf("http://%s\n", addr)
		err := fasthttp.ListenAndServe(addr, router)
		if err != nil {
			return
		}
	}()

	<-forever
}

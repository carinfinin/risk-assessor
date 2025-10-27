package main

import (
	"fmt"
	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/carinfinin/risk-assessor/internal/logger"
	"github.com/carinfinin/risk-assessor/internal/server"
	"github.com/carinfinin/risk-assessor/internal/service"
)

func main() {
	fmt.Println("start")

	cfg := config.New("")

	err := logger.Configure(cfg)
	if err != nil {
		fmt.Println(err)
	}
	log, err := logger.Get()
	if err != nil {
		fmt.Println(err)
	}

	s := service.New()
	router := server.NewRouter(cfg, s)

	svr := server.New(cfg, router)
	svr.Start()

	log.Info("start server")

	/*
		конфиг
		логер
		кеш
		воркеры
		сервер
	*/

}

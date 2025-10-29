package main

import (
	"fmt"
	"github.com/carinfinin/risk-assessor/internal/config"
	"github.com/carinfinin/risk-assessor/internal/encryption"
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

	keyProvider := encryption.NewInMemoryKeyProvider()
	//// Генерируем или загружаем ключи
	//keyV1 := make([]byte, 32) // AES-256
	//rand.Read(keyV1)
	//keyProvider.AddKey("key_v1_2024", keyV1)

	// Создаем encryptor
	encryptor := encryption.NewEncryptor(keyProvider)
	// Запускаем сервисы
	s := service.New(encryptor)
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

package main

import (
	"log"

	"github.com/aripkur/go-learn-shop/external/database"
	"github.com/aripkur/go-learn-shop/internal/config"
)

func main() {
	filename := "cmd/api/config.yaml"

	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}
}

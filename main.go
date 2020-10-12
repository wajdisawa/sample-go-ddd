package main

import (
	"log"
	"sample-go-ddd/config"
	infrastructure "sample-go-ddd/infrastructure/persistence"
	"sample-go-ddd/interfaces"
	"strconv"
)

func main() {
	db, err := config.NewDBConnection()
	if err != nil{
		log.Fatal(err)
		return
	}
	infra := infrastructure.NewUserRepository(db)
	handler := interfaces.NewHandler(infra)
	port, err := strconv.ParseInt(config.GetEnvWithDefault("APP_PORT", "9000"), 0, 0)
	if err != nil{
		log.Fatal(err)
		return
	}
	if err:= handler.Run(int(port)); err !=nil{
		log.Fatal(err)
	}
}

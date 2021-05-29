package main

import (
	"fmt"
	"log"

	"github.com/TiagoCiceri/go-desafios/application/repositories"
	"github.com/TiagoCiceri/go-desafios/domain"
	"github.com/TiagoCiceri/go-desafios/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Tiago",
		Email:    "tiago@gmail.com",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert((&user))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tacoProject/internal/client/postgresql"
	"tacoProject/pkg/routes"
)

func main() {
	postgresql.LoadEnvVariables()

	db, err := postgresql.OpenDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.Initialize(router, db)

	router.Run()

}

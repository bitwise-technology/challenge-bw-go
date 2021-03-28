package main

import (
	"log"

	"github.com/challenge-bw-go/src/database"
	"github.com/challenge-bw-go/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Initialize()

	routes.Setup(r)
	r.Run(":3333")

	log.Printf("Listening on port 3333")
}

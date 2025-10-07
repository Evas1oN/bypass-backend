package main

import (
	"fmt"
	"log"
	"pavrat/bypass/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// todo: Connect to DB

	r := gin.Default()

	r.POST("/addUser", func(ctx *gin.Context) {
		// todo SELECT nextval('ipaddress')
	})

	r.Run(fmt.Sprintf("%s:%s", config.Server.Address, config.Server.Port))
}

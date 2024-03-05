package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	setupRoutes(router)

	router.Run("localhost:8082")
}

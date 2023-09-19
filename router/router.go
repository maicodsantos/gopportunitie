package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()                    // Inicializa o Router utilizando as configurações Default do Gin
	router.GET("/ping", func(c *gin.Context) { // Define a rota /ping
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080") // Executa o servidor na porta 8080 (default
}

package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default() // Inicializa o Router utilizando as configurações Default do Gin

	initializeRoutes(router) // Inicializa as rotas

	router.Run(":8080") // Executa o servidor na porta 8080 (default
}

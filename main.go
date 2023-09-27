package main

import (
	"github.com/maicodsantos/gopportunitie/config"
	"github.com/maicodsantos/gopportunitie/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init() // Inicializa o Banco de Dados
	if err != nil {
		logger.Errorf("config initialization failed: %v", err)
		return
	}

	router.Initialize() // Inicializa o Router

}

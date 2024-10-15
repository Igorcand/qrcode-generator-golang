package main

import (
	"log"
	"qrcode-generator/internal/core/application/use_cases"
	"qrcode-generator/internal/core/adapters/api/router"
)

func main() {
	// Inicialize o serviço de conversão de QRCode
	converterService := application.NewConverterService()

	// Inicialize o roteador da API
	r := router.SetupRouter(&converterService)

	// Inicia o servidor
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

package router

import (
	application "qrcode-generator/internal/core/application/use_cases"
	
	"qrcode-generator/internal/core/adapters/api/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas da API
func SetupRouter(converterService *application.ConverterService) *gin.Engine {
	router := gin.Default()

	// Define a rota para converter e salvar QR Codes
	router.POST("/convert", func(c *gin.Context) {
		handler.HandleConvertQRCode(c, converterService)
	})

	return router
}

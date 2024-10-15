package handler

import (
	"log"
	"net/http"
	application "qrcode-generator/internal/core/application/use_cases"
	"qrcode-generator/internal/core/domain/link"
	"qrcode-generator/internal/core/domain/qrcode"

	"github.com/gin-gonic/gin"
)

// Função handler para a rota de conversão de QR code
func HandleConvertQRCode(c *gin.Context, converterService *application.ConverterService) {
	var req struct {
		Link string `json:"link" binding:"required"`
	}

	// Faz o bind dos dados JSON para a struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Gera o QR code
	link := link.NewLink()
	link.Url = req.Link

	qrcode := qrcode.NewQRCode()
	qrcode.Link = *link
	qrcode.Format = "png"
	qrCode, err := converterService.ConverterLinkToQRCode(qrcode)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar o QR code"})
		return
	}

	// Define o cabeçalho para indicar que a resposta é uma imagem PNG
	c.Header("Content-Type", "image/png")
	c.Writer.Write(qrCode.Image) // Retorna a imagem do QR code como resposta
}

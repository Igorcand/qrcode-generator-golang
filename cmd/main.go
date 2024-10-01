package main

import (
	"log"
	"net/http"
	application "qrcode-generator/internal/core/application/use_cases"
	"qrcode-generator/internal/core/domain/link"
	"qrcode-generator/internal/core/domain/qrcode"

	"github.com/gin-gonic/gin"
)

func main(){
	//mongoClient, err := db.NewMongoClient("mongodb://localhost:27017")
	//if err != nil {
    //    log.Fatalf("Erro ao conectar com o MongoDB: %v", err)
    //}

	//linkRepo := repository.NewMongoLinkRepository(mongoClient, "converter_db", "links")
	//qrcodeRepo := repository.NewMongoQRCodeRepository(mongoClient, "converter_db", "qrcodes")

	converterService := application.NewConverterService()

	router := gin.Default()
	// Define a rota para converter e salvar QR Codes
    router.POST("/convert", func(c *gin.Context) {
        handleConvertQRCode(c, &converterService)
    })

    // Inicia o servidor
    router.Run(":8080")
}


// Função handler para a rota de conversão de QR code
func handleConvertQRCode(c *gin.Context, converterService *application.ConverterService) {
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
    c.Writer.Write(qrCode.Image)  // Retorna a imagem do QR code como resposta
}
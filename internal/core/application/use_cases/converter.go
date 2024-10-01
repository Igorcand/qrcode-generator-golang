package application

import (
	"bytes"
	"errors"
	"log"
	"qrcode-generator/internal/core/domain/link"
	"qrcode-generator/internal/core/domain/qrcode"

	qrcode_gen "github.com/skip2/go-qrcode"
)

type ConverterService struct{}

func NewConverterService() ConverterService {
	return ConverterService{}
}

func (s *ConverterService) ConverterLinkToQRCode(qrcode *qrcode.QRCode) (*qrcode.QRCode, error){
	if !link.IsValidLink(qrcode.Link.Url){
		log.Println("erro validando link")
		return nil, errors.New("invalid link")
	}

	// Gera o QR code diretamente usando a biblioteca go-qrcode
    qrCodeData, err := qrcode_gen.New(qrcode.Link.Url, qrcode_gen.Medium)
    if err != nil {
		log.Println("erro gerando qrcode")
        return nil, err
    }

	//err = qrCodeData.WriteFile(256, "teste.png")
	//if err != nil {
	//	return nil, err
	//}

	var buffer bytes.Buffer
	err = qrCodeData.Write(256, &buffer) // 256 é o tamanho da imagem (em pixels)
	if err != nil {
		log.Println("erro buffer")
		return nil, err
	}

	qrcode.Image = buffer.Bytes()

	err = qrcode.Validate()
	if err != nil{
		log.Println("erro validando")
		return nil, err
	}

	return qrcode, nil

}
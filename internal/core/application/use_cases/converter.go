package application

import (
	"bytes"
	"errors"
	"qrcode-generator/internal/core/domain"
	"github.com/skip2/go-qrcode"
)

type ConverterService struct{}

func NewConverterService() ConverterService {
	return ConverterService{}
}

func (s *ConverterService) ConverterLinkToQRCode(qrCode *domain.QRCode) (*domain.QRCode, error){
	if !domain.IsValidLink(qrCode.Link.Url){
		return nil, errors.New("invalid link")
	}

	// Gera o QR code diretamente usando a biblioteca go-qrcode
    qrCodeData, err := qrcode.New(qrCode.Link.Url, qrcode.Medium)
    if err != nil {
        return nil, err
    }

	//err = qrCodeData.WriteFile(256, "teste.png")
	//if err != nil {
	//	return nil, err
	//}

	var buffer bytes.Buffer
	err = qrCodeData.Write(256, &buffer) // 256 é o tamanho da imagem (em pixels)
	if err != nil {
		return nil, err
	}

	qrCode.Image = buffer.Bytes()

	err = qrCode.Validate()
	if err != nil{
		return nil, err
	}

	return qrCode, nil

}
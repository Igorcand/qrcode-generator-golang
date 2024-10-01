package service

import (
	application "qrcode-generator/internal/core/application/use_cases"
	"qrcode-generator/internal/core/domain/link"
	"qrcode-generator/internal/core/domain/qrcode"
)


type ConverterService struct {
    linkRepo link.LinkRepository
	qrcodeRepo qrcode.QRCodeRepository
}

func NewConverterService( linkRepo link.LinkRepository,qrcodeRepo qrcode.QRCodeRepository) *ConverterService {
    return &ConverterService{
        linkRepo:    linkRepo,
        qrcodeRepo:  qrcodeRepo,
    }
}

func (s *ConverterService) ConvertLinkToQRCode(link string) (*qrcode.QRCode, error) {
    // 1. Valida o link
	converter := application.NewConverterService()
	qrcode := qrcode.NewQRCode()
	qrcode.Link.Url = link
	qrcode, err := converter.ConverterLinkToQRCode(qrcode)
	if err != nil{
		return nil, err
	}
    return qrcode, nil
}
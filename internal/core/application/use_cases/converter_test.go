package application_test

import (
	application "qrcode-generator/internal/core/application/use_cases"
	"qrcode-generator/internal/core/domain"
	"testing"
	//"log"
	"github.com/stretchr/testify/require"
)

func TestConvertLinkToQRCode(t *testing.T){
	link := domain.NewLink()
	link.Url = "https://www.google.com"
	err := link.Validate()
	require.Nil(t, err)

	qrcode := domain.NewQRCode()
	qrcode.Link = *link
	qrcode.Format = "png"

	service := application.NewConverterService()

	qrcode, err = service.ConverterLinkToQRCode(qrcode)

	require.Nil(t, err)
	require.NotNil(t, qrcode)

}
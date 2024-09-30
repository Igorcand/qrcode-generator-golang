package qrcode_test

import (
	"qrcode-generator/internal/core/domain/qrcode"
	"qrcode-generator/internal/core/domain/link"


	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidadeIfQRCodeIsEmpty(t *testing.T){
	qrcode := qrcode.NewQRCode()
	err := qrcode.Validate()
	require.Error(t, err)
}

func TestValidadeIfIdIsNotUuid(t *testing.T){
	qrcode := qrcode.NewQRCode()
	qrcode.ID = "abc"
	qrcode.Image = []byte("image")
	qrcode.Format = "png"
	qrcode.CreatedAt = time.Now()
	qrcode.Link = *link.NewLink()

	err := qrcode.Validate()
	require.Error(t, err)
}

func TestValidadeNotValidFormat(t *testing.T){
	qrcode := qrcode.NewQRCode()
	qrcode.Image = []byte("image")
	qrcode.Format = "abc"
	qrcode.Link = *link.NewLink()
	err := qrcode.Validate()
	require.Error(t, err)
}

func TestQRCodeIsValid(t *testing.T){
	qrcode := qrcode.NewQRCode()
	qrcode.Image = []byte("image")
	qrcode.Format = "png"
	qrcode.Link = *link.NewLink()
	qrcode.Link.Url = "https://www.google.com"
	err := qrcode.Validate()
	require.Nil(t, err)
}
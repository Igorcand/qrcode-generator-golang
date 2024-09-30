package domain_test

import (
	"qrcode-generator/internal/core/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidadeIfQRCodeIsEmpty(t *testing.T){
	qrcode := domain.NewQRCode()
	err := qrcode.Validate()
	require.Error(t, err)
}

func TestValidadeIfIdIsNotUuid(t *testing.T){
	qrcode := domain.NewQRCode()
	qrcode.ID = "abc"
	qrcode.Image = []byte("image")
	qrcode.Format = "png"
	qrcode.CreatedAt = time.Now()
	qrcode.LinkID = "abc"

	err := qrcode.Validate()
	require.Error(t, err)
}

func TestValidadeNotValidFormat(t *testing.T){
	qrcode := domain.NewQRCode()
	qrcode.Image = []byte("image")
	qrcode.Format = "abc"
	qrcode.LinkID = "123"
	err := qrcode.Validate()
	require.Error(t, err)
}

func TestQRCodeIsValid(t *testing.T){
	qrcode := domain.NewQRCode()
	qrcode.Image = []byte("image")
	qrcode.Format = "png"
	qrcode.LinkID = "123"
	err := qrcode.Validate()
	require.Nil(t, err)
}
package qrcode

import (
	"time"
	"errors"
	"github.com/asaskevich/govalidator"
	"qrcode-generator/internal/core/domain/link"
	uuid "github.com/satori/go.uuid"
)

var ALLOWED_FORMATS = []string{"png", "jpeg"}

type QRCode struct{
	ID			string		`valid:"uuid"`
	Image		[]byte		`valid:"-"`
	Format		string		`valid:"notnull"`
	Link		link.Link		`valid:"-"`
	CreatedAt	time.Time	`valid:"-"`
}

func init(){
	govalidator.SetFieldsRequiredByDefault(true)
}

func (qrcode *QRCode) prepare(){
	qrcode.ID = uuid.NewV4().String()
	qrcode.CreatedAt = time.Now()
}

func NewQRCode() *QRCode{
	qrcode := &QRCode{}
	qrcode.prepare()
	return qrcode
}

// Função para verificar se um formato está permitido
func isFormatAllowed(format string) bool {
	for _, allowedFormat := range ALLOWED_FORMATS {
		if format == allowedFormat {
			return true
		}
	}
	return false
}

func (qrcode *QRCode) Validate() error{
	_, err := govalidator.ValidateStruct(qrcode)
	if err != nil{
		return err
	}

	// Verificação manual do campo Image
	if len(qrcode.Image) == 0 {
		return errors.New("image cannot be empty")
	}

	if !isFormatAllowed(qrcode.Format) {
		return errors.New("invalid format")
	}

	// Retorna nil se tudo estiver correto
	return nil
}
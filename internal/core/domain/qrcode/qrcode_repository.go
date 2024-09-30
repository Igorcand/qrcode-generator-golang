package qrcode

type QRCodeRepository interface {
	Save(qrcode *QRCode) error 
	FindByID(id string) (*QRCode, error)
}
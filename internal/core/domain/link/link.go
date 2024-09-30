package link

import (
	"time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Link struct{
	ID			string		`valid:"uuid"`
	Url			string		`valid:"notnull"`
	CreatedAt	time.Time	`valid:"-"`
}

func init(){
	govalidator.SetFieldsRequiredByDefault(true)
}

func (link *Link) prepare(){
	link.ID = uuid.NewV4().String()
	link.CreatedAt = time.Now()
}

func NewLink() *Link{
	link := &Link{}
	link.prepare()
	return link
}

func (link *Link) Validate() error{
	_, err := govalidator.ValidateStruct(link)
	if err != nil{
		return err
	}
	return nil
}

// IsValidLink valida a estrutura do link
func IsValidLink(link string) bool {
    // Lógica simplificada de validação de URL
    return link != "" && len(link) > 5
}
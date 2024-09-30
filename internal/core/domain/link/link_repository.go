package link

type LinkRepository interface {
	Save(link *Link) error 
	FindByURL(url string) (*Link, error)
}
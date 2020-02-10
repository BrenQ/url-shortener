package service

import (
	"context"
	"urlshortener/database"
	"urlshortener/model"
	"urlshortener/repository"
)

// Link interface
type LinkServiceInterface interface {
	Create(url model.Url)
}

type LinkService struct {
	LinkRepository repository.LinkRepositoryInterface
	Db * database.Database
}

func NewLinkService() LinkServiceInterface {
	return LinkService {
		LinkRepository: repository.NewLinkRepository(),
		Db : database.NewConnection(),
	}
}

func (l LinkService) Create (url model.Url ) {

	 link := l.LinkRepository.FindByShortUrl(url.Short)

	if link.Short != "" {
		return
	}

	l.Db.InsertOne(context.Background() ,url)
}







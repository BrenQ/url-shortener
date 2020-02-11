package service

import (
	"context"
	"errors"
	"urlshortener/database"
	"urlshortener/model"
	"urlshortener/repository"
)

// Link interface
type LinkServiceInterface interface {
	Create(url model.Url) error
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

func (l LinkService) Create (url model.Url ) error {
	// Check if exist a short url with same code
	 _, res := l.LinkRepository.FindByShortUrl(url.Short)

	if res == nil {
		return errors.New("url already exists")
	}

	_, err := l.Db.InsertOne(context.Background(), url)

	return err
}







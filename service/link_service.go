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
	Get(url string) (*model.Url, error)
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

func (l LinkService) Create(url model.Url ) error {
	// Check if exist a short url with same code
	 _, res := l.LinkRepository.FindByShortUrl(url.Short)

	if res == nil {
		return errors.New("url already exists")
	}
	// Insert data from url struct
	_, err := l.Db.InsertOne(context.Background(), url)

	return err
}

func (l LinkService ) Get(code string) (*model.Url , error) {
	return l.LinkRepository.FindByShortUrl(code)
}







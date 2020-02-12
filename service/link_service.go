package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"urlshortener/database"
	"urlshortener/model"
	"urlshortener/repository"
)

// Link interface
type LinkServiceInterface interface {
	Create(url model.Url) (*mongo.InsertOneResult,error)
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
// Insert url data
func (l LinkService) Create(url model.Url ) (*mongo.InsertOneResult,error) {
	return l.Db.InsertOne(context.Background(), url)
}
// Get a url data by short code
func (l LinkService ) Get(code string) (*model.Url , error) {
	return l.LinkRepository.FindByShortUrl(code)
}







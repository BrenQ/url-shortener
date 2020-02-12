package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"urlshortener/database"
	"urlshortener/model"
)

// Repository interface
type LinkRepositoryInterface interface {
	FindByShortUrl(url string) (*model.Url, error)
}

type LinkRepository struct {
	Db * database.Database
}

func NewLinkRepository() LinkRepositoryInterface  {
	return LinkRepository{
		Db: database.NewConnection(),
	}
}

func (l LinkRepository ) FindByShortUrl(url string ) (*model.Url,error) {
	var link model.Url

	err := l.Db.FindOne(context.Background(), bson.D{{"short", url}}).
		Decode(&link)

	if err != nil {
		return nil , err
	}

	return &link , nil
}



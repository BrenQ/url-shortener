package service

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"urlshortener/database"
	"urlshortener/model"
	"urlshortener/repository"
)

var (
	SaveRepo func(url model.Url) (*mongo.InsertOneResult,error)
	FindByShortUrlRepo func(code string)(*model.Url,error)
)

func NewLinkRepositoryMock() repository.LinkRepositoryInterface {
	return LinkRepositoryMock {
		database.NewConnection(),
	}
}

func NewLinkServiceMock ()  LinkServiceInterface {
	return LinkService {
		LinkRepository: NewLinkRepositoryMock(),
		Db:             database.NewConnection(),
	}
}

type LinkRepositoryMock struct {
	Db * database.Database
}

func (l LinkRepositoryMock) FindByShortUrl(url string) (*model.Url, error){
	return FindByShortUrlRepo(url)
}

func (l LinkRepositoryMock) Save (url model.Url) (*mongo.InsertOneResult, error){
	return SaveRepo(url)
}

//// Tests

func TestLinkServiceCreate_Success(t *testing.T) {

	SaveRepo = func(url model.Url) (result *mongo.InsertOneResult, err error) {
		return &mongo.InsertOneResult{}, nil
	}

	service := NewLinkServiceMock()
	result, err := service.Create(model.Url{
		Original: "https://www.google.com",
		Short:    "kaIo9a",
	})

	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Unable to generate short url")
	}

	assert.IsType(t, &mongo.InsertOneResult{},result)
}

func TestLinkServiceCreate_Error(t *testing.T) {

	SaveRepo = func(url model.Url) (result *mongo.InsertOneResult, err error) {
		return nil, errors.New("creation fail")
	}

	service := NewLinkServiceMock()
	_, err := service.Create(model.Url{
		Original: "https://www.google.com",
		Short:    "kaIo9a",
	})

	if err == nil {
		t.Errorf("Url was generate")
	}
}

func TestLinkServiceGet_ShortUrlFound(t *testing.T) {

	shortUrl := "xLA13"
	FindByShortUrlRepo = func(code string) (url *model.Url, err error) {
		return &model.Url{
			Original: "https://www.google.com",
			Short:    shortUrl,
		},nil
	}

	service := NewLinkServiceMock()
	res , err := service.Get("xLA13")

	if err != nil {
		t.Errorf("Short url not found")
	}

	assert.Equal(t,shortUrl, res.Short )
}

func TestLinkServiceGet_ShortUrlNotFound(t *testing.T) {

	FindByShortUrlRepo = func(code string) (url *model.Url, err error) {
		return nil,errors.New("not found")
	}

	service := NewLinkServiceMock()
	_ , err := service.Get("xLA13")

	if err == nil {
		t.Errorf("Short url found")
	}
}





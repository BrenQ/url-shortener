package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"testing"
	"urlshortener/database"
	"urlshortener/model"
	"urlshortener/repository"
	"urlshortener/service"
)

var (
	GetLinkService func(code string) (*model.Url , error)
	CreateLinkService func(url model.Url) (*mongo.InsertOneResult, error)
)


func NewLinkControllerMock () LinkControllerInterface {
	return LinkController{ LinkService: NewLinkServiceMock()}
}

func NewLinkServiceMock ()  service.LinkServiceInterface {
	return LinkServiceMock {
		LinkRepository: repository.NewLinkRepository(),
		Db:             database.NewConnection(),
	}
}


// Mock Link service
type LinkServiceMock struct {
	LinkRepository repository.LinkRepositoryInterface
	Db             *database.Database
}

func (lm LinkServiceMock) Create (url model.Url)  (*mongo.InsertOneResult, error) {
	return CreateLinkService(url)
}

func (lm LinkServiceMock) Get (code string) (*model.Url, error) {
	return GetLinkService(code)
}

// Http response
type Response struct {
	Message string`json:"message"`
}

func TestLinkControllerGetLink_CheckSuccessRedirect(t *testing.T) {

	link := "https://www.linkedin.com"
	GetLinkService = func(code string) (url *model.Url, err error) {
		return &model.Url{
			Original: link,
			Short:    "ALXKxaa",
		},nil
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.GET("/:code", controller.GetLink)

	req := httptest.NewRequest("GET", "/ASK928" , nil )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t,link, rec.Header().Get("Location"))
	assert.Equal(t, 301, rec.Code)
}

func TestLinkControllerGetLink_InvalidUriParam(t *testing.T) {

	link := "https://www.linkedin.com"
	GetLinkService = func(code string) (url *model.Url, err error) {
		return &model.Url{
			Original: link,
			Short:    "ALXKxaa",
		},nil
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.GET("/:code", controller.GetLink)

	bodyReq := "{ url : LsPOaA }"

	req := httptest.NewRequest("GET", "/" , bytes.NewBufferString(bodyReq) )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, 404, rec.Code)
}

func TestLinkControllerGetLink_CheckUrlNotFound(t *testing.T) {
	GetLinkService = func(code string) (url *model.Url, err error) {
		return nil, errors.New("not found")
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.GET("/:code", controller.GetLink)

	req := httptest.NewRequest("GET", "/klAoEr" , nil )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	// Response
	var response Response
	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t,"Code not registered", response.Message)
}

func TestLinkControllerCreateLink_CheckCreateSuccess(t *testing.T) {
	CreateLinkService = func(url model.Url) (result *mongo.InsertOneResult, err error) {
		return &mongo.InsertOneResult{}, nil
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.POST("/links", controller.CreateLink)

	bodyReq := `{ "url" : "https://www.facebook.com" }`

	req := httptest.NewRequest("POST", "/links" , bytes.NewBufferString(bodyReq) )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	var response struct{
		Link string
		Orig string
		Short string
	}

	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t , 200 ,rec.Code)
	assert.Equal(t, "https://www.facebook.com", response.Orig )
	assert.NotNil(t,response.Link)
	assert.NotNil(t,response.Short)
}

func TestLinkControllerCreateLink_WithInvalidUrl(t *testing.T) {

	CreateLinkService = func(url model.Url) (result *mongo.InsertOneResult, err error) {
		return &mongo.InsertOneResult{}, nil
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.POST("/links", controller.CreateLink)

	bodyReq := `{ "url" : "htt1://www.facebook.com" }`

	req := httptest.NewRequest("POST", "/links" , bytes.NewBufferString(bodyReq) )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	var response Response

	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t , http.StatusBadRequest, rec.Code)
	assert.Equal(t,"Url invalid", response.Message)
}

func TestLinkControllerCreateLink_WithInvalidBody(t *testing.T) {

	CreateLinkService = func(url model.Url) (result *mongo.InsertOneResult, err error) {
		return &mongo.InsertOneResult{}, nil
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.POST("/links", controller.CreateLink)

	bodyReq := `{ "link" : "htt1://www.facebook.com" }`

	req := httptest.NewRequest("POST", "/links" , bytes.NewBufferString(bodyReq) )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	var response struct{
		Message string
		Err string
	}

	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t , http.StatusBadRequest, rec.Code)
	assert.Equal(t,"Invalid request", response.Message)
	assert.NotNil(t,response.Err)
}

func TestLinkControllerCreateLink_ShortUrlNonCreated(t *testing.T) {

	CreateLinkService = func(url model.Url) (result *mongo.InsertOneResult, err error) {
		return nil, errors.New("non created")
	}

	controller :=  NewLinkControllerMock()

	r := gin.Default()
	r.POST("/links", controller.CreateLink)

	bodyReq := `{ "url" : "https://www.facebook.com" }`

	req := httptest.NewRequest("POST", "/links" , bytes.NewBufferString(bodyReq) )
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	var response Response

	_ = json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t , http.StatusUnprocessableEntity, rec.Code)
	assert.Equal(t,"Unable to process request. Try Again", response.Message)
}

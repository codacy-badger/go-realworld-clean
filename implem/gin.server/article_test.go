package server_test

import (
	"testing"

	"net/http/httptest"

	"net/http"

	server "github.com/err0r500/go-realworld-clean/implem/gin.server"
	jwt "github.com/err0r500/go-realworld-clean/implem/jwt.authHandler"
	"github.com/err0r500/go-realworld-clean/implem/mock.uc"
	"github.com/err0r500/go-realworld-clean/testData"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/baloo.v3"
)

const artPath = "/api/articles"

func TestRouterHandler_articlePost(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedArticle := testData.Article("jane")
	ucHandler := uc.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().
		ArticlePost(gomock.Any()).
		Return(&expectedArticle, nil).
		Times(1)

	jwtHandler := jwt.NewTokenHandler("mySalt")

	gE := gin.Default()
	server.NewRouter(ucHandler, jwtHandler).SetRoutes(gE)
	ts := httptest.NewServer(gE)
	defer ts.Close()

	authToken, err := jwtHandler.GenUserToken(testData.User("jane").Name)
	assert.NoError(t, err)

	baloo.New(ts.URL).
		Post(artPath).
		AddHeader("Authorization", authToken).
		BodyString(`{
  "article": {
    "title": "` + expectedArticle.Title + `",
    "description": "` + expectedArticle.Description + `",
    "body": "` + expectedArticle.Body + `",
    "tagList": [
      "` + expectedArticle.TagList[0] + `"
    ]
  }
}`).
		Expect(t).
		Status(http.StatusCreated).
		Done()
}

//
//func TestRouterHandler_articlePut(t *testing.T) {
//	type fields struct {
//		ucHandler   uc.Handler
//		authHandler uc.AuthHandler
//		Logger      uc.Logger
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rH := RouterHandler{
//				ucHandler:   tt.fields.ucHandler,
//				authHandler: tt.fields.authHandler,
//				Logger:      tt.fields.Logger,
//			}
//			rH.articlePut(tt.args.c)
//		})
//	}
//}
//
//func TestRouterHandler_articleGet(t *testing.T) {
//	type fields struct {
//		ucHandler   uc.Handler
//		authHandler uc.AuthHandler
//		Logger      uc.Logger
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rH := RouterHandler{
//				ucHandler:   tt.fields.ucHandler,
//				authHandler: tt.fields.authHandler,
//				Logger:      tt.fields.Logger,
//			}
//			rH.articleGet(tt.args.c)
//		})
//	}
//}
//
//func TestRouterHandler_articleDelete(t *testing.T) {
//	type fields struct {
//		ucHandler   uc.Handler
//		authHandler uc.AuthHandler
//		Logger      uc.Logger
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rH := RouterHandler{
//				ucHandler:   tt.fields.ucHandler,
//				authHandler: tt.fields.authHandler,
//				Logger:      tt.fields.Logger,
//			}
//			rH.articleDelete(tt.args.c)
//		})
//	}
//}

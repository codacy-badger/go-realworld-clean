package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/err0r500/go-realworld-clean/implem/gin.server"
	"github.com/err0r500/go-realworld-clean/implem/jwt.authHandler"
	logger "github.com/err0r500/go-realworld-clean/implem/logrus.logger"
	"github.com/err0r500/go-realworld-clean/implem/mock.uc"
	"github.com/err0r500/go-realworld-clean/testData"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/baloo.v3"
)

var userPutPath = "/api/users"

func TestUserPut_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jane := testData.User("jane")
	ucHandler := uc.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().
		UserEdit(testData.User("rick").Name, gomock.Any()).
		Return(&jane, nil).
		Times(1)

	gE := gin.Default()
	jwtHandler := jwt.NewTokenHandler("mySalt")
	router := server.NewRouter(ucHandler, jwtHandler)
	router.Logger = logger.SimpleLogger{}
	router.SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	token, err := jwtHandler.GenUserToken(testData.User("rick").Name)
	assert.NoError(t, err)
	baloo.New(ts.URL).
		Put(userPutPath).
		AddHeader("Authorization", token).
		BodyString(`{
  			"user": {
				"bio": "` + testData.User("rick").Email + `",
				"name": "` + testData.User("rick").Name + `"
  			}
		}`).
		Expect(t).
		Status(http.StatusOK).
		Done()
}

package server

import (
	"net/http"

	formatter "github.com/err0r500/go-realworld-clean/implem/json.formatter"
	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) userGet(c *gin.Context) {
	log := rH.log(c.Request.URL.Path)

	userID, err := rH.getUserName(c)
	if err != nil {
		log(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	user, token, err := rH.ucHandler.UserGet(userID)
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": formatter.NewUserResp(*user, token)})
}

package server

import (
	"net/http"

	"github.com/err0r500/go-realworld-clean/implem/json.formatter"
	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) profileFollowDelete(c *gin.Context) {
	log := rH.log(c.Request.URL.Path)

	userID, err := rH.getUserName(c)
	if err != nil {
		log(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	user, err := rH.ucHandler.ProfileUpdateFollow(userID, c.Param("username"), false)
	if err != nil {
		log(err)

		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, formatter.NewProfileFromDomain(*user, false))
}

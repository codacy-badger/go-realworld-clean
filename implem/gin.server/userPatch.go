package server

import (
	"net/http"

	"github.com/err0r500/go-realworld-clean/uc"
	"github.com/gin-gonic/gin"
)

//New user details.
type userPutRequest struct {
	User struct {
		Email    string `json:"email,omitempty"`
		Name     string `json:"username,omitempty"`
		Bio      string `json:"bio,omitempty"`
		Image    string `json:"image,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"user,required"`
}

func (rH RouterHandler) userPatch(c *gin.Context) {
	log := rH.log(c.Request.URL.Path)
	userID, err := rH.getUserName(c)
	if err != nil {
		log(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	req := &userPutRequest{}
	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := rH.ucHandler.UserEdit(userID, map[uc.UpdatableProperty]*string{
		uc.Email:     &req.User.Email,
		uc.Name:      &req.User.Name,
		uc.Bio:       &req.User.Bio,
		uc.ImageLink: &req.User.Image,
		uc.Password:  &req.User.Password,
	})
	if err != nil {
		log(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
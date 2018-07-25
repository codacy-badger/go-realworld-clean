package server

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/err0r500/go-realworld-clean/uc"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	ucHandler   uc.Handler
	authHandler uc.AuthHandler
	Logger      uc.Logger
}

func NewRouter(i uc.Handler, auth uc.AuthHandler) RouterHandler {
	return RouterHandler{
		ucHandler:   i,
		authHandler: auth,
	}
}

func NewWithLogger(i uc.Handler, auth uc.AuthHandler, logger uc.Logger) RouterHandler {
	return RouterHandler{
		ucHandler:   i,
		authHandler: auth,
		Logger:      logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")

	profiles := api.Group("/profiles")
	profiles.GET("/:username", rH.profileGet)                                        // Get a profile of a user of the system. Auth is optional
	profiles.POST("/:username/follow", rH.jwtMiddleware(), rH.profileFollowPost)     // Follow a user by username
	profiles.DELETE("/:username/follow", rH.jwtMiddleware(), rH.profileFollowDelete) // Unfollow a user by username

	users := api.Group("/users")
	users.POST("", rH.userPost)            // Register a new user
	users.POST("/login", rH.userLoginPost) // Login for existing user

	users.GET("", rH.jwtMiddleware(), rH.userGet)     // Gets the currently logged-in user
	users.PUT("", rH.jwtMiddleware(), rH.userPatch)   // WARNING : it's a in fact a PATCH request in the API contract !!!
	users.PATCH("", rH.jwtMiddleware(), rH.userPatch) // just in case it's fixed one day....

	articles := api.Group("/articles")
	// GET articleCollection
	articles.GET("", rH.articlesFilteredGet)
	api.GET("articlesFeed", rH.jwtMiddleware(), rH.articlesFeedGet) //unused, see GET /:slug below
	articles.POST("", rH.jwtMiddleware(), rH.articlePost)
	articles.GET("/:slug", func(c *gin.Context) { // ugly api contract !
		if c.Param("slug") == "feed" {
			rH.jwtMiddleware()(c)
			rH.articlesFeedGet(c)
			return
		}

		rH.articleGet(c)
	})
	articles.PUT("/:slug", rH.jwtMiddleware(), rH.articlePut)
	articles.DELETE("/:slug", rH.jwtMiddleware(), rH.articleDelete)
	articles.POST("/:slug/favorite", rH.jwtMiddleware(), rH.articleFavoritePost)
	articles.DELETE("/:slug/favorite", rH.jwtMiddleware(), rH.articleFavoriteDelete)
	////article comments
	articles.GET("/:slug/comments", rH.articleCommentsGet)
	articles.POST("/:slug/comments", rH.jwtMiddleware(), rH.articleCommentPost)
	articles.DELETE("/:slug/comments/:id", rH.jwtMiddleware(), rH.articleCommentDelete)

	//tags
	api.GET("/tags", rH.tagsGet)
}

// TODO : implement these routes

func (RouterHandler) articleFavoritePost(c *gin.Context)   {}
func (RouterHandler) articleFavoriteDelete(c *gin.Context) {}
func (RouterHandler) articleCommentsGet(c *gin.Context)    {}
func (RouterHandler) articleCommentPost(c *gin.Context)    {}
func (RouterHandler) articleCommentDelete(c *gin.Context)  {}
func (RouterHandler) tagsGet(c *gin.Context)               {}

const userNameKey = "userNameKey"

func (rH RouterHandler) jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName, err := rH.authHandler.GetUserName(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(userNameKey, userName)
		c.Next()
	}
}

func (RouterHandler) getUserName(c *gin.Context) (string, error) {
	userName, ok := c.Keys[userNameKey].(string)
	if !ok {
		return "", errors.New("userNameKey not in gin Context")
	}
	if userName == "" {
		return "", errors.New("empty userNameKey in gin Context")
	}
	return userName, nil
}

// log is used to "partially apply" the title to the rH.logger.Log function
// so we can see in the logs from which route the log comes from
func (rH RouterHandler) log(title string) func(...interface{}) {
	return func(logs ...interface{}) {
		rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}

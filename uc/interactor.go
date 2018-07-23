package uc

import (
	"github.com/err0r500/go-realworld-clean/domain"
)

type Handler interface {
	ProfileGet(userName string) (profile *domain.Profile, err error)
	ProfileUpdateFollow(loggedInUserID, username string, follow bool) (user *domain.User, err error)

	UserCreate(username, email, password string) (user *domain.User, token string, err error)
	UserLogin(email, password string) (user *domain.User, token string, err error)
	UserGet(userID string) (user *domain.User, token string, err error)
	UserEdit(userID string, newUser map[UpdatableProperty]*string) (user *domain.User, err error)

	ArticlesFeed(username string, limit, offset int) ([]domain.Article, error)
}

// NewHandler : the interactor constructor, use this in order to avoid null pointers at runtime
func NewHandler(logger Logger, uRW UserRW, arw ArticleRW, validator UserValidator, handler AuthHandler) Handler {
	return interactor{
		logger:        logger,
		userRW:        uRW,
		articleRW:     arw,
		userValidator: validator,
		authHandler:   handler,
	}
}

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases
type interactor struct {
	logger        Logger
	userRW        UserRW
	articleRW     ArticleRW
	userValidator UserValidator
	authHandler   AuthHandler
}

// Logger : only used to log stuff
type Logger interface {
	Log(...interface{})
}

type AuthHandler interface {
	GenUserToken(userName string) (token string, err error)
	GetUserName(token string) (userID string, err error)
}

type UserRW interface {
	Create(username, email, password string) (*domain.User, error)
	GetByName(userName string) (*domain.User, error)
	GetByEmailAndPassword(email, password string) (*domain.User, error)
	Save(user domain.User) error
}

type ArticleRW interface {
	GetByAuthorsNameOrderedByMostRecentAsc(usernames []string) ([]domain.Article, error)
}

type UserValidator interface {
	CheckUser(user domain.User) error
}

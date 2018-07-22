package uc

import (
	"log"

	"github.com/err0r500/go-realworld-clean/uc"
	"github.com/golang/mock/gomock"
)

// MockedInteractor : is used in order to update its properties accordingly to each test conditions
type MockedInteractor struct {
	Logger        uc.Logger
	UserRW        *MockUserRW
	UserValidator *MockUserValidator
	AuthHandler   *MockAuthHandler
}

type simpleLogger struct{}

func (simpleLogger) Log(logs ...interface{}) {
	log.Println(logs)
}

//NewMockedInteractor : the MockedInteractor constructor
func NewMockedInteractor(mockCtrl *gomock.Controller) MockedInteractor {
	return MockedInteractor{
		Logger:        simpleLogger{},
		UserRW:        NewMockUserRW(mockCtrl),
		UserValidator: NewMockUserValidator(mockCtrl),
		AuthHandler:   NewMockAuthHandler(mockCtrl),
	}
}

//GetUCHandler : returns a uc.interactor in order to call its methods aka the use cases to test
func (i MockedInteractor) GetUCHandler() uc.Handler {
	return uc.NewHandler(i.Logger, i.UserRW, i.UserValidator, i.AuthHandler)
}

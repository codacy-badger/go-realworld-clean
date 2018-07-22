package uc

import (
	"github.com/golang/mock/gomock"
	
	"github.com/err0r500/go-realworld-clean/uc"
)

// MockedInteractor is used in order to update its properties accordingly to 
// each test conditions
type MockedInteractor struct {}

// NewMockedInteractor is the MockedInteractor constructor
func NewMockedInteractor(mockCtrl *gomock.Controller) MockedInteractor {
	return MockedInteractor{}
}

//GetUCInteractor : returns a uc.Interactor in order to call its methods aka the use cases to test
func (i MockedInteractor) GetUCInteractor() uc.Interactor {
	return uc.Interactor{}
}

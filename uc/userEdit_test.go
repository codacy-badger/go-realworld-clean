package uc_test

import (
	"testing"

	mock "github.com/err0r500/go-realworld-clean/implem/mock.uc"
	"github.com/err0r500/go-realworld-clean/testData"
	"github.com/err0r500/go-realworld-clean/uc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEditUser_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedUser := testData.User("rick")
	expectedUser.Email = testData.User("jane").Email
	rick := testData.User("rick")
	jane := testData.User("jane")
	i := mock.NewMockedInteractor(mockCtrl)
	i.UserRW.EXPECT().GetByName(rick.Name).Return(&rick, nil).Times(1)
	i.UserValidator.EXPECT().CheckUser(expectedUser).Return(nil).Times(1)
	i.UserRW.EXPECT().Save(expectedUser).Return(nil).Times(1)

	retUser, err := i.GetUCHandler().UserEdit(rick.Name, map[uc.UpdatableProperty]*string{
		uc.Email: &jane.Email,
		uc.Bio:   testData.User("jane").Bio, //nil
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, *retUser)

}

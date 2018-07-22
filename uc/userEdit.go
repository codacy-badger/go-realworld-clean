package uc

import (
	"github.com/err0r500/go-realworld-clean/domain"
)

type UpdatableProperty int

const (
	Email UpdatableProperty = iota
	Name
	Bio
	ImageLink
	Password
)
//UserEdit(userID string, newUser map[UpdatableProperty]*string) (user *domain.User, err error)
func (i interactor) UserEdit(userName string, fieldsToUpdate map[UpdatableProperty]*string) (*domain.User, error) {
	user, err := i.userRW.GetByName(userName)
	if err != nil {
		return nil, err
	}
	if user.Name != userName {
		return nil, errWrongUser
	}
	if user == nil {
		return nil, errUserNotFound
	}

	domain.UpdateUser(user,
		domain.SetUserName(fieldsToUpdate[Name]),
		domain.SetUserEmail(fieldsToUpdate[Email]),
		domain.SetUserBio(fieldsToUpdate[Bio]),
		domain.SetUserImageLink(fieldsToUpdate[ImageLink]),
		domain.SetUserPassword(fieldsToUpdate[Password]),
	)

	if err := i.userValidator.CheckUser(*user); err != nil {
		return nil, err
	}

	if err := i.userRW.Save(*user); err != nil {
		return nil, err
	}

	return user, nil
}

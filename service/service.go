package service

import (
	"go_mysql/domain"
	"go_mysql/utils"

	"github.com/rs/xid"
)

func CreateUser(user domain.User) *utils.Resterr {
	if err := user.Vaildate(); err != nil {
		return err
	}
	guid := xid.New()
	user.ID = guid.String()
	pass := utils.HashPasswordMD5(user.Password)
	user.Password = pass
	insertErr := user.Create()
	if insertErr != nil {
		return insertErr
	}
	return nil
}

func FindUser(user domain.User) (*domain.User, *utils.Resterr) {
	if err := user.EmailVaildate(); err != nil {
		return nil, err
	}
	result, insertErr := user.FindUser()
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func FindAlluser(user domain.User) ([]domain.User, *utils.Resterr) {

	result, insertErr := user.FindAlluser()
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func DeleteUserByEmail(user domain.User) (int64, *utils.Resterr) {
	if err := user.EmailVaildate(); err != nil {
		return 0, err
	}
	result, insertErr := user.DeleteUser()
	if insertErr != nil {
		return 0, insertErr
	}
	return result, nil
}

func UpdateUserByEmail(user domain.User) (int64, *utils.Resterr) {
	if err := user.EmailVaildate(); err != nil {
		return 0, err
	}
	result, insertErr := user.UpdateUser()
	if insertErr != nil {
		return 0, insertErr
	}
	return result, nil
}

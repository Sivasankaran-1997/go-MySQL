package domain

import (
	"go_mysql/utils"
)

func (user *User) Create() *utils.Resterr {

	var emailcount int64
	Instance.Model(&User{}).Where("email = ?", user.Email).Count(&emailcount)

	if emailcount > 0 {
		return utils.BadRequest("Email Already  Register")
	}

	value := Instance.Create(&user)

	if value.Error != nil {
		return utils.InternalErr("Can't Insert the Database")
	}

	return nil
}

func (user *User) FindUser() (*User, *utils.Resterr) {

	var emailcount int64
	Instance.Model(&User{}).Where("email = ?", user.Email).Count(&emailcount)

	if emailcount == 0 {
		return nil, utils.BadRequest("Email Not Found")
	}

	Instance.Find(&user, "email = ?", user.Email)
	user.Password = ""

	return user, nil
}

func (user User) FindAlluser() ([]User, *utils.Resterr) {

	var users []User
	Instance.Find(&users)

	return users, nil
}

func (user *User) DeleteUser() (int64, *utils.Resterr) {

	var emailcount int64
	Instance.Model(&User{}).Where("email = ?", user.Email).Count(&emailcount)

	if emailcount == 0 {
		return 0, utils.BadRequest("Email Not Found")
	}

	res := Instance.Delete(&user, "email = ?", user.Email)

	return res.RowsAffected, nil
}

func (user *User) UpdateUser() (int64, *utils.Resterr) {

	var emailcount int64
	Instance.Model(&User{}).Where("email = ?", user.Email).Count(&emailcount)

	if emailcount == 0 {
		return 0, utils.BadRequest("Email Not Found")
	}

	res := Instance.Model(&User{}).Where("email = ?", user.Email).Updates(map[string]interface{}{"Name": user.Name})

	return res.RowsAffected, nil
}

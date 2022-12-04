package domain

import (
	"go_mysql/utils"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type User struct {
	ID       string `json:id"`
	Name     string `json:name bson:"name,omitempty"`
	Email    string `json:email bson:"email,omitempty"`
	Password string `json:password bson:"password,omitempty"`
}

func (user *User) Vaildate() *utils.Resterr {
	if strings.TrimSpace(user.Name) == "" {
		return utils.BadRequest("Name Required")
	}

	if strings.TrimSpace(user.Email) == "" {
		return utils.BadRequest("Email Required")
	}

	if strings.TrimSpace(user.Password) == "" {
		return utils.BadRequest("PassWord Required")
	}

	return nil
}

func (user *User) EmailVaildate() *utils.Resterr {

	if strings.TrimSpace(user.Email) == "" {
		return utils.BadRequest("Email Required")
	}

	return nil
}

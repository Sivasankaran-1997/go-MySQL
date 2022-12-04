package contoller

import (
	"encoding/json"
	"go_mysql/domain"
	"go_mysql/service"
	"go_mysql/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		json.NewEncoder(w).Encode(resterr)
		return
	}

	inserterr := service.CreateUser(newUser)

	if inserterr != nil {
		json.NewEncoder(w).Encode(inserterr)
		return
	}
	json.NewEncoder(w).Encode(newUser)

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		json.NewEncoder(w).Encode(resterr)
		return
	}

	findvalue, finderr := service.FindUser(newUser)

	if finderr != nil {
		json.NewEncoder(w).Encode(finderr)
		return
	}
	json.NewEncoder(w).Encode(findvalue)

}

func FindAlluser(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User

	findvalue, finderr := service.FindAlluser(newUser)

	if finderr != nil {
		json.NewEncoder(w).Encode(finderr)
		return
	}
	json.NewEncoder(w).Encode(findvalue)

}

func DeleteUserByEmail(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		json.NewEncoder(w).Encode(resterr)
		return
	}

	deletevalue, deleteerr := service.DeleteUserByEmail(newUser)

	if deleteerr != nil {
		json.NewEncoder(w).Encode(deleteerr)
		return
	}
	json.NewEncoder(w).Encode(deletevalue)

}

func UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		json.NewEncoder(w).Encode(resterr)
		return
	}

	updatevalue, updateerr := service.UpdateUserByEmail(newUser)

	if updateerr != nil {
		json.NewEncoder(w).Encode(updateerr)
		return
	}
	json.NewEncoder(w).Encode(updatevalue)

}

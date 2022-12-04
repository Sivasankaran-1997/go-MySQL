package app

import "go_mysql/contoller"

func Routers() {
	r.HandleFunc("/users/create", contoller.CreateUser).Methods("POST")
	r.HandleFunc("/users/getbyemail", contoller.FindUser).Methods("GET")
	r.HandleFunc("/users/getall", contoller.FindAlluser).Methods("GET")
	r.HandleFunc("/users/deletebyemail", contoller.DeleteUserByEmail).Methods("DELETE")
	r.HandleFunc("/users/updatebyemail", contoller.UpdateUserByEmail).Methods("PATCH")
}

package controller

import (
	"encoding/json"
	"internetshop/helper"
	"internetshop/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	repository helper.UserRepository
}

func NewUserRepository(repository helper.UserRepository) *UserController {
	return &UserController{repository: repository}
}

func (c *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.UnregUser

	_ = json.NewDecoder(r.Body).Decode(&user) //check what instade of this var
	c.repository.AddUser(user)
}

// gittest
func CreateUnregUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user model.UnregUser
	_ = json.NewDecoder(r.Body).Decode(&user) //check what instade of this var

	u := helper.CreateUserInDB(user)
	// var s string = user.ID + user.Basket
	json.NewEncoder(w).Encode(u)

}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	users := helper.GetAllUsers()
	json.NewEncoder(w).Encode(users)

}
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	searchedUser := helper.GetOneUser(params["id"])
	if searchedUser == nil {
		json.NewEncoder(w).Encode("No user with id: " + params["id"])
		return
	}

	//json.NewEncoder(w).Encode(searchedUserInarray)
	json.NewEncoder(w).Encode(searchedUser)

}
func DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	helper.DeleteOneUser(params["id"])
	json.NewEncoder(w).Encode("User with id :" + params["id"] + " was deleted")

}
func DeleteALlUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := helper.DeleteALlUsers()
	json.NewEncoder(w).Encode(strconv.Itoa(int(count)) + "users was deleted")
}

func AddCommodityToUserBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	//fmt.Println()
	helper.AddCommodityToUserBasket(params["id"], params["commodity"])
}
func GetSetCoockies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	helper.GetSetCoockies(w, r)
}

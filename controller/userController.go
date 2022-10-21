package controller

import (
	"encoding/json"
	"fmt"
	helper "internetshop/helper"
	"internetshop/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type UserController struct {
	repository helper.UserRepository
}

func NewUserRepository(repository helper.UserRepository) *UserController {
	return &UserController{repository: repository}
}

// gittest
func (c *UserController) CreateUnregUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user model.UnregUser
	_ = json.NewDecoder(r.Body).Decode(&user) //check what instade of this var

	//u := helper.CreateUserInDB(user)
	u := c.repository.CreateUserInDB(user)
	// var s string = user.ID + user.Basket
	json.NewEncoder(w).Encode(u)

}
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	//users := helper.GetAllUsers()
	users, err := c.repository.GetAllUsers()
	if err != nil {
		json.NewEncoder(w).Encode(users)
	}
	json.NewEncoder(w).Encode(users)

}
func (c *UserController) GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	//searchedUser := helper.GetOneUser(params["id"])
	searchedUser, err := c.repository.GetOneUser(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("No user with id: " + params["id"])
		return
	}
	//json.NewEncoder(w).Encode(searchedUserInarray)
	json.NewEncoder(w).Encode(searchedUser)

}
func (c *UserController) DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	//helper.DeleteOneUser(params["id"])
	c.repository.DeleteOneUser(params["id"])
	json.NewEncoder(w).Encode("User with id :" + params["id"] + " was deleted")

}
func (c *UserController) DeleteALlUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	//count := helper.DeleteALlUsers()
	count := c.repository.DeleteALlUsers()
	json.NewEncoder(w).Encode(strconv.Itoa(int(count)) + "users was deleted")
}

func (c *UserController) AddCommodityToUserBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	//helper.AddCommodityToUserBasket(params["id"], params["commodity"])
	c.repository.AddCommodityToUserBasket(params["id"], params["commodity"])
}

func (c *UserController) SetCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	//user := CreateUnregUserInDB()
	// user := UserRepository.CreateUnregUserInDB(NewMongo())
	userID := c.repository.CreateUnregUserInDB()

	fmt.Println(userID)
	cookie := http.Cookie{Name: "id", Value: userID, Expires: expiration}
	http.SetCookie(w, &cookie)
	return &cookie
}

func (c *UserController) GetCoockie(w http.ResponseWriter, r *http.Request) (*http.Cookie, bool) {
	cookieid, _ := r.Cookie("id")
	if cookieid == nil {
		fmt.Println("cookie is nil")
		return cookieid, false
	} else {
		return cookieid, true
	}

}
func (c *UserController) GetSetCoockies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// rcookie := r.Cookies()
	//cookie, check := GetCoockie(w, r)
	cookie, check := c.GetCoockie(w, r)

	fmt.Println(cookie, check)
	if !check {
		a := c.SetCookie(w, r)
		fmt.Println(a)
		//fmt.Println(c)cookie
	}

}

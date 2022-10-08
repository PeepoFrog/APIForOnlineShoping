package handler

import (
	"fmt"
	"internetshop/controller"
	"net/http"
)

type Mongo struct {
}
type Postgre struct {
}
type Controller interface {
	CreateUnregUser()
}
handler
func (m Mongo) CreateUnregUser(w http.ResponseWriter, r *http.Request) {
	controller.CreateUnregUser(w, r)

	//метод из контроллера который добавляет в монго юзера
}
func (m Postgre) AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(m)

	//метод из контроллера который добавляет в постгре юзера
}

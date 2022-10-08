package router

import (
	"internetshop/controller"
	"internetshop/helper"

	"github.com/gorilla/mux"
)
handlers
func Router() *mux.Router {
	//
	//
	// commodities routers
	router := mux.NewRouter()
	router.HandleFunc("/api/commodity", controller.GetAllCommodities).Methods("GET")
	router.HandleFunc("/api/commodity/{id}", controller.GetOneCommodity).Methods("GET")
	router.HandleFunc("/api/commodity", controller.CreateCommodity).Methods("POST")
	router.HandleFunc("/api/commodity/{id}&{price}", controller.SetPrice).Methods("PUT")
	router.HandleFunc("/api/commodity/{id}&{quantity}", controller.SetQuantity).Methods("PUT")
	router.HandleFunc("/api/commodity/{id}", controller.DeleteOneCommodity).Methods("DELETE")
	router.HandleFunc("/api/commodity", controller.DeleteALlCommodities).Methods("DELETE")
	//
	//
	// testing
	router.HandleFunc("/api/coockie", controller.GetSetCoockies).Methods("GET")
	router.HandleFunc("/api/user/testing/{userID}&{itemID}", helper.Testing).Methods("GET")

	//router.HandleFunc("/api/user/adduser").Methods("GET")

	//
	//
	// users routers
	router.HandleFunc("/api/user", controller.CreateUnregUser).Methods("POST")
	router.HandleFunc("/api/user", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.GetOneUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.DeleteOneUser).Methods("DELETE")
	router.HandleFunc("/api/user", controller.DeleteALlUsers).Methods("DELETE")
	router.HandleFunc("/api/user/{id}&{commodity}", controller.AddCommodityToUserBasket).Methods("PUT")

	//
	return router

}

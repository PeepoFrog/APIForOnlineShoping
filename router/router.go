package router

import (
	"internetshop/controller"
	"internetshop/helper"

	"github.com/gorilla/mux"
)

func init() {

}
func Router() *mux.Router {
	router := mux.NewRouter()
	//
	//
	// commodities routers
	mongo := helper.NewMongo()
	_ = mongo
	postgres := helper.NewPostgre()
	_ = postgres

	userController := controller.NewUserRepository(mongo)
	router.HandleFunc("/api/usertest", userController.AddUser).Methods("POST")

	//

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
	router.HandleFunc("/api/coockie", userController.GetSetCoockies).Methods("GET")
	router.HandleFunc("/api/user/testing/{userID}&{itemID}", helper.Testing).Methods("GET")

	//
	//
	// users routers

	router.HandleFunc("/api/user", userController.CreateUnregUser).Methods("POST")
	router.HandleFunc("/api/user", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", userController.GetOneUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", userController.DeleteOneUser).Methods("DELETE")
	router.HandleFunc("/api/user", userController.DeleteALlUsers).Methods("DELETE")
	router.HandleFunc("/api/user/{id}&{commodity}", userController.AddCommodityToUserBasket).Methods("PUT")

	//
	return router

}

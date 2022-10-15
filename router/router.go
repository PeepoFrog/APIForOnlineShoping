package router

import (
	"internetshop/controller"
	helper "internetshop/helper"

	"github.com/gorilla/mux"
)

func init() {

}
func Router() *mux.Router {
	router := mux.NewRouter()
	mongo := helper.NewMongo()
	_ = mongo
	postgres := helper.NewPostgre()
	_ = postgres
	mongoUserController := controller.NewUserRepository(mongo)
	mongoCommodityController := controller.NewCommodityRepository(mongo)
	router.HandleFunc("/api/usertest", mongoUserController.AddUser).Methods("POST")
	//
	//
	// commodities routers
	router.HandleFunc("/api/commodity", mongoCommodityController.GetAllCommodities).Methods("GET")
	router.HandleFunc("/api/commodity/{id}", mongoCommodityController.GetOneCommodity).Methods("GET")
	router.HandleFunc("/api/commodity", mongoCommodityController.CreateCommodity).Methods("POST")
	router.HandleFunc("/api/commodity/{id}&{price}", mongoCommodityController.SetPrice).Methods("PUT")
	router.HandleFunc("/api/commodity/{id}&{quantity}", mongoCommodityController.SetQuantity).Methods("PUT")
	router.HandleFunc("/api/commodity/{id}", mongoCommodityController.DeleteOneCommodity).Methods("DELETE")
	router.HandleFunc("/api/commodity", mongoCommodityController.DeleteALlCommodities).Methods("DELETE")
	//
	//
	// testing
	router.HandleFunc("/api/coockie", mongoUserController.GetSetCoockies).Methods("GET")
	router.HandleFunc("/api/user/testing/{userID}&{itemID}", helper.Testing).Methods("GET")
	//
	//
	// users routers
	router.HandleFunc("/api/user", mongoUserController.CreateUnregUser).Methods("POST")
	router.HandleFunc("/api/user", mongoUserController.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", mongoUserController.GetOneUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", mongoUserController.DeleteOneUser).Methods("DELETE")
	router.HandleFunc("/api/user", mongoUserController.DeleteALlUsers).Methods("DELETE")
	router.HandleFunc("/api/user/{id}&{commodity}", mongoUserController.AddCommodityToUserBasket).Methods("PUT")
	//
	return router

}

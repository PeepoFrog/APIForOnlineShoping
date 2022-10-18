package router

import (
	"internetshop/controller"
	helper "internetshop/helper"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	mongo := helper.NewMongo()
	_ = mongo
	postgres := helper.NewPostgre()

	userController := controller.NewUserRepository(mongo)
	commodityController := controller.NewCommodityRepository(mongo)
	postgreCommodityComtroller := controller.NewCommodityRepository(postgres)
	router.HandleFunc("/api/usertest", userController.AddUser).Methods("POST")

	//
	//
	// commodities routers
	router.HandleFunc("/api/commodity", commodityController.GetAllCommodities).Methods("GET")
	router.HandleFunc("/api/commodity/{id}", commodityController.GetOneCommodity).Methods("GET")
	router.HandleFunc("/api/commodity", commodityController.CreateCommodity).Methods("POST")
	router.HandleFunc("/api/commodity/price/{id}&{price}", commodityController.SetPrice).Methods("PUT")
	router.HandleFunc("/api/commodity/quantity/{id}&{quantity}", commodityController.SetQuantity).Methods("PUT")
	router.HandleFunc("/api/commodity/{id}", commodityController.DeleteOneCommodity).Methods("DELETE")
	router.HandleFunc("/api/commodity", commodityController.DeleteALlCommodities).Methods("DELETE")
	//
	//
	// testing
	router.HandleFunc("/test/cookie", userController.GetSetCoockies).Methods("GET")
	router.HandleFunc("/test/{userID}&{itemID}", helper.Testing).Methods("GET")
	router.HandleFunc("/test", postgreCommodityComtroller.GetAllCommodities).Methods("GET")
	router.HandleFunc("/test{id}&{price}", postgreCommodityComtroller.GetAllCommodities).Methods("GET")

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

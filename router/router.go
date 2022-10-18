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

	mongoUserController := controller.NewUserRepository(mongo)
	commodityController := controller.NewCommodityRepository(mongo)
	postgreCommodityComtroller := controller.NewCommodityRepository(postgres)
	router.HandleFunc("/api/usertest", mongoUserController.AddUser).Methods("POST")

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
	router.HandleFunc("/test/cookie", mongoUserController.GetSetCoockies).Methods("GET")
	router.HandleFunc("/test/{userID}&{itemID}", helper.Testing).Methods("GET")
	router.HandleFunc("/test", postgreCommodityComtroller.GetAllCommodities).Methods("GET")
	router.HandleFunc("/test{id}&{price}", postgreCommodityComtroller.GetAllCommodities).Methods("GET")

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

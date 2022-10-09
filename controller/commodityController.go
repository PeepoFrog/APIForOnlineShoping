package controller

import (
	"encoding/json"
	"internetshop/helper"
	"internetshop/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommodityController struct {
	repository helper.CommodityRepository
}

func NewCommodityRepository(repository helper.CommodityRepository) *CommodityController {
	return &CommodityController{repository: repository}
}
func (c *CommodityController) CreateCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var commodity model.Commodity
	_ = json.NewDecoder(r.Body).Decode(&commodity) //check what instade of this var
	c.repository.InsertComodity(commodity)
	json.NewEncoder(w).Encode(&commodity)
}
func (c *CommodityController) SetPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	newP, err := strconv.ParseFloat(params["price"], 64)
	if err != nil {
		log.Panic(err)
	}
	c.repository.SetPrice(params["id"], newP)
	json.NewEncoder(w).Encode(params["id"])

}
func (c *CommodityController) SetQuantity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	newP, err := strconv.Atoi(params["price"])
	if err != nil {
		log.Panic(err)
	}
	c.repository.SetQuantity(params["id"], newP)
	json.NewEncoder(w).Encode(params["id"] + params["price"])

}

func (c *CommodityController) DeleteOneCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	c.repository.DeleteOneCommodity(params["id"])
	json.NewEncoder(w).Encode("Commodity with id :" + params["id"] + " was deleted")

}
func (c *CommodityController) DeleteALlCommodities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := c.repository.DeleteALlCommodities()
	json.NewEncoder(w).Encode(strconv.Itoa(int(count)) + "goods was deleted")

}

func (c *CommodityController) GetAllCommodities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := c.repository.GetAllCommodities()
	json.NewEncoder(w).Encode(allMovies)

}
func (c *CommodityController) GetOneCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	searched, _ := c.repository.GetOneCommodity(params["id"])
	json.NewEncoder(w).Encode(searched)

}

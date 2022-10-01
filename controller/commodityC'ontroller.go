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

func init() {

}
func CreateCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var commodity model.Commodity
	_ = json.NewDecoder(r.Body).Decode(&commodity) //check what instade of this var
	helper.InsertComodity(commodity)
	json.NewEncoder(w).Encode(&commodity)
}
func SetPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	newP, err := strconv.ParseFloat(params["price"], 64)
	if err != nil {
		log.Panic(err)
	}
	helper.SetPrice(params["id"], newP)
	json.NewEncoder(w).Encode(params["id"])

}
func SetQuantity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	newP, err := strconv.Atoi(params["price"])
	if err != nil {
		log.Panic(err)
	}
	helper.SetQuantity(params["id"], newP)
	json.NewEncoder(w).Encode(params["id"] + params["price"])

}

func DeleteOneCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	helper.DeleteOneCommodity(params["id"])
	json.NewEncoder(w).Encode("Commodity with id :" + params["id"] + " was deleted")

}
func DeleteALlCommodities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := helper.DeleteALlCommodities()
	json.NewEncoder(w).Encode(strconv.Itoa(int(count)) + "goods was deleted")

}

func GetAllCommodities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := helper.GetAllCommodities()
	json.NewEncoder(w).Encode(allMovies)

}
func GetOneCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	searched, _ := helper.GetOneCommodity(params["id"])
	json.NewEncoder(w).Encode(searched)

}

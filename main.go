package main

import (
	"fmt"
	"internetshop/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("****SERVER-----MONGO----DB-----STARTing****")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("****LISTENING AP PORT 4000****")

}

//sending in this format
// {
// 	"movie": "iron man 3",
// 	"watched": false
//   }

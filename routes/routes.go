package routes

import (
	"log"
	"net/http"

	"golang-crud-sql/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	myRouter.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	myRouter.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	myRouter.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
	//log.Fatal(http.ListenAndServe(config.AppConfig.Port, myRouter))
}

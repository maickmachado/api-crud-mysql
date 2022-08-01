package controllers

import (
	"encoding/json"
	"fmt"
	"golang-crud-sql/database"
	"golang-crud-sql/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("respondendo")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("func CreateProduct")
	// Defines a new product variable.
	var product entities.Product
	//Decodes the Body of the Incoming JSON request and maps it to the newly created product variable.
	json.NewDecoder(r.Body).Decode(&product)
	//Using GORM, we try to create a new product by passing in the parsed product.
	//This would ideally create a new record in the products table for us.
	database.Instance.Create(&product)
	//Returns the newly created product data back to the client.
	json.NewEncoder(w).Encode(product)
}

//takes in the productId and queries against the database if the ID is found in the table
func checkIfProductExists(productId string) bool {
	var product entities.Product
	//acha o primeiro item que tem match com as condições especificadass
	//passar como parâmetro o endereço de memória da variável e segundo parametro o ID
	database.Instance.First(&product, productId)
	//If there are no records found for the ID, GORM would return the ID as 0
	return product.ID != 0
	// if product.ID == 0 {
	// 	return false
	// }
	// return true
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	//Gets the Product Id from the Query string of the request. To be clear, the client would have to send a GET
	//request to the host/api/products/{id} to get the details of a particular product. This means that the product
	//id will be passed as a part of the Request URL. MUX makes it simple for us to extract this passed product id
	//from the request URL.
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	//With the help of GORM, the product table is queried with the product Id. This would fill in the product details
	//to the newly created product variable.
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	//we encode the product variable and send it back to the client.
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chega ate aqui")
	//Here we define an empty new list of products.
	var products []entities.Product
	//Maps all the available products into the product list variable.
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Finally, it simply encodes the products variable and returns it back to the client.
	//Another thing to note here is that we are also writing an HTTP Status code of 200 OK to the Header of the response.
	//You may also use this for other handlers if you like.
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//Here too, MUX extracts the id from the URL and assigns the value to the id variable. Then the code checks if the passed
	//product Id actually exists in the product table.
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	//If found, GORM queries the product record to the product variable. The JSON decoder then converts the request body
	//to a product variable, which is then saved to the database table.
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Extracts the id to be deleted from the request URL. Checks if the ID is actually available in the product table.
	//Then we create a new product variable.
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	//GORM deletes the product by ID.
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

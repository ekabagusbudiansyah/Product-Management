package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type Database struct {
	Products []Product
}

var (
	data = make(map[int]Product)
)

func SetJSONResp(w http.ResponseWriter, message []byte, httpCode int) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(httpCode)
	w.Write(message)
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		message := []byte(`{"Invalid http method"}`)
		SetJSONResp(w, message, http.StatusMethodNotAllowed)
		return
	}
	var product = Database{Products: []Product{}}
	for _, value := range data {
		product.Products = append(product.Products, value)

	}
	productJSON, err := json.Marshal(&product)
	if err != nil {
		message := []byte(`{"Message": "Error when parsing data"}`)
		SetJSONResp(w, message, http.StatusInternalServerError)
		return
	}
	SetJSONResp(w, productJSON, http.StatusOK)
}

func addProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		message := []byte(`{"Message": "Invalid http method"}`)
		SetJSONResp(w, message, http.StatusMethodNotAllowed)
		return
	}
	var product Product
	payload := r.Body
	defer r.Body.Close()
	err := json.NewDecoder(payload).Decode(&product)
	if err != nil {
		massage := []byte(`{"Message": "Invalid when parsing data"}`)
		SetJSONResp(w, massage, http.StatusInternalServerError)
		return
	}
	data[product.ID] = product
	message := []byte(`{"Message": "Succes add product"}`)
	SetJSONResp(w, message, http.StatusOK)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		message := []byte(`{"Message": "Invalid http method"}`)
		SetJSONResp(w, message, http.StatusMethodNotAllowed)
		return
	}
	if _, ok := r.URL.Query()["id"]; !ok {
		message := []byte(`{"Message": "Required product id"}`)
		SetJSONResp(w, message, http.StatusBadRequest)
		return
	}
	id := r.URL.Query()["id"][0]
	idInt, _ := strconv.Atoi(id)
	product, ok := data[idInt]
	if !ok {
		message := []byte(`{"Message": "Product not found"}`)
		SetJSONResp(w, message, http.StatusNotFound)
		return
	}
	delete(data, idInt)
	productJSON, err := json.Marshal(&product)
	if err != nil {
		message := []byte(`{"Message": "Error when parsing data"}`)
		SetJSONResp(w, message, http.StatusInternalServerError)
		return
	}
	message := []byte(`{"Message": "Succes delete product"}`)
	SetJSONResp(w, message, http.StatusOK)
	SetJSONResp(w, productJSON, http.StatusOK)
}

func main() {
	data[1] = Product{ID: 1, Name: "Book", Price: 5000}
	data[2] = Product{ID: 2, Name: "Bolpoint", Price: 3000}

	http.HandleFunc("/get-products", getProducts)
	http.HandleFunc("/add-product", addProduct)
	http.HandleFunc("/delete-product", deleteProduct)

	fmt.Println("Server run in localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

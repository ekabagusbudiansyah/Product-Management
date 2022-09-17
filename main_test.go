package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestGetProducts(t *testing.T) {
	data[1] = Product{ID: 1, Name: "Book", Price: 5000}

	req, err := http.NewRequest("GET", "/get-products", nil)
	if err != nil {
		t.Fatal(err) 
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getProducts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Products":[{"id":1,"name":"Book","price":5000}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAddProduct(t *testing.T) {

	testBody := `{"id": 1, "name": "Book", "price": 4000}`

	req, err := http.NewRequest("POST", "/add-product", bytes.NewBufferString(testBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Message": "Succes add product"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteProduct(t *testing.T) {

	data[1] = Product{ID: 1, Name: "Book", Price: 5000}

	req, err := http.NewRequest("DELETE", "/delete-product", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Message": "Succes delete product"}{"id":1,"name":"Book","price":5000}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

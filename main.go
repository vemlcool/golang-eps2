package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/TMaulana26/go-ep-2/app"
)

var products []app.Product

func main() {
	http.HandleFunc("/api/v1/products.json", productJSONHandler)
	http.HandleFunc("/api/v1/products.xml", productXMLHandler)
	http.HandleFunc("/api/v1/add-products.json", addProductsJSONHandler)
	http.ListenAndServe(":8081", nil)
}

func productJSONHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		{
			Base: app.Base{
				ID:   1,
				Name: "Lakungan",
			},
			Price: 5000000000,
			Category: app.Category{
				Base: app.Base{ID: 2, Name: "Ecommerce"},
			},
		},
		{
			Base: app.Base{
				ID:   2,
				Name: "My Score Card",
			},
			Price: 600000000000,
			Category: app.Category{
				Base: app.Base{ID: 1, Name: "Project Management"},
			},
		},
	}

	// products[0].SetIDAndName(1, "Lakungan")
	// products[1].SetIDAndName(2, "My Score Card")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

func productXMLHandler(w http.ResponseWriter, r *http.Request) {
	products := []app.Product{
		{
			Base: app.Base{
				ID:   1,
				Name: "Lakungan",
			},
			Price: 5000000000,
			Category: app.Category{
				Base: app.Base{ID: 2, Name: "Ecommerce"},
			},
		},
		{
			Base: app.Base{
				ID:   2,
				Name: "My Score Card",
			},
			Price: 600000000000,
			Category: app.Category{
				Base: app.Base{ID: 1, Name: "Project Management"},
			},
		},
	}

	products[0].SetIDAndName(1, "Lakungan")
	products[1].SetIDAndName(2, "My Score Card")

	w.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(products)
}

func addProductsJSONHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var newProducts app.Product
	err = json.Unmarshal(body, &newProducts)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	products = append(products, newProducts)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newProducts)

}

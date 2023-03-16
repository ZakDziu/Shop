package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shop/internal/core/services/product"
)

func GetAllProducts(
	productService *product.ProductService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		products, err := productService.GetAll(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&requestErrorFmt{Error: err.Error()})

			return
		}

		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			log.Panic(err)
		}
	}
}

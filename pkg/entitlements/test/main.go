package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v81"

	stripeclient "github.com/theopenlane/core/pkg/entitlements/stripe"
)

func main() {
	stripeClient := stripeclient.NewStripeClient()

	http.HandleFunc("/create-customer", func(w http.ResponseWriter, r *http.Request) {
		var params stripe.CustomerParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		customer, err := stripeClient.CreateCustomer(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(customer)
	})

	http.HandleFunc("/update-customer", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			ID     string                `json:"id"`
			Params stripe.CustomerParams `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		customer, err := stripeClient.UpdateCustomer(params.ID, &params.Params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(customer)
	})

	// Add similar handlers for products, prices, and plans

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}

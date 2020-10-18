package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/prakashsingha/orderAPI/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder creates a new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var order services.Order
	json.NewDecoder(r.Body).Decode(&order)

	// TODO: Validate model
	payment := services.Payment{
		Status: services.PaymentPending,
	}
	paymentIDHex, err := services.CreatePayment(payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Payment creation failed. Error: "` + err.Error() + ` }`))
		return
	}

	oID, HexErr := primitive.ObjectIDFromHex(paymentIDHex)
	if HexErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Invalid paymentID. Error: "` + err.Error() + ` }`))
		return
	}

	order.PaymentID = oID
	newOrder, err := services.CreateOrder(order)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"message": "Order creation failed. Error: "` + err.Error() + ` }`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newOrder)
}

// GetOrders returns a list of orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	urlQuery := r.URL.Query()
	searchParam := services.SearchParam{
		HotelName:     urlQuery.Get("hotelName"),
		CustomerName:  urlQuery.Get("cname"),
		CustomerEmail: urlQuery.Get("cemail"),
		CustomerPhone: urlQuery.Get("cphone"),
	}

	orders, err := services.GetOrders(searchParam)
	returnErrorResponse(w, r, err, http.StatusInternalServerError)

	if orders == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "No any order found"}`))
		return
	}

	json.NewEncoder(w).Encode(orders)
}

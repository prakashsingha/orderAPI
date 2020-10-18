package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/prakashsingha/orderAPI/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentVM struct {
	ID            primitive.ObjectID `json:"id,omitempty"`
	Status        int                `json:"status"`
	Code          string             `json:"code,omitempty"`
	ConfirmedDate *time.Time         `json:"confirmedDate,omitempty"`
	CardNo        string             `json:"cardNo,omitempty"`
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	_id := r.URL.Query().Get("id")
	_field := r.URL.Query().Get("field")
	payment, err := services.GetPayment(_id)
	returnErrorResponse(w, r, err, http.StatusInternalServerError)

	vm := PaymentVM{}
	if _field == "status" {
		vm.ID = payment.ID
		vm.Status = payment.Status
	} else {
		vm.ID = payment.ID
		vm.Status = payment.Status
		vm.ConfirmedDate = payment.ConfirmedDate
		vm.CardNo = payment.CardNo
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vm)
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var vm PaymentVM
	json.NewDecoder(r.Body).Decode(&vm)

	payment := services.Payment{
		ID:     vm.ID,
		Status: vm.Status,
		CardNo: vm.CardNo,
	}
	result, err := services.UpdatePayment(payment)
	returnErrorResponse(w, r, err, http.StatusInternalServerError)

	w.WriteHeader(http.StatusOK)
	s := fmt.Sprintf("{\"modifiedCount\": %v}", result)
	w.Write([]byte(s))
}

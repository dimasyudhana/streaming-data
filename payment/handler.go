package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router         *mux.Router
	PaymentService *PaymentService
}

func NewServer(paymentService *PaymentService) *Server {
	s := &Server{
		Router:         mux.NewRouter(),
		PaymentService: paymentService,
	}
	s.routes()

	return s
}

func (s Server) MakePayment() http.HandlerFunc {

	type Response struct {
		PaymentID string `json:"payment_id"`
		OrderID   string `json:"order_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var command NewPaymentCommand
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&command)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err)
			return
		}

		paymentID, err := s.PaymentService.MakePayment(command)
		if err != nil {
			log.Println(err)
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, Response{PaymentID: paymentID, OrderID: command.OrderID}, nil)
	}
}

func (s Server) GetPayments() http.HandlerFunc {

	type Payment struct {
		PaymentID string  `json:"payment_id"`
		OrderID   string  `json:"order_id"`
		Value     float64 `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		payments, err := s.PaymentService.GetPayment()
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		resp := make([]Payment, 0)
		for _, p := range payments {
			resp = append(resp, Payment{
				PaymentID: p.ID,
				OrderID:   p.OrderID,
				Value:     p.Value,
			})
		}

		WriteSuccessResponse(w, http.StatusOK, resp, nil)
	}

}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}, headMap map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	if headMap != nil && len(headMap) > 0 {
		for key, val := range headMap {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func WriteFailResponse(w http.ResponseWriter, statusCode int, error interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(error)
	w.Write(jsonData)
}

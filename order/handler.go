package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router       *mux.Router
	OrderService *OrderService
}

func NewServer(orderService *OrderService) *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		OrderService: orderService,
	}
	s.routes()

	return s
}

func (s *Server) CreateOrder() http.HandlerFunc {

	type Response struct {
		ID string `json:"order_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var command NewOrderCommand
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&command)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err)
			return
		}

		id, err := s.OrderService.NewOrder(command)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		log.Println(id)

		WriteSuccessResponse(w, http.StatusOK, &Response{ID: id}, nil)
	}
}

func (s *Server) GetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := s.OrderService.GetOrders()
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, nil)
			return
		}

		resp := make([]Order, 0)
		for _, o := range orders {
			resp = append(resp, o)
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

package main

import (
	"context"
	"log"
	"net/http"

	commonlib "stream/commonlib"
)

func Run() {
	conf, err := commonlib.Load(DefaultConfig, "")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	em, err := commonlib.New(context.Background(), conf.Get("event"))
	if err != nil {
		log.Fatal("Failed to initialize event emitter:", err)
	}

	orderService := NewOrderService(em)
	orderHandler := NewServer(orderService)

	commonlib.ConsoleGreet(conf.GetString("name"), conf.GetString("version"), "", conf.GetInt("port"))
	log.Fatal(http.ListenAndServe(":50050", orderHandler.Router))
}

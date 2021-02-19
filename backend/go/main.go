package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"video-crm/request"
)

var (
	ServerAddr = os.Getenv("VC_addr")
)

const (
	Na     = iota
	Active = iota
)

//func
type Device struct {
	ID         int    `json:"id"`
	Hardware   string `json:"hardware"`
	Owner      string `json:"owner"`
	Status     int    `json:"status"`
	CreateDate int    `json:"create_date"`
	Uptime     int    `json:"uptime"`
	UpdateTime int    `json:"update_time"`
	Info       string `json:"info"`
	Token      string `json:"token"`
}

func NewServer(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:         ServerAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
		TLSConfig:    nil,
		Handler:      mux,
	}
}

func main() {
	logger := log.New(os.Stdout, "vcrm > ", log.LstdFlags|log.Lshortfile)

	mux := http.NewServeMux()
	h := request.NewHandlers(logger)
	h.SetupRequest(mux, "/add_device", h.AddDevice)
	h.SetupRequest(mux, "/info", h.Info)

	logger.Println("Starting...")
	srv := NewServer(mux)

	log.Fatal(srv.ListenAndServe())
}

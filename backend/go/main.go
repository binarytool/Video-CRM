package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"video-crm/request"
)

var (
	ServerAddr     = os.Getenv("VC_addr")
	ServerUserName = os.Getenv("VC_uname")
	ServerPassword = os.Getenv("VC_psw")
)

const (
	Na     = iota
	Active = iota
)

//func

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

	h := request.NewHandlers(logger, ServerUserName, ServerPassword)
	h.InitDB()
	//defer h.DB.Close()

	h.SetupRequest(mux, "/info", h.Info)
	h.SetupRequest(mux, "/device", h.Device)

	logger.Println("Starting...")
	srv := NewServer(mux)

	log.Fatal(srv.ListenAndServe())
}

package request

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) ProfileRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		h.logger.Println("Request processed in %s", time.Now().Sub(start))
	}
}

func (h *Handlers) SetupRequest(mux *http.ServeMux, path string, fn func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, h.ProfileRequest(fn))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

// Handler funcs

func (h *Handlers) AddDevice(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Request add_device")

	switch r.Method {
	case http.MethodPut:
		var ks []string
		var vs []string
		for k, v := range r.URL.Query() {
			ks = append(ks, k)
			vs = append(vs, v[0])
		}

		req_s := fmt.Sprintf("INSERT INTO device (%s) values ('%s')", strings.Join(ks, ","), strings.Join(vs, "','"))
		h.logger.Println(req_s)
		db, err := sql.Open("mysql", "root:@/video-crm")
		insert, err := db.Query(req_s)
		defer insert.Close()
		defer db.Close()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", req_s)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("done"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) Info(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Request processed")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Info"))
}

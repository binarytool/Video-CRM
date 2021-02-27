package request

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RequestHandler struct {
	Logger   *log.Logger
	DB       *gorm.DB
	Username string
	Password string
}

type Device struct {
	ID        int       `json:"id"`
	Hardware  string    `json:"hardware"`
	Owner     string    `json:"owner"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"create_date"`
	Uptime    int       `json:"uptime"`
	UpdatedAt time.Time `json:"update_time"`
	Info      string    `json:"info"`
	Token     string    `json:"token"`
}

type Content struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	caption   string `json:"caption"`
	Size      int    `json:"status"`
	Info      string `json:"info"`
	mediainfo string `json:"mediainfo"`
}

type Stats struct {
	DeviceId  int
	CreatedAt time.Time
	Event     string
	Value     string
}

func (h *RequestHandler) ProfileRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		h.Logger.Println("Request processed in %s", time.Now().Sub(start))
	}
}

func (h *RequestHandler) InitDB() {
	h.Logger.Println("Try setup SQL connection")
	dsn := h.Username + ":" + h.Password + "@tcp(127.0.0.1:3306)/video-crm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db, err := sql.Open("mysql", h.Username+":"+h.Password+"@/video-crm")
	h.DB = db
	if err != nil {
		panic(err)
	}
	h.Logger.Println("Connection established")
}

func (h *RequestHandler) SetupRequest(mux *http.ServeMux, path string, fn func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, h.ProfileRequest(fn))
}

func NewHandlers(logger *log.Logger, uname string, psw string) *RequestHandler {
	return &RequestHandler{
		Logger:   logger,
		DB:       nil,
		Username: uname,
		Password: psw,
	}
}

// Handler funcs

func (h *RequestHandler) Device(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println("Request add_device")

	switch r.Method {
	case http.MethodPut:
		devicestr := r.URL.Query()["device"]
		var device Device
		bytes := []byte(devicestr[0])
		err := json.Unmarshal(bytes, &device)
		if err != nil {
			panic(err)
		}

		result := h.DB.Create(&device)
		if result.Error != nil {
			panic(result.Error)
		}

		h.Logger.Printf("Added device with id = %d\n", device.ID)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("done"))
	case http.MethodGet:
		var devices []Device
		result := h.DB.Find(&devices)

		if result.Error != nil {
			panic(result.Error)
		}

		s, err := json.Marshal(devices)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(s)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RequestHandler) Init(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println("TODO: populate db schema")

	//h.DB.AutoMigrate(&Device{}, &Content{}, Stats{})
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Not implemented"))
}

func (h *RequestHandler) Info(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println("Request processed")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Info"))
}

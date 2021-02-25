package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
	"video-crm/request"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	timestampf = "2006-01-02 15:04:05"
)

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestDevice(t *testing.T) {
	tt := []struct {
		name       string
		api        string
		method     string
		input      *request.Device
		want       string
		statusCode int
	}{
		{
			name:       "create schema",
			api:        "/init",
			method:     http.MethodPut,
			input:      nil,
			want:       "ok",
			statusCode: http.StatusOK,
		},
		{
			name:   "add device",
			api:    "/device",
			method: http.MethodPut,
			input: &request.Device{
				Hardware: "thw 1",
				Owner:    "towner 1",
				Status:   Active,
				Info:     "tinfo 1",
				Token:    RandStringRunes(16),
			},
			want:       "done",
			statusCode: http.StatusOK,
		},
		{
			name:   "add device",
			api:    "/device",
			method: http.MethodPut,
			input: &request.Device{
				Hardware: "thw 2",
				Owner:    "towner 2",
				Status:   Active,
				Info:     "tinfo 2",
				Token:    RandStringRunes(16),
			},
			want:       "done",
			statusCode: http.StatusOK,
		},
		{
			name:   "add device",
			api:    "/device",
			method: http.MethodPut,
			input: &request.Device{
				Hardware: "thw 3",
				Owner:    "towner 3",
				Status:   Active,
				Info:     "tinfo 3",
				Token:    RandStringRunes(16),
			},
			want:       "done",
			statusCode: http.StatusOK,
		},
		{
			name:       "get device",
			api:        "/device",
			method:     http.MethodGet,
			input:      nil,
			want:       "done",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		logger := log.New(os.Stdout, "vcrm > ", log.LstdFlags|log.Lshortfile)
		h := request.NewHandlers(logger, ServerUserName, ServerPassword)
		h.InitDB()
		//defer h.DB.Close()

		t.Run(tc.name, func(t *testing.T) {
			if tc.api == "/init" {
				fmt.Println("Init part")
				//r := httptest.NewRequest(tc.method, "/init", nil)
				//q := r.URL.Query()
			}
			if tc.api == "/device" {
				r := httptest.NewRequest(tc.method, "/device", nil)
				q := r.URL.Query()

				if tc.method == http.MethodPut {
					device := &request.Device{
						Hardware: tc.input.Hardware,
						Owner:    tc.input.Owner,
						Status:   1,
						Uptime:   0,
						Info:     tc.input.Info,
						Token:    tc.input.Token,
					}

					s, err := json.Marshal(device)
					if err != nil {
						fmt.Println(err)
						return
					}

					q.Add("device", string(s))

					r.URL.RawQuery = q.Encode()
					responseRecorder := httptest.NewRecorder()
					h.Device(responseRecorder, r)

					if responseRecorder.Code != tc.statusCode {
						t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
					}

					if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
						t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
					}
				} else if tc.method == http.MethodGet {
					r.URL.RawQuery = q.Encode()
					responseRecorder := httptest.NewRecorder()
					h.Device(responseRecorder, r)
					if responseRecorder.Code != tc.statusCode {
						t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
					}
					if len(strings.TrimSpace((responseRecorder.Body.String()))) == 0 {
						t.Errorf("Want non empty responce")
					}
				}
			}
		})
	}
}

func setup() {
	h := request.NewHandlers(log.New(os.Stdout, "vcrm > ", log.LstdFlags|log.Lshortfile), ServerUserName, ServerPassword)
	h.InitDB()
	r := httptest.NewRequest(http.MethodPut, "/init", nil)
	responseRecorder := httptest.NewRecorder()
	h.Init(responseRecorder, r)
}

func shutdown() {

}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

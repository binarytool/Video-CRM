package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"video-crm/request"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestAddDevice(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		input      *Device
		want       string
		statusCode int
	}{
		{
			name:   "add device",
			method: http.MethodPut,
			input: &Device{
				ID:         0,
				Hardware:   "test hw",
				Owner:      "test owner",
				Status:     Active,
				CreateDate: 0,
				Uptime:     0,
				UpdateTime: 0,
				Info:       "test info",
				Token:      RandStringRunes(16),
			},
			want:       "done",
			statusCode: http.StatusOK,
		},
		{
			name:       "get device",
			method:     http.MethodGet,
			input:      nil,
			want:       "done",
			statusCode: http.StatusOK,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.method == http.MethodPut {
				r := httptest.NewRequest(tc.method, "/add_device", nil)
				q := r.URL.Query()
				q.Add("hardware", tc.input.Hardware)
				q.Add("owner", tc.input.Owner)
				q.Add("status", strconv.Itoa(1))
				q.Add("create_date", strconv.Itoa(tc.input.CreateDate))
				q.Add("uptime", strconv.Itoa(tc.input.Uptime))
				q.Add("update_time", strconv.Itoa(tc.input.UpdateTime))
				q.Add("info", tc.input.Info)
				q.Add("token", tc.input.Token)
				r.URL.RawQuery = q.Encode()

				responseRecorder := httptest.NewRecorder()
				logger := log.New(os.Stdout, "vcrm > ", log.LstdFlags|log.Lshortfile)
				h := request.NewHandlers(logger)
				h.AddDevice(responseRecorder, r)

				if responseRecorder.Code != tc.statusCode {
					t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
				}

				if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
					t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
				}
			}
		})
	}
}

func TestThreeElements(t *testing.T) {
	//t.Error("No")
}

package server

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type DecodedRequest struct {
	Str string `json:"inputString"`
}

type DecodedResponse struct {
	Str string `json:"outputString"`
}

func APIVersionGetRequest(w http.ResponseWriter, _ *http.Request) {
	ans := []byte("v1.0.0")
	_, err := w.Write(ans)
	if err != nil {
		return
	}
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	var req DecodedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatal("Invalid request")
		return
	}

	decodedStr, err := base64.StdEncoding.DecodeString(req.Str)
	if err != nil {
		log.Fatal("Failed to encode")
		return
	}
	response := DecodedResponse{string(decodedStr)}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("Failed to encode")
		return
	}
}

func HardOpGetRequest(w http.ResponseWriter, r *http.Request) {
	Time := time.Duration(10+rand.Intn(11)) * time.Second
	time.Sleep(Time)
	if rand.Intn(2) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

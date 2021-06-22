package controllers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, code int, payload interface{}) {
	out, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

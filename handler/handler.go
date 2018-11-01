package handler

import (
	"fmt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "200", "message": "OK"}`)
}

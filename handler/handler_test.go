package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(RegisterHandler))
	req := httptest.NewRequest("GET", "/register", nil)
	rec := httptest.NewRecorder()

	log.Printf("%v", rec.Code)
	log.Printf("%v", req)

}

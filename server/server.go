package server

import (
	"log"
	"net/http"

	"github.com/Nobv/webauthn/color"
	"github.com/Nobv/webauthn/handler"
	"github.com/gorilla/mux"
)

func NewServer(addr string) *http.Server {

	router := InitRouting()

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func StartServer(server *http.Server) {
	log.Println(color.Coloring(color.Cyan, "starting server..."))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}

func InitRouting() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/register", handler.RegisterHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}

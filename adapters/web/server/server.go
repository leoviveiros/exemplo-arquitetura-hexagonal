package server

import (
	"go-hexagonal/adapters/web/handler"
	"go-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (s *WebServer) Start() {
	router := mux.NewRouter()
	middleware := negroni.New(negroni.NewLogger())

	handler.NewProductHandler(router, middleware, s.Service)

	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 	10 * time.Second,
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "LOG:", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
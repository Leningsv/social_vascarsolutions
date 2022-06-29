package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"social_vascarsolutions/internal/entities"
	"social_vascarsolutions/internal/handlers"
)

func main() {
	entities.MigrateEntities()
	mux := mux.NewRouter()
	mux.HandleFunc("/api/report/post", handlers.ReportPost).Methods(http.MethodPost)
	mux.HandleFunc("/api/report/comment", handlers.ReportComment).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":3001", mux))
}

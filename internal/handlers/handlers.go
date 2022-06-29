package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_vascarsolutions/internal/models"
	"social_vascarsolutions/internal/services"
)

func ReportPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var post models.PostReport
	if err := decoder.Decode(&post); err == nil {
		reportService := services.ReportService{}
		report := reportService.CreatePost(post)
		Send(w, report)
	} else {
		SendBadRequest(w)
	}
}

func ReportComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var comment models.CommentReport
	if err := decoder.Decode(&comment); err == nil {
		reportService := services.ReportService{}
		report := reportService.CreateComment(comment)
		Send(w, report)
	} else {
		SendBadRequest(w)
	}
}

func Send(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	output, _ := json.Marshal(&data)
	fmt.Fprintf(w, string(output))
}

func SendBadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	output, _ := json.Marshal(map[string]string{
		"message": "bad request",
	})
	fmt.Fprintf(w, string(output))
}

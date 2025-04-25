package handlers

import (
	"encoding/json"
	"mini_BA/db"
	"net/http"
)

type GetReportsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetReportsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GetReportsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	reports, err := db.GetReportsByEmailAndPassword(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Failed to retrieve reports", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"email":   req.Email,
		"reports": reports,
	})
}

package apistream

import (
	"encoding/json"
	"net/http"

	"SLtwitchApi/db"
	"SLtwitchApi/models"
	"SLtwitchApi/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type StreamerInput struct {
	Username   string `json:"username"`
	Fraktion   string `json:"fraktion"`
	CharName   string `json:"charname"`
	ShortStory string `json:"shortStory"`
}

func AddStreamHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var input StreamerInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate username length
	if len(input.Username) > 25 {
		http.Error(w, "Username is too long (max 25)", http.StatusBadRequest)
		return
	}

	// Validate ShortStory length
	if len(input.ShortStory) > 150 {
		http.Error(w, "ShortStory is too long (max 150)", http.StatusBadRequest)
		return
	}

	// Check if username already exists
	var count int64
	db.SQL.Model(&models.Streamer{}).Where("user_name = ?", input.Username).Count(&count)
	if count > 0 {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// Set default values
	stream := models.Streamer{
		Username:    input.Username,
		Fraktion:    input.Fraktion,
		Charname:    input.CharName,
		ShortStory:  input.ShortStory,
		Active:      false,
		DataFetched: false,
		ID:          uuid.New(),
	}

	// Save stream to database
	if err := db.SQL.Create(&stream).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusOK) // Set HTTP status code to 200 OK
	json.NewEncoder(w).Encode(map[string]bool{
		"success": true,
	})
}

func DeleteStreamHandler(w http.ResponseWriter, r *http.Request) {
	// Get UserID from URL parameter
	userID, ok := mux.Vars(r)["UserID"]
	if !ok {
		http.Error(w, "UserID not found in request", http.StatusBadRequest)
		return
	}
	// Check if stream exists
	var stream models.Streamer
	if err := db.SQL.Where("id = ?", userID).First(&stream).Error; err != nil {
		http.Error(w, "Stream not found", http.StatusBadRequest)
		return
	}

	// Delete stream from database
	if err := db.SQL.Delete(&stream).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusOK) // Set HTTP status code to 200 OK
	json.NewEncoder(w).Encode(map[string]bool{
		"success": true,
	})
}

func ListAllActiveStreamers(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(services.GetAllStreamerInformation())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

package api

import (
	"encoding/json"
	"net/http"

	"github.com/buoto/civ-6-bot/turn"
)

// ListAllHandler lists all turn notifications and returns them as csv
func ListAllHandler(w http.ResponseWriter, r *http.Request) {
  repo := turn.NewTurnNotificationRepository()
  notifications, err := repo.GetAll()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  w.Header().Set("Content-Type", "text/csv")
  w.WriteHeader(http.StatusOK)
  encoder := json.NewEncoder(w)
  encoder.Encode(notifications)
}

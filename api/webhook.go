package api

import (
	"encoding/json"
	"net/http"

	"github.com/buoto/civ-6-bot/turn"
)

// WebhookHandler handles incoming webhooks and stores them in db
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var notification turn.Notification
	if err := decoder.Decode(&notification); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := turn.NewTurnNotificationRepository()
	if err := repo.Create(notification.GameName, notification.PlayerName, notification.TurnNumber); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

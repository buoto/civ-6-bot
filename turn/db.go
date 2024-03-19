package turn

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// NotificationRepository is a repository for turn notifications
type NotificationRepository struct {
	db *sql.DB
}

// NewTurnNotificationRepository creates a new TurnNotificationRepository
func NewTurnNotificationRepository() *NotificationRepository {
	return &NotificationRepository{getDB()}
}

// Create creates a new turn notification
func (r *NotificationRepository) Create(gameName string, playerName string, turnNumber int) error {
	_, err := r.db.Exec("INSERT INTO turn_notifications (game_name, player_name, turn_number) VALUES ($1, $2, $3)", gameName, playerName, turnNumber)
	return err
}

// GetAll gets all turn notifications
func (r *NotificationRepository) GetAll() ([]Notification, error) {
	rows, err := r.db.Query("SELECT game_name, player_name, turn_number FROM turn_notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := []Notification{}
	for rows.Next() {
		var notification Notification
		err := rows.Scan(&notification.GameName, &notification.PlayerName, &notification.TurnNumber)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

// get pq client for database set up in vercel
func getDB() *sql.DB {
	// get the database url from the environment
	url := os.Getenv("DATABASE_URL")
	// open a connection to the database
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

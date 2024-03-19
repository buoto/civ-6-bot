package turn

// Notification is a struct that holds a turn notification
type Notification struct {
	GameName   string `json:"value1"`
	PlayerName string `json:"value2"`
	TurnNumber int    `json:"value3"`
}

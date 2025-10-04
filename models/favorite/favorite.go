package favorite_model

type Favorite struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Asset  string `json:"asset"` // e.g. "AAPL", "BTC", etc.
}

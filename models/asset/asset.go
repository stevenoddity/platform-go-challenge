package asset_model

type Asset struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"` // e.g. "AAPL", "BTC", etc.
}

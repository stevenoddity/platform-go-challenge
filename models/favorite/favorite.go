package favorite_model

// Favorite represents a user's favorite asset in the system.
type Favorite struct {
	ID     int    `json:"id"`      // Unique identifier for the favorite
	UserID int    `json:"user_id"` // Identifier for the user who owns the asset
	Asset  string `json:"asset"`   // Identifier for the asset who belongs to the user
}

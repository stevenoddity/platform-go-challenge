package utils

import (
	favorite_model "gwi/models/favorite"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateFavorite(fav *favorite_model.Favorite) error {
	return validate.Struct(fav)
}

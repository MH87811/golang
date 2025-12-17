package repositories

import (
	"shop/internal/models"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(ID uint) (models.User, error)
}

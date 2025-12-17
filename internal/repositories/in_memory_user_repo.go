package repositories

import (
	"fmt"
	"shop/internal/models"
)

type InMemoryUserRepo struct {
	users map[string]models.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]models.User),
	}
}

func (r *InMemoryUserRepo) Save(user models.User) (models.User, error) {
	r.users[user.Email] = user
	return user, nil
}

func (r *InMemoryUserRepo) FindByEmail(email string) (models.User, error) {
	user, ok := r.users[email]
	if !ok {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) FindByID(id uint) (models.User, error) {
	user, ok := r.users[string(id)]
	if !ok {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

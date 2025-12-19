package repositories

import (
	"fmt"
	"shop/internal/models"
)

type InMemoryUserRepo struct {
	users      map[uint]models.User
	emailIndex map[string]uint
	autoID     uint
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:      make(map[uint]models.User),
		emailIndex: make(map[string]uint),
		autoID:     1,
	}
}

func (r *InMemoryUserRepo) Save(user models.User) (models.User, error) {
	user.ID = r.autoID
	r.autoID++

	r.users[user.ID] = user
	r.emailIndex[user.Email] = user.ID
	return user, nil
}

func (r *InMemoryUserRepo) FindByEmail(email string) (models.User, error) {
	id, ok := r.emailIndex[email]
	if !ok {
		return models.User{}, fmt.Errorf("user not found")
	}
	return r.users[id], nil
}

func (r *InMemoryUserRepo) FindByID(id uint) (models.User, error) {
	user, ok := r.users[id]
	if !ok {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

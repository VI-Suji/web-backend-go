package managers

import "github.com/VI-Suji/web-backend-go/models"

type UserManager struct {
	// dbClient
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMrg *UserManager) Create(user *models.User) (*models.User, error) {
	return nil, nil
}

package user

import "github.com/MasoudHeydari/golang-crud-rest-api/models"

type UserRepository interface {
	CreateNewUser(newUser *models.User) uint
	GetAllUsers() []*models.User
}

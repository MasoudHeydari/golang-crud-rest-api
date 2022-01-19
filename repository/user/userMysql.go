package user

import (
	"github.com/MasoudHeydari/golang-crud-rest-api/models"
	"github.com/jinzhu/gorm"
)

type mySqlUserRepository struct {
	DbConn *gorm.DB
}

func NewSQLUserRepository(dbConn *gorm.DB) UserRepository {
	return &mySqlUserRepository{
		DbConn: dbConn,
	}
}

func (mysql mySqlUserRepository) CreateNewUser(newUser *models.User) uint {
	// TODO check user email and reject already registered email
	mysql.DbConn.Create(newUser)
	return newUser.ID
}

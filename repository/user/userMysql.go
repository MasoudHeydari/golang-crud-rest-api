package user

import (
	"fmt"
	"github.com/MasoudHeydari/golang-crud-rest-api/models"
	"github.com/jinzhu/gorm"
	"log"
)

type mySqlUserRepository struct {
	DbConn *gorm.DB
}

func NewSQLUserRepository(dbConn *gorm.DB) UserRepository {
	return &mySqlUserRepository{
		DbConn: dbConn,
	}
}

func (mysql *mySqlUserRepository) CreateNewUser(newUser *models.User) uint {
	// TODO check user email and reject already registered email
	mysql.DbConn.Create(newUser)
	return newUser.ID
}

func (mysql *mySqlUserRepository) GetAllUsers() []*models.User {
	fmt.Println("retrieving all users from MySQL")
	var users []*models.User

	// get all the user
	result := mysql.DbConn.Find(&users)
	if result.Error != nil {
		log.Fatalln(result.Error)
		return nil
	}
	fmt.Printf("number of all users: %v", len(users))
	return users
}

func (mysql *mySqlUserRepository) Update(userId uint, userToUpdate *models.User) *models.User {
	savedUser := models.User{}
	mysql.DbConn.First(&savedUser, userId)
	// update all columns that user passed via request
	mysql.DbConn.Model(&savedUser).Updates(userToUpdate)
	return &savedUser
}

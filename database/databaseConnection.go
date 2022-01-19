package database

import (
	"fmt"
	"github.com/MasoudHeydari/golang-crud-rest-api/config"
	"github.com/MasoudHeydari/golang-crud-rest-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Sql *gorm.DB
}

var dbConnection = &DB{}

func ConnectToDB() (*DB, error) {
	sqlCredentials := getMySqlCredentials()
	fmt.Println(sqlCredentials)

	// connect to db
	db, err := gorm.Open("mysql", sqlCredentials)
	//defer db.Close()

	if err != nil {
		fmt.Println("err")
		panic("Failed to connect to the database")
	}
	fmt.Println("mysql db opened successfully")
	dbConnection.Sql = db
	dbConnection.Sql.AutoMigrate(&models.User{})

	return dbConnection, err
}

func getMySqlCredentials() (mySqlCredentials string) {

	appConfig := config.GetAppConfig()
	// Example: user:password@tcp(host:port)/db_name
	mySqlCredentials = fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DB_NAME"],
	)

	return
}

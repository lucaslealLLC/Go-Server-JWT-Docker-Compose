package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	port     string = os.Getenv("PORT")
	username string = os.Getenv("USERNAME")
	password string = os.Getenv("PASSWORD")
	host     string = os.Getenv("HOST")
	database string = os.Getenv("DATABASE")
)

func ConnectDatabase() *gorm.DB {
	var err error
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		var errorMessage string = fmt.Sprintf("Failed to connect to database: %s", err)
		panic(errorMessage)
	}

	return DB
}

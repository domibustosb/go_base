package datalayers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// ConnectDB - Create db connection
func ConnectDB() {
	var dbErr error
	log.Println("Se conectó")
	DBDriver := os.Getenv("DB_DRIVER")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBSSLMode := os.Getenv("DB_SSL_MODE")
	GormDebug := os.Getenv("Gorm_DEBUG")
	GormBlockGlobalUpdate := os.Getenv("GORM_BLOCK_GLOBAL_UPDATE")

	if db, dbErr = gorm.Open(DBDriver,
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			DBHost,
			DBPort,
			DBName,
			DBUser,
			DBPassword,
			DBSSLMode,
		),
	); dbErr != nil {
		panic(dbErr)
	}

	if b, err := strconv.ParseBool(GormDebug); err != nil {
		db.LogMode(b)
	}

	if b, err := strconv.ParseBool(GormBlockGlobalUpdate); err != nil {
		db.BlockGlobalUpdate(b)
	}
}

// CloseConnection - Used to close db connection
func CloseConnection() {
	db.Close()
}

// GetDB - Verify if exist db connection, if not exist create db connection
func GetDB() *gorm.DB {
	if db == nil {
		ConnectDB()
	}
	return db
}

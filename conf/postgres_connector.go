package conf

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DbConfig struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

var Db *gorm.DB

// InitializeDb takes the database connection information as arguments .
// Creates a connection to the db and set the connection to the pointer Db *gorm.DB .
func InitializeDb(dbHost string, username string, dbName string, port int, password string) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=disable", dbHost, username, dbName, port, password) //Build connection string
	if conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
		return
	} else {
		Db = conn
		fmt.Printf("We are connected to database: %v\n", Db)
	}
	return
}

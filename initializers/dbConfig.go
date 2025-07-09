package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

var DB *gorm.DB
var err error

func DatabaseConfig(cfg DBConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode)
	log.Println("dsn:", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database ::: " + err.Error())
	}

	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	log.Fatalf("Failed to connect to the database: %v", err)
	//}
	//defer conn.Close(context.Background())
	//
	//// Example query to test connection
	//var version string
	//if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
	//	log.Fatalf("Query failed: %v", err)
	//}
	//
	//log.Println("Connected to:", version)

}

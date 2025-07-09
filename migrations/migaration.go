package migrations

import (
	"github.com/arbeeorlar/initializers"
	"github.com/arbeeorlar/models"
	"os"
)

func init() {
	initializers.LoadEnvVariable()
	config := initializers.DBConfig{
		Host:     os.Getenv("DB_HOST_2"), //DB_HOST_SUPA_BASE
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"), //_SUPA_BASE
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_INSTANCE"),
		SSLMode:  "require", //"disable", //
	}
	initializers.DatabaseConfig(config)
}

func AutoMigrate() {
	db := initializers.DB
	db.AutoMigrate(
		&models.Faculty{},
		&models.Department{},
		&models.Level{},
		&models.Grade{},
		&models.Student{},
		&models.Course{},
		&models.Score{},
	)
	initializers.SeedAll(db)
}

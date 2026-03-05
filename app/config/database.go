package config

import (
	"fiber-sample-project/app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, name, port, sslMode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Database doesn't exist, create it
		log.Println("⚠️  Creating database...")
		if err := createDatabase(user, pass, host, port, sslMode, name); err != nil {
			panic(fmt.Sprintf("Failed to create database: %v", err))
		}

		// Try connecting again
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Failed to connect: %v", err))
		}
	}

	DB = database
	log.Println("🚀 Database Connected!")
}

func createDatabase(user, pass, host, port, sslMode, dbName string) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s",
		host, user, pass, port, sslMode,
	)

	tempDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := tempDB.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	return tempDB.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName)).Error
}

func Migration() {
	stmt := &gorm.Statement{DB: DB}
	for _, m := range models.AllModels() {
		err := DB.AutoMigrate(m)
		if err != nil {
			fmt.Println("Migration failed:", err)
			continue
		}

		stmt.Parse(m)
		fmt.Println("Table created successfully:", stmt.Schema.Table)

		// 🔥 Check if model supports DropColumns
		if dropper, ok := m.(models.ColumnDropper); ok {
			columns := dropper.DropColumns()
			DropColumns(m, columns)
		}

	}

	log.Println("✅ Database migrations completed")

}

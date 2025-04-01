package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/livingdolls/go-template/internal/config"
)

func main() {
	if err := config.LoadConfig("config"); err != nil {
		log.Fatalf("failed to load configuration file: %v", err)
	}

	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		config.Config.Database.Driver,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name,
	)

	// buat instance migrasi
	m, err := migrate.New("file://migrations", dsn)

	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// periksa argumen untuk menentukan perintah migrasi
	if len(os.Args) < 2 {
		log.Fatal("please provide an argument: up | down | force <version>")
	}

	action := os.Args[1]
	switch action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("Please provide a version number for force command")
		}

		version, convErr := strconv.Atoi(os.Args[2])
		if convErr != nil {
			log.Fatalf("Invalid version number: %v", err)
		}

		err = m.Force(version)
	default:
		log.Fatalf("Invalid command. Use 'up' or 'down'.")
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed %v", err)
	}

	log.Printf("Migration completed successfully!")
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	_ "github.com/domibustosb/go_base/db/migrations"
	_ "github.com/domibustosb/go_base/db/seeders"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}
	command := args[1]
	DBDriver := os.Getenv("DB_DRIVER")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBSSLMode := os.Getenv("DB_SSL_MODE")

	db, err := goose.OpenDBWithDriver(DBDriver, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		DBHost,
		DBPort,
		DBName,
		DBUser,
		DBPassword,
		DBSSLMode,
	))
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

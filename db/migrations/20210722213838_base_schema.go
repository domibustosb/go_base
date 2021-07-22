package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upBaseSchema, downBaseSchema)
}

func upBaseSchema(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		CREATE TABLE public.animes (
			id SERIAL NOT NULL,
			name VARCHAR(50),
			created_at timestamptz,
			updated_at timestamptz
		);
	`)
	if err != nil {
		log.Println("Error en la creación. Error: ", err.Error())
	}
	return nil
}

func downBaseSchema(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE public.animes;`)
	if err != nil {
		log.Println("Error en la creación. Error: ", err.Error())
	}
	return nil
}

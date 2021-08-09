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
		INSERT INTO public.animes (name, created_at, updated_at) values ('one piece', now(),now());
	`)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}
	return nil
}

func downBaseSchema(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
	delete from public.animes where name='one piece';
`)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}
	return nil
}

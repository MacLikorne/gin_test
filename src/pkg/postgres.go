package pkg

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"
)

func ConnectToDatabase() *pg.DB {
	addr := fmt.Sprintf("%s:%s",
		GetEnvironmentVariable("QOVERY_POSTGRESQL_ZF3A4FF96_HOST"),
		GetEnvironmentVariable("QOVERY_POSTGRESQL_ZF3A4FF96_PORT"),
	)
	options := &pg.Options{
		User:     GetEnvironmentVariable("QOVERY_POSTGRESQL_ZF3A4FF96_LOGIN"),
		Password: GetEnvironmentVariable("QOVERY_POSTGRESQL_ZF3A4FF96_PASSWORD"),
		Addr:     addr,
		Database: GetEnvironmentVariable("QOVERY_POSTGRESQL_ZF3A4FF96_DEFAULT_DATABASE_NAME"),
	}

	db := pg.Connect(options)
	if db == nil {
		log.Fatal("Failed to connect to database. Closing app.")
	}

	log.Print("Connected to database.")

	return db
}

func CreateNumberTable(db *pg.DB) error {
	createError := db.CreateTable(&Number{}, &orm.CreateTableOptions{IfNotExists: true})
	if createError != nil {
		return createError
	}

	return nil
}

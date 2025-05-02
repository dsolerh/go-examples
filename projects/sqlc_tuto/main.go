package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"reflect"
	authors_repo "sqlc_tuto/pkg/repos/authors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var databaseUrl = "postgres://test:pass@localhost:5432/test?sslmode=disable"

func runMigrations() error {
	migrator, err := migrate.New(
		"file://migrations",
		databaseUrl,
	)
	if err != nil {
		return err
	}
	migrator.Up() // or m.Steps(2) if you want to explicitly set the number of migrations to run
	log.Println("database migrated")

	return nil
}

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	log.Println("connected")

	repo := authors_repo.New(conn)
	// list all authors
	authors, err := repo.ListAuthors(ctx)
	if err != nil {
		return err
	}
	for _, lar := range authors {
		fmt.Printf("lar: %v\n", lar)
	}

	// create an author
	insertedAuthor, err := repo.CreateAuthor(ctx, authors_repo.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: pgtype.Text{
			String: "Co-author of The C Programming Language and The Go Programming Language",
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := repo.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	// run the migrations
	if err := runMigrations(); err != nil {
		log.Fatal(err)
	}
	// run the app
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

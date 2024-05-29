package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func InitDB() {
	var err error
	db, err = sqlx.Open("sqlite3", "file:mydb.duckdb?cache=shared&mode=memory")
	if err != nil {
		log.Fatal(err)
	}
	createTables()
}

func GetDB() *sqlx.DB {
	return db
}

func createTables() {
	createUserTable := `
    CREATE TABLE IF NOT EXISTS Users (
        UserId TEXT PRIMARY KEY,
        Username TEXT NOT NULL UNIQUE,
        Password TEXT NOT NULL,
        Email TEXT NOT NULL UNIQUE,
		Address TEXT NULL,
		DateOfBirth DATE NULL,
		IsActive INTEGER DEFAULT 1,
		IsOrgSet INTEGER DEFAULT 0,
		IsLocked INTEGER DEFAULT 0,
		IsProfileComplete INTEGER DEFAULT 0,
		IsMobileSet INTEGER DEFAULT 0,
		IsMFAEnabled INTEGER DEFAULT 0
    );`

	createSampleProgram := `
	CREATE TABLE SampleProgram (
		SampleProgramId INTEGER PRIMARY KEY,
		Name TEXT NOT NULL,
		IsActive BOOLEAN NOT NULL,
		CreatedOn TIMESTAMP NOT NULL,
		LastModifiedOn TIMESTAMP NOT NULL
	);	
	`

	_, err := db.Exec(createUserTable)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	_, err = db.Exec(createSampleProgram)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

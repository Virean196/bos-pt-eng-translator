package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Virean196/bos-pt-eng-translator/internal/database"
	_ "modernc.org/sqlite"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to get user home dir: %s", err)
	}
	// Create DB connection
	dbPath := "/.local/share/bos-pt-eng-translator/trans.db"
	dbString := filepath.Join(homeDir, dbPath)
	dbConn, err := sql.Open("sqlite", dbString)
	db := database.New(dbConn)

	if err != nil {
		log.Fatalf("Unable to open SQL connection: %s", err)
	}
	if db != nil {
		log.Print("Connection sucessful!")
	}

	// Check ARGS for translator
	if len(os.Args) < 2 {
		log.Fatal("usage: bos <phrase>")
	}
	phrase := os.Args[1]
	translation, err := db.GetPhrase(phrase)
	if translation == "" {
		err = db.CreatePhrase(phrase)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Entry created!")
	} else {
		fmt.Print(translation)
	}
}

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
	langs := map[string]string{
		"en": "en-US|bs-BA",
		"ba": "bs-BA|en-US",
		"pt": "pt|bs-BA",
	}

	// Default langPair from Bosnian to English
	langPair := "bs-BA|en-US"

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

	// Check ARGS for translator
	if len(os.Args) > 2 {
		arg := os.Args[2]
		if pair, ok := langs[arg]; ok {
			langPair = pair
		} else {
			log.Fatalf("Unsupported language: %s, supported: en, bs, pt", arg)
		}
	}
	if len(os.Args) < 2 {
		log.Fatal("usage: bos <phrase> or bos -help")
	}
	phrase := os.Args[1]
	if phrase == "-help" {
		fmt.Printf("- Usage: bos <phrase> <optional: language pair>\n- Available language pairs: en (EN:BA), ba (BA:EN), pt (PT:BA)\n- Language pair defaults to Bosnian to English")
		os.Exit(1)
	}
	translation, err := db.GetPhrase(phrase)
	if translation == "" {
		translation, err = db.CreatePhrase(phrase, langPair)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Translation: %s", translation)
	} else {
		fmt.Printf("Translation: %s", translation)
	}
}

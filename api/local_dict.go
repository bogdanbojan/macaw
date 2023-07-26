package api

import (
	"bufio"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed sqlite_websters_unabridged_dictionary.sql
var embedSQL []byte

type Dictionary struct {
	Word       string
	Wordtype   string
	Definition string
}

func HandleLocalResponse(word string) ([]string, error) {
	var dict Dictionary
	var definitions []string

	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "tmpDict.sql")

	err = os.WriteFile(tmpFile, embedSQL, 0o400)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", tmpFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Setup for MYSQL with Docker-Compose.
	//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/webster_dictionary")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer db.Close()

	res, err := db.Query(fmt.Sprintf(`SELECT entries.definition FROM entries WHERE entries.word = "%s"`, cases.Title(language.AmericanEnglish).String(word)))
	if err != nil {
		return nil, err
	}

	for res.Next() {
		err := res.Scan(&dict.Definition)

		if err != nil {
			return nil, err
		}
		definitions = append(definitions, dict.Definition)
	}

	return definitions, nil
}

func ExtractWords(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

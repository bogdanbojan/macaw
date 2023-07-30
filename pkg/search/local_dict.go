package search

import (
	"bufio"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed assets/sqlite_websters_unabridged_dictionary.sql
var embedSQL []byte

type Dictionary struct {
	Word       string
	Wordtype   string
	Definition string
}

func GetLocalDefinition(word string) (string, error) {
	var dict Dictionary
	var definitions []string

	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "tmpDict.sql")

	err = os.WriteFile(tmpFile, embedSQL, 0o400)
	if err != nil {
		return "", err
	}

	db, err := sql.Open("sqlite3", tmpFile)
	if err != nil {
		return "", err
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf(`SELECT entries.definition FROM entries WHERE entries.word = "%s"`, cases.Title(language.AmericanEnglish).String(word)))
	if err != nil {
		return "", err
	}

	for res.Next() {
		err := res.Scan(&dict.Definition)

		if err != nil {
			return "", err
		}
		definitions = append(definitions, dict.Definition)
	}

	if definitions == nil {
		return "", errors.New("Word not found")
	}
	// TODO: Think of better name handling here.
	var def string
	for i, v := range definitions {
		def += fmt.Sprintf("[%d] %s \n", i, v)
	}

	return def, nil
}

func ExtractWords(fileName string) []string {
	fileName = filepath.Clean(fileName)
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

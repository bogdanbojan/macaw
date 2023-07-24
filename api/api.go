package api

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	gowiki "github.com/trietmn/go-wiki"
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

func SearchWiki(word string) (string, error) {
	res, err := gowiki.Summary(word, 5, -1, false, true)
	if err != nil {
		return "", err
	}

	return res, nil
}

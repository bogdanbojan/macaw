package api

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Response struct {
	Word       string     `json:"word"`
	Phonetic   string     `json:"phonetic"`
	Phonetics  Phonetics  `json:"phoenitcs"`
	Meanings   []Meanings `json:"meanings"`
	License    License    `json:"license"`
	SourceUrls []string   `json:"sourceUrls"`
}

type Phonetics struct {
	Text      string `json:"text"`
	Audio     string `json:"audio"`
	SourceURL string `json:"sourceUrl,omitempty"`
	License   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license,omitempty"`
}

type Meanings struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Definitions  []struct {
		Definition string `json:"definition"`
		Synonyms   []any  `json:"synonyms"`
		Antonyms   []any  `json:"antonyms"`
		Example    string `json:"example,omitempty"`
	} `json:"definitions"`
	Synonyms []any `json:"synonyms"`
	Antonyms []any `json:"antonyms"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
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

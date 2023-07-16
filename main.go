package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	gowiki "github.com/trietmn/go-wiki"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed sql-scripts/sqlite_websters_unabridged_dictionary.sql
var embedSQL []byte

type Dictionary struct {
	Word       string
	Wordtype   string
	Definition string
}

func main() {
	ShowGUI()

	// fileName := "words.txt"
	// words := extractWords(fileName)
	// apiRequest(words)
}

//	func extractWords(fileName string) []string {
//		file, err := os.Open(fileName)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//
//		var words []string
//		for scanner.Scan() {
//			words = append(words, scanner.Text())
//		}
//
//		if err := scanner.Err(); err != nil {
//			log.Fatal(err)
//		}
//
//		return words
//	}
// func apiRequest(words []string) {
// 	reqURL := "https://api.dictionaryapi.dev/api/v2/entries/en/"
// 	for _, w := range words {
// 		resp, err := http.Get(reqURL + w)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
//
// 		handleLocalResponse(w)
// 		handleServerResponse(resp, w)
// 	}
// }

func handleLocalResponse(word string) ([]string, error) {
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
		log.Println(dict.Definition)
	}

	return definitions, nil
}

//
//	func handleServerResponse(resp *http.Response, word string) {
//		defer resp.Body.Close()
//		body, err := io.ReadAll(resp.Body)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//		var r []Response
//		err = json.Unmarshal(body, &r)
//
//		if err != nil {
//			fmt.Println("Could not find word in the dictionary")
//			searchWiki(word)
//		}
//
//		s, _ := json.MarshalIndent(r, "", "\t")
//		fmt.Println(string(s))
//	}

func SearchWiki(word string) (string, error) {
	res, err := gowiki.Summary(word, 5, -1, false, true)
	if err != nil {
		return "", err
	}

	return res, nil
}

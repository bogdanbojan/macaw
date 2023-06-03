package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
		License   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"license,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
			Example    string `json:"example,omitempty"`
		} `json:"definitions"`
		Synonyms []any `json:"synonyms"`
		Antonyms []any `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}

func main() {
	fileName := "words.txt"
	words := extractWords(fileName)
	apiRequest(words)
}

func extractWords(fileName string) []string {
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

func apiRequest(words []string) {
	reqURL := "https://api.dictionaryapi.dev/api/v2/entries/en/"
	for _, w := range words {
		resp, err := http.Get(reqURL + w)
		if err != nil {
			fmt.Println(err)
		}
		handleResponse(resp)
	}
}

func handleResponse(resp *http.Response) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var r []Response
	json.Unmarshal(body, &r)

	s, _ := json.MarshalIndent(r, "", "\t")
	fmt.Println(string(s))
}

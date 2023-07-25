package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func ApiRequest(words []string) ([]string, error) {
	reqURL := "https://api.dictionaryapi.dev/api/v2/entries/en/"
	var response []string
	for _, w := range words {
		resp, err := http.Get(reqURL + w)
		if err != nil {
			fmt.Println(err)
		}

		// TODO: Handle error when request does not find the word.
		s, err := HandleServerResponse(resp, w)
		if err != nil {
			return nil, err
		}
		response = append(response, s)
	}
	return response, nil

}

func HandleServerResponse(resp *http.Response, word string) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var r []Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}

	s, _ := json.MarshalIndent(r, "", "\t")
	return string(s), nil
}

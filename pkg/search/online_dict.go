package search

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response represents the structure of the API call json response that we get
// from api.dictionaryapi.dev.
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

// GetOnlineDefinition fetches the definition from the online API and unmarshals 
// it into a Response struct based on the word parameter.
func GetOnlineDefinition(word string) (string, error) {
    reqURL := "https://api.dictionaryapi.dev/api/v2/entries/en/"
	resp, err := http.Get(reqURL + word)
	if err != nil {
		return "", err
	}

	// TODO: Handle error when request does not find the word.
	response, err := HandleServerResponse(resp, word)
	if err != nil {
		return "", err
	}
	return response, nil

}

// HandleServerResponse unmarshals, indents and formats the http response into a string.
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

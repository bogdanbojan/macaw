package main

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

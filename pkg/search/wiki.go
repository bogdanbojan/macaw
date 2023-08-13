package search

import gowiki "github.com/trietmn/go-wiki"

// GetWikipediaSummary fetches the summary of a word using the Wikipedia API.
func GetWikipediaSummary(word string) (string, error) {
	res, err := gowiki.Summary(word, 5, -1, false, true)
	if err != nil {
		return "", err
	}

	return res, nil
}

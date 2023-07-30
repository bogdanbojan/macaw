package search

import (
	"fmt"
)

type Searcher interface {
	Definition(word string) (string, error)
	Definitions(words []string) (string, error)
}

type Sources struct {
	LocalDictionary
	OnlineDictionary
	WikipediaSummary
	SliderValues []float64
}

type LocalDictionary struct {
	// sliderValue float64
}
type OnlineDictionary struct {
	// sliderValue float64
}
type WikipediaSummary struct {
	// sliderValue float64
}

func (*LocalDictionary) Definition(word string) (string, error) {
	res, err := GetLocalDefinition(word)
	if err != nil {
		return "", err
	}

	return res, nil
}

// TODO: Think about if we want this code repetition or we can just pass a
// a func as a parameter.
func (*LocalDictionary) Definitions(words []string) (string, error) {
	var results []string
	var failedResults []string
	for _, w := range words {
		res, err := GetLocalDefinition(w)
		if err != nil {
			failedResults = append(failedResults, w)
			continue
		}
		results = append(results, fmt.Sprint(w+"\n")+res)
	}

	var res string
	for _, v := range results {
		res += fmt.Sprintf(" %s \n", v)
	}

	if len(failedResults) != 0 {
		res += "Could not find the following words: \n"
		for _, v := range failedResults {
			res += v + "\n"
		}
	}

	return res, nil
}

func (*OnlineDictionary) Definition(word string) (string, error) {
	res, err := GetOnlineDefinition(word)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (*OnlineDictionary) Definitions(words []string) (string, error) {
	var results []string
	var failedResults []string
	for _, w := range words {
		res, err := GetOnlineDefinition(w)
		if err != nil {
			failedResults = append(failedResults, w)
			continue
		}
		results = append(results, fmt.Sprint(w+"\n")+res)
	}

	var res string
	for _, v := range results {
		res += fmt.Sprintf(" %s \n", v)
	}

	if len(failedResults) != 0 {
		res += "Could not find the following words: \n"
		for _, v := range failedResults {
			res += v + "\n"
		}
	}

	return res, nil
}

func (*WikipediaSummary) Definition(word string) (string, error) {
	res, err := GetWikipediaSummary(word)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (*WikipediaSummary) Definitions(words []string) (string, error) {
	var results []string
	var failedResults []string
	for _, w := range words {
		res, err := GetWikipediaSummary(w)
		if err != nil {
			failedResults = append(failedResults, w)
			continue
		}
		results = append(results, fmt.Sprint(w+"\n")+res)
	}

	var res string
	for _, v := range results {
		res += fmt.Sprintf(" %s \n", v)
	}

	if len(failedResults) != 0 {
		res += "Could not find the following words: \n"
		for _, v := range failedResults {
			res += v + "\n"
		}
	}

	return res, nil
}

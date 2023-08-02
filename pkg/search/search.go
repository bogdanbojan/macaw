package search

import (
	"context"
	"fmt"
	"log"
)

// TODO: Duplicate vars due to circular import..
var (
	LOCAL  = "LOCAL"
	ONLINE = "ONLINE"
	WIKI   = "WIKI"
)

// TODO: Not sure the context options should live here..
var OPTIONS = contextKey("OPTIONS")

type contextKey string

func (c contextKey) String() string {
	return string(c) + "key"
}

// TODO: Have a struct as a response.
type Searcher interface {
	Search(ctx context.Context, words []string) (definitions string, err error)
}

type Sources struct {
	LocalDictionary
	OnlineDictionary
	WikipediaSummary
}

// TODO: Think about if the definitions field is even worth keeping.
type LocalDictionary struct {
	definitions string
}

type OnlineDictionary struct {
	definitions string
}

type WikipediaSummary struct {
	definitions string
}

func (s *Sources) Search(ctx context.Context, words []string) (definitions string, err error) {
	opt := ctx.Value(OPTIONS).(map[string]float64)

	switch {
	case opt[LOCAL] == 1:
        // Check that we searched for the term so we don't always pattern match
        // on this option. This is ugly.
        opt[LOCAL] = 0
		s.LocalDictionary.definitions, err = s.LocalDictionary.Search(words)
		if err != nil {
			return "", err
		}
		return s.LocalDictionary.definitions, nil

	case opt[ONLINE] == 1:
        opt[ONLINE] = 0
		s.OnlineDictionary.definitions, err = s.OnlineDictionary.Search(words)
		if err != nil {
			return "", err
		}
		return s.OnlineDictionary.definitions, nil

	case opt[WIKI] == 1:
        opt[WIKI] = 0
		s.WikipediaSummary.definitions, err = s.WikipediaSummary.Search(words)
		if err != nil {
			return "", err
		}
		return s.WikipediaSummary.definitions, nil
	}

	return "", nil
}

func (*LocalDictionary) Search(words []string) (definitions string, err error) {
	dd, fdd := defineWords(words, GetLocalDefinition)
	if len(dd) == 1 {
		return dd[0], nil
	}

	d := toString(dd, fdd)

	return d, nil
}

func (*OnlineDictionary) Search(words []string) (definitions string, err error) {
	dd, fdd := defineWords(words, GetOnlineDefinition)
	if len(dd) == 1 {
        log.Println(dd[0])
		return dd[0], nil
	}

	d := toString(dd, fdd)

	return d, nil
}

func (*WikipediaSummary) Search(words []string) (definitions string, err error) {
	dd, fdd := defineWords(words, GetWikipediaSummary)
	if len(dd) == 1 {
		return dd[0], nil
	}

	d := toString(dd, fdd)

	return d, nil
}

type getDefFunc func(word string) (string, error)

func defineWords(words []string, f getDefFunc) (definitions, failedDefinitions []string) {
	var dd []string
	var fd []string
	for _, w := range words {
		d, err := f(w)
		if err != nil {
			fd = append(fd, w)
			continue
		}
		dd = append(dd, fmt.Sprint(w+"\n")+d)
	}

	return dd, fd
}

func toString(definitions, failedDefinitions []string) string {
	var stringdd string
	for _, d := range definitions {
		stringdd += fmt.Sprintf(" %s \n", d)
	}

	if len(failedDefinitions) != 0 {
		stringdd += "Could not find the following words: \n"
		for _, fd := range failedDefinitions {
			stringdd += fd + "\n"
		}
	}
	return stringdd
}

package search

import (
	"context"
)

// Existing options to search from.
var (
	LOCAL  = "LOCAL"
	ONLINE = "ONLINE"
	WIKI   = "WIKI"
)

// ContextKeyOptions represents the options for the key for the context value which 
// fetches the possibilities to search from.
var ContextKeyOptions = contextKey("OPTIONS")

// contextKey represents the type of value for the context key.
type contextKey string

// String() formats the context key to a custom string.
func (c contextKey) String() string {
	return string(c) + "key"
}

// OptionsKey checks that the options are present in the ctx.
func OptionsKey(ctx context.Context) (string, bool) {
	optStr, ok := ctx.Value(ContextKeyOptions).(string)
	return optStr, ok
}

// Searcher interface represents the default search behaviour. 
type Searcher interface {
	Search(ctx context.Context, words []string) []Definition
}

// Sources manages all the search sources.
type Sources struct {
	localDictionary
	onlineDictionary
	wikipediaSummary
}

// Definition is the response of the Searcher interface and the default way to 
// represent a definition within the app.
type Definition struct {
	Word string
	Text string
	Ok   bool
}

// localDictionary holds the definitions collected from the local SQLite database.
type localDictionary struct {
	dd []Definition
}

// onlineDictionary holds the definitions collected from the dictionary API.
type onlineDictionary struct {
	dd []Definition
}

// wikipediaSummary holds the definitions collected from the wikipedia API.
type wikipediaSummary struct {
	dd []Definition
}

// Search gets the proper definition collection, based on the search flag that 
// exists in the passed context.
func (s *Sources) Search(ctx context.Context, words []string) []Definition {
	opt := ctx.Value(ContextKeyOptions).(map[string]float64)

	switch {
	case opt[LOCAL] == 1:
		opt[LOCAL] = 0
		s.localDictionary.Search(words)
		return s.localDictionary.dd

	case opt[ONLINE] == 1:
		opt[ONLINE] = 0
		s.onlineDictionary.Search(words)
		return s.onlineDictionary.dd

	case opt[WIKI] == 1:
		opt[WIKI] = 0
		s.wikipediaSummary.Search(words)
		return s.wikipediaSummary.dd
	}

	return nil
}

// Search collects the definitions from the local dictionary.
func (ld *localDictionary) Search(words []string) {
	ld.dd = defineWords(words, GetLocalDefinition)
}

// Search collects the definitions from the dictionary API.
func (od *onlineDictionary) Search(words []string) {
	od.dd = defineWords(words, GetOnlineDefinition)
}

// Search collects the definitions from the wikipedia API.
func (ws *wikipediaSummary) Search(words []string) {
	ws.dd = defineWords(words, GetWikipediaSummary)
}

// getDefFunc is the signature of any definition/summary getter.
type getDefFunc func(word string) (string, error)

// defineWords calls a passed in function in order to get the correct definition
// for the helper Search function which have the Source fields as a method receiver.
func defineWords(words []string, f getDefFunc) []Definition {
	var dd []Definition
	for _, w := range words {
		var d Definition
		var err error

		d.Text, err = f(w)
		if err != nil {
			// The default bool value is false.
			dd = append(dd, Definition{Word: w})
			continue
		}
		d.Word = w
		d.Ok = true
		dd = append(dd, d)
	}

	return dd
}

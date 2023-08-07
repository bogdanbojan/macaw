package search

import (
	"context"
)

var (
	LOCAL  = "LOCAL"
	ONLINE = "ONLINE"
	WIKI   = "WIKI"
)

var ContextKeyOptions = contextKey("OPTIONS")

type contextKey string

func (c contextKey) String() string {
	return string(c) + "key"
}

// OptionsKey gets the option setup from the context.
func OptionsKey(ctx context.Context) (string, bool) {
	optStr, ok := ctx.Value(ContextKeyOptions).(string)
	return optStr, ok
}

type Searcher interface {
	Search(ctx context.Context, words []string) []Definition
}

type Sources struct {
	localDictionary
	onlineDictionary
	wikipediaSummary
}

type Definition struct {
	Word string
	Text string
	Ok   bool
}

type localDictionary struct {
	dd []Definition
}

type onlineDictionary struct {
	dd []Definition
}

type wikipediaSummary struct {
	dd []Definition
}

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

func (ld *localDictionary) Search(words []string) {
	ld.dd = defineWords(words, GetLocalDefinition)
}

func (od *onlineDictionary) Search(words []string) {
	od.dd = defineWords(words, GetOnlineDefinition)
}

func (ws *wikipediaSummary) Search(words []string) {
	ws.dd = defineWords(words, GetWikipediaSummary)
}

type getDefFunc func(word string) (string, error)

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

package main

import (
	_ "embed"

	"github.com/bogdanbojan/macaw/gui"
	_ "github.com/mattn/go-sqlite3"
	gowiki "github.com/trietmn/go-wiki"
)

func main() {
	gui.ShowGUI()

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

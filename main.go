package main

import (
	"github.com/bogdanbojan/loci/gui"
	_ "github.com/go-sql-driver/mysql"
)

type Dictionary struct {
	Word       string
	Wordtype   string
	Definition string
}

func main() {
    gui.ShowGUI()
//	fileName := "words.txt"
//	words := extractWords(fileName)
//	apiRequest(words)
}

// func extractWords(fileName string) []string {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 
// 	scanner := bufio.NewScanner(file)
// 
// 	var words []string
// 	for scanner.Scan() {
// 		words = append(words, scanner.Text())
// 	}
// 
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 
// 	return words
// }
// 
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
// func handleLocalResponse(word string) {
// 	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/webster_dictionary")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("Cannot ping db")
// 	}
// 
// 	res, err := db.Query(fmt.Sprintf(`SELECT entries.definition FROM entries WHERE entries.word = "%s"`, word))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 
// 	for res.Next() {
// 		var dict Dictionary
// 		err := res.Scan(&dict.Definition)
// 
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 
// 		fmt.Println("Definition: ", dict)
// 	}
// 
// 	fmt.Println("==============================================================")
// 
// }
// 
// func handleServerResponse(resp *http.Response, word string) {
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var r []Response
// 	err = json.Unmarshal(body, &r)
// 
// 	if err != nil {
// 		fmt.Println("Could not find word in the dictionary")
// 		searchWiki(word)
// 	}
// 
// 	s, _ := json.MarshalIndent(r, "", "\t")
// 	fmt.Println(string(s))
// }
// 
// func searchWiki(word string) {
// 	res, err := gowiki.Summary(word, 5, -1, false, true)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("Summary: %v\n", res)
// }

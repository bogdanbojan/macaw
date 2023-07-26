package main

import (
	"github.com/bogdanbojan/macaw/gui"
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

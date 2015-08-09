package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
)

// translate is looking for translation
// its using dict.cc
func translate(word string) {
	doc, err := goquery.NewDocument(fmt.Sprintf("http://dict.cc/?s=%s", word))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	re := regexp.MustCompile("\\d")

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		first := s.Find("td.td7nl").First().Text()
		last := s.Find("td.td7nl").Last().Text()
		if len(last) > 0 || len(last) > 0 {
			fmt.Printf("%s - %s\n", first, re.ReplaceAllString(last, ""))
		}
	})
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Need a word to translate, like ./go-translator hello")
		os.Exit(1)
	}
	translate(os.Args[1])
}

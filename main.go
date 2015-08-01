package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
)

func translate(word string) {
	doc, err := goquery.NewDocument(fmt.Sprintf("http://dict.cc/?s=%s", word))
	if err != nil {
		log.Fatal(err)
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
	translate(os.Args[1])
}

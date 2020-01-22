package contacts

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
)

func Convert(filepath string, ignoreDeletedAccounts bool) {
	// Request the HTML page.
	content := getFileAsDocument(filepath)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		log.Fatalln(err)
	}

	contacts := parseHtml(*doc)
	for _, c := range contacts {
		numbers := ""
		for i, n := range c.numbers {
			if i == 0 {
				numbers += n
			} else {
				numbers += ", " + n
			}
		}
		msg := fmt.Sprintf("%s: %s", c.name, numbers)
		log.Println(msg)
	}
}

func getFileAsDocument(path string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	//defer file.Close()

	return file
}

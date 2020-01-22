package contacts

import (
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
	res := ""
	for i, c := range contacts {
		if i != 0 {
			res += "\n"
		}
		vCard := createTemplate(c)
		res += vCard
	}

	log.Println(res)
}

func getFileAsDocument(path string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	//defer file.Close()

	return file
}

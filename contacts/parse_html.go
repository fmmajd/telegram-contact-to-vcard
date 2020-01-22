package contacts

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func parseHtml(doc goquery.Document) []contact {
	contacts := map[string]contact{}
	doc.Find(".entry.clearfix").Each(func(i int, entry *goquery.Selection) {
		name := entry.Find(".name").Text()
		name = trimContactInfo(name)

		phone := entry.Find(".details_entry.details").Text()
		phone = trimContactInfo(phone)

		contacts = appendToContacts(contacts, name, phone)
	})

	var result []contact
	for _, v := range contacts {
		result = append(result, v)
	}
	return result
}

func appendToContacts(contacts map[string]contact, name string, number string) map[string]contact {
	contact, exists := contacts[name]
	var numbers []string
	if exists {
		numbers = append(contact.numbers, number)
	} else {
		contact.name = name
		numbers = []string {number}
	}

	contact.numbers = numbers
	contacts[name] = contact

	return contacts
}

func trimContactInfo(s string) string {
	newlineChars := []string{
		"\r\n",
		"\r",
		"\n",
		"\t",
		"\\s",
		"\v",
		"\f",
		"\u0085",
		"\u2028",
		"\u2029",
	}
	for _, ch := range newlineChars {
		s = strings.Trim(s, ch)
	}
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")

	//this part is very important. for some reason, go can not trim \s end-of-line characters from these docs
	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, "")

	return s
}

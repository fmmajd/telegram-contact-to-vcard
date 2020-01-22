package contacts

import (
	"fmt"
	"strings"
	"time"
)

var vcard_template_example = `BEGIN:VCARD
VERSION:3.0
item1.EMAIL;TYPE=INTERNET,pref:bfnny5@yahoo.com
TEL;TYPE=CELL,VOICE,pref:+989333958248
NOTE:\n\n
REV:20151108T093348Z
N:Abadi;Nastaran;;;
FN:
item1.X-ABLABEL:_$!<Other>!$_
PRODID:ez-vcard 0.9.14-fc
END:VCARD`

var vcard_template = `BEGIN:VCARD
VERSION:3.0
%s
REV:%s
N:%s
FN:%s
PRODID:ez-vcard 0.9.14-fc
END:VCARD`

func createTemplate(c contact) string {
	numbers := vCardNumbersLines(c)
	firstName, lastName := firstNameAndLastName(c)
	nAttr := vCardNAttribute(firstName, lastName)
	fnnAttr := vCardFNAttribute(firstName, lastName)
	revAttr := vCardRevAttribute()
	vCard := fmt.Sprintf(vcard_template, numbers, revAttr, nAttr, fnnAttr)

	return vCard
}

func vCardNumbersLines(c contact) string{
	numbers := ""
	for i, n := range c.numbers {
		if i != 0 {
			numbers += "\n"
		}
		numbers += fmt.Sprintf("TEL;TYPE=CELL,VOICE,pref:%s", n)
	}

	return numbers
}

func firstNameAndLastName(c contact) (string, string) {
	name := c.name
	nameSplit := strings.SplitAfterN(name, " ", 2)
	firstName := ""
	lastName := ""
	if len(nameSplit) > 0 {
		firstName = nameSplit[0]
	}
	if len(nameSplit) > 1 {
		lastName = nameSplit[1]
	}

	return firstName, lastName
}

func vCardNAttribute(firstName string, lastName string) string {
	firstName = strings.TrimRight(firstName, " ")
	nAttr := fmt.Sprintf("%s;%s;;;", lastName, firstName)

	return nAttr
}

func vCardFNAttribute(firstName string, lastName string) string {
	nAttr := fmt.Sprintf("%s%s", firstName, lastName)

	return nAttr
}

func vCardRevAttribute() string {
	//example: 20151108T093348Z
	now := time.Now()
	format := "20060102T030405Z"
	return now.Format(format)
}
package main

import (
	"fatemeh.dev/telegram-contacts-converter/contacts"
	"log"
	"os"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	filename := curDir + "/data/lists/contacts.html"
	ignoreDeletedAccounts := true

	contacts.Convert(filename, ignoreDeletedAccounts)
}
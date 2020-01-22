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
	beginPath := curDir + "/data/lists/contacts.html"
	targetPath := curDir + "/data/res.vcf"
	ignoreDeletedAccounts := true

	contacts.Convert(beginPath, targetPath, ignoreDeletedAccounts)
}
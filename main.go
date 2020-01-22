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

	cmdArgs := os.Args
	ignoreDeletedAccounts := true
	if len(cmdArgs) > 1 {
		arg := cmdArgs[1]
		if arg == "--add-deleted" {
			ignoreDeletedAccounts = false
		}
	}

	contacts.Convert(beginPath, targetPath, ignoreDeletedAccounts)
}
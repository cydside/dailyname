package main

import (
	"log"

	"github.com/cydside/dailyname_prova/cmd"
)

//______________________________________________________________________________

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
}

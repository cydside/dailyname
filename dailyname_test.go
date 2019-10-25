package dailyname_test

import (
	"log"
	"testing"

	"github.com/cydside/dailyname/infrastructure"
)

//______________________________________________________________________________

// TestOneNode testa il
// go test -v -run ^TestCli$
func TestCli(t *testing.T) {
	err := infrastructure.Execute()
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
}

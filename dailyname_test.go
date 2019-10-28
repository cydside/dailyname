package dailyname_test

import (
	"log"
	"testing"

	"github.com/cydside/dailyname"
)

//______________________________________________________________________________

// TestOneNode testa il
// go test -v -run ^TestCli$
func TestCli(t *testing.T) {
	var err error

	ur := &dailyname.UserReq{
		DateFrom: "2019",
		DateTo:   "",
	}

	dailyname.SetCommands(ur)

	log.Printf("%v", ur)

	ur = &dailyname.UserReq{
		DateFrom: "2020-02",
		DateTo:   "",
	}

	dailyname.SetCommands(ur)

	log.Printf("%v", ur)

	ur = &dailyname.UserReq{
		DateFrom: "2020-02",
		DateTo:   "2020-10-03",
	}

	dailyname.SetCommands(ur)

	log.Printf("%v", ur)

	ur = &dailyname.UserReq{
		DateFrom: "2020-02-10",
		DateTo:   "2020-02-15",
	}

	err = dailyname.SetCommands(ur)
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%v", ur)

	ur = &dailyname.UserReq{
		DateFrom: "2020-03",
		DateTo:   "2020-02",
	}

	err = dailyname.SetCommands(ur)
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%v", ur)
	// err := infrastructure.Execute()
	// if err != nil {
	// 	log.Printf("ERROR: %s\n", err.Error())
	// }
}

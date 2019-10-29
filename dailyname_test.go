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
	var arrstr []string
	var ur dailyname.UserReq

	ur = dailyname.UserReq{
		DateFrom: "2019",
		DateTo:   "",
	}

	arrstr, err = dailyname.GetDailyNames(&ur)
	if err != nil {
		t.Errorf("Error(): %s", err.Error())
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)

	ur = dailyname.UserReq{
		Lang:      "it_IT",
		DateFrom:  "2019-05",
		DateTo:    "",
		LoneOrSub: true,
		DoW:       -1,
		Content:   "RX,TX",
	}

	arrstr, err = dailyname.GetDailyNames(&ur)
	if err != nil {
		t.Errorf("Error(): %s", err.Error())
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)

	ur = dailyname.UserReq{
		Lang:      "it_IT",
		DateFrom:  "2019-02",
		DateTo:    "",
		LoneOrSub: true,
		// DoW:       2,
		// Content:   "RX,TX",
	}

	arrstr, err = dailyname.GetDailyNames(&ur)
	if err != nil {
		t.Errorf("Error(): %s", err.Error())
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)

	ur = dailyname.UserReq{
		Lang:     "it_IT",
		DateFrom: "2020-02",
		DateTo:   "2020-10-03",
	}

	arrstr, err = dailyname.GetDailyNames(&ur)
	if err != nil {
		t.Errorf("Error(): %s", err.Error())
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)

	ur = dailyname.UserReq{
		Lang:     "it_IT",
		DateFrom: "2020-02-10",
		DateTo:   "2020-02-15",
	}

	arrstr, err = dailyname.GetDailyNames(&ur)
	if err != nil {
		t.Errorf("Error(): %s", err.Error())
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)

	ur = dailyname.UserReq{
		Lang:     "it_IT",
		DateFrom: "2020-03",
		DateTo:   "2020-02",
	}

	arrstr, err = dailyname.GetDailyNames(&ur)
	if err == nil {
		t.Errorf("Error(): %s", "Swap the dates!")
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)
}

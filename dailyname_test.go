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

	// ur = dailyname.UserReq{
	// 	Lang:     "it_IT",
	// 	DateFrom: "2019",
	// 	DateTo:   "",
	// }

	// arrstr, err = dailyname.GetDailyNames(&ur)
	// if err != nil {
	// 	log.Printf("%s", err.Error())
	// }

	// log.Printf("%v", ur)
	// log.Printf("%v", arrstr)

	ur = dailyname.UserReq{
		Lang:      "it_IT",
		DateFrom:  "2019-02",
		DateTo:    "",
		LoneOrSub: true,
		DoW:       2,
		Content:   "RX,TX",
	}

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
		log.Printf("%s", err.Error())
	}

	log.Printf("%v", ur)
	log.Printf("%v", arrstr)

	// ur = &dailyname.UserReq{
	// 	DateFrom: "2020-02",
	// 	DateTo:   "",
	// }

	// dailyname.GetDailyNames(ur)

	// log.Printf("%v", ur)

	// ur = &dailyname.UserReq{
	// 	DateFrom: "2020-02",
	// 	DateTo:   "2020-10-03",
	// }

	// dailyname.GetDailyNames(ur)

	// log.Printf("%v", ur)

	// ur = &dailyname.UserReq{
	// 	DateFrom: "2020-02-10",
	// 	DateTo:   "2020-02-15",
	// }

	// err = dailyname.GetDailyNames(ur)
	// if err != nil {
	// 	log.Printf("%s", err.Error())
	// }

	// log.Printf("%v", ur)

	// ur = &dailyname.UserReq{
	// 	DateFrom: "2020-03",
	// 	DateTo:   "2020-02",
	// }

	// err = dailyname.GetDailyNames(ur)
	// if err != nil {
	// 	log.Printf("%s", err.Error())
	// }

	// log.Printf("%v", ur)
	// err := infrastructure.Execute()
	// if err != nil {
	// 	log.Printf("ERROR: %s\n", err.Error())
	// }
}

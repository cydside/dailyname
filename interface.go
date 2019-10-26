//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

import (
	"time"

	"github.com/pkg/errors"
)

//______________________________________________________________________________

// SetCommands imposta le opzioni disponibili per la crazione dei nomi
func SetCommands(obj *UserReq) error {
	var err error

	err = validateCommands(obj)

	return err
}

//______________________________________________________________________________

// type UserReq struct {
// 	Lang      string `json:"lang"`          // Language selected(Italian('it_IT' default) and English('en_US') are pre-loaded).
// 	PathSep   string `json:"pathsep"`       // Path separator for the desired OS only on LoneOrSub = false.
// 	DateFrom  string `json:"datefrom"`      // The starting date in format YYYYMMDD(eg.: 2019-11-08 or 2019-11 or 2019) on missing data will be selected the first avaiable. If the final date won't be supplied the last day of the year will be the selected.
// 	DateTo    string `json:"dateto"`        // The final datein format YYYYMMDD(eg.: 2020-01-16 or  2020-01) on missing data will be selected the last avaiable without crossing year, if the starting date won't be supplied the first day of the year will be the selected.
// 	Suffix    string `json:"suffix"`        // Add suffix to day folder's name.
// 	Prefix    string `json:"prefix"`        // Add prefix to day folder's name.
// 	LoneOrSub bool   `json:"loneorsub"`     // True: A long forlder name per day(eg: 2017-01-22); False: a forlder for the year, subfolders for the months and subfolders for the days(default).
// 	Duration  uint16 `json:"duration"`      // Duration in number of the days(1 to 366)
// 	DoW       uint8  `json:"dayoftheweek"`  // Day's name added: Monday(long format), Mon(short format)
// 	DoM       uint8  `json:"dayofthemonth"` // Month's name added: Jenuary(long format), Jen(short format)
// 	DoY       bool   `json:"dayoftheyear"`  // Julian date added: 001 for jenuary first
// }

func validateCommands(obj *UserReq) error {
	var err error

	if obj.DateFrom, obj.DateTo, err = checkPeriod(obj.DateFrom, obj.DateTo); err != nil {
		return err
	}

	return err
}

//______________________________________________________________________________

func checkPeriod(start, end string) (string, string, error) {
	var dateStart, dateEnd time.Time
	var err error

	if start == "" && end == "" {
		return "", "", errors.New("No period specified, bye!")
	}

	m := "Wrong start date, bye!"

	switch len(start) {
	case 10:
		if dateStart, err = checkStringDate(start); err != nil {
			return "", "", errors.Wrap(err, m)
		}
	case 7:
		if dateStart, err = checkStringDate(start + "-01"); err != nil {
			return "", "", errors.Wrap(err, m)
		}
	case 4:
		if dateStart, err = checkStringDate(start + "-01-01"); err != nil {
			return "", "", errors.Wrap(err, m)
		}
	default:
		return "", "", errors.New(m)
	}

	m = "Wrong end date, bye!"

	switch len(end) {
	case 10:
		if dateEnd, err = checkStringDate(end); err != nil {
			return "", "", errors.Wrap(err, m)
		}
	case 7:
		if dateEnd, err = checkStringDate(end + "-01"); err != nil {
			return "", "", errors.Wrap(err, m)
		}
	case 0:
		if dateEnd, err = checkStringDate(start[:4] + "-12-31"); err != nil {
			return "", "", errors.Wrap(err, m)
		}
	default:
		return "", "", errors.New(m)
	}

	de := dateStart.AddDate(1, 0, 0)
	if de.Before(dateEnd) {
		return "", "", errors.New("No more than a year is allowed, bye!")
	}

	return dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"), nil
}

//______________________________________________________________________________

func checkStringDate(sDate string) (time.Time, error) {
	return time.Parse("2006-01-02", sDate)
}

//______________________________________________________________________________
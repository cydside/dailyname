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

	if err = checkPeriod(obj); err != nil {
		return err
	}

	return err
}

//______________________________________________________________________________

func checkPeriod(obj *UserReq) error {
	var dateStart, dateEnd time.Time
	var err error

	start := obj.DateFrom
	end := obj.DateTo

	ms := "Wrong start date, bye!"
	me := "No valid end date, bye!"

	if start == "" {
		return errors.New(ms)
	}

	switch len(start) {
	case 10:
		if dateStart, err = checkStringDate(start); err != nil {
			return errors.Wrap(err, ms)
		}

		if len(end) == 0 {
			if obj.Duration > 0 && obj.Duration <= 366 {
				dateEnd = dateStart.AddDate(0, 0, obj.Duration)
			} else {
				return errors.New(me)
			}
		} else if len(end) == 7 {
			if dateEnd, err = checkStringDate(end + "-01"); err != nil {
				return errors.Wrap(err, me)
			}
			_, dateEnd = monthInterval(dateEnd)
		} else if len(end) == 10 {
			if dateEnd, err = checkStringDate(end); err != nil {
				return errors.Wrap(err, me)
			}
		} else {
			return errors.New(me)
		}
	case 7:
		if dateStart, err = checkStringDate(start + "-01"); err != nil {
			return errors.Wrap(err, ms)
		}

		if len(end) == 0 {
			_, dateEnd = monthInterval(dateStart)
		} else if len(end) == 7 {
			if dateEnd, err = checkStringDate(end + "-01"); err != nil {
				return errors.Wrap(err, me)
			}
			_, dateEnd = monthInterval(dateEnd)
		} else if len(end) == 10 {
			if dateEnd, err = checkStringDate(end); err != nil {
				return errors.Wrap(err, me)
			}
		} else {
			return errors.New(me)
		}
	case 4:
		if dateStart, err = checkStringDate(start + "-01-01"); err != nil {
			return errors.Wrap(err, ms)
		}

		if dateEnd, err = checkStringDate(start + "-12-31"); err != nil {
			return errors.Wrap(err, ms)
		}
	default:
		return errors.New(ms)
	}

	if dateEnd.Before(dateStart) {
		td := dateEnd
		dateEnd = dateStart
		dateStart = td
	}

	de := dateStart.AddDate(1, 0, 0)
	if de.Before(dateEnd) {
		return errors.New("No more than a year is allowed, bye!")
	}

	obj.DateFrom = dateStart.Format("2006-01-02")
	obj.DateTo = dateEnd.Format("2006-01-02")

	return nil
}

//______________________________________________________________________________

func checkStringDate(sDate string) (time.Time, error) {
	return time.Parse("2006-01-02", sDate)
}

//______________________________________________________________________________

func monthInterval(tDate time.Time) (firstOfMonth, lastOfMonth time.Time) {
	currentYear, currentMonth, _ := tDate.Date()
	firstOfMonth = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth = firstOfMonth.AddDate(0, 1, -1)

	return
}

//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

import (
	"os"
	"time"

	"github.com/pkg/errors"
)

//______________________________________________________________________________

// GetDailyNames it returns a string array of daily names
func GetDailyNames(obj *UserReq) ([]string, error) {
	var err error
	var arrstr []string

	if err = setCommands(obj); err != nil {
		return nil, err
	}

	if arrstr, err = createDailyNames(obj); err != nil {
		return nil, err
	}

	return arrstr, err
}

//______________________________________________________________________________

// setCommands it sets the default values if not supplied and checks the mandatory ones
func setCommands(obj *UserReq) error {
	var err error

	if obj.Lang == "" {
		obj.Lang = "en_US"
	}

	err = validateCommands(obj)

	return err
}

//______________________________________________________________________________

func validateCommands(obj *UserReq) error {
	var err error

	if err = checkPeriod(obj); err != nil {
		return err
	}

	if err = checkLang(obj); err != nil {
		return err
	}

	if err = checkSubFolder(obj); err != nil {
		return err
	}

	return err
}

//______________________________________________________________________________

func checkLang(obj *UserReq) error {
	if _, ok := defaultDaysName[locale(obj.Lang)]; !ok {
		return errors.New("Language not defined")
	}

	return nil
}

//______________________________________________________________________________

func checkSubFolder(obj *UserReq) error {
	if obj.PathSep == "" {
		obj.PathSep = string(os.PathSeparator)
	}

	return nil
}

//______________________________________________________________________________

func checkPeriod(obj *UserReq) error {
	var dateStart, dateEnd time.Time
	var err error

	start := obj.DateFrom
	end := obj.DateTo

	ms := "Wrong start date, bye"
	me := "No valid end date, bye"

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
				dateEnd = dateStart.AddDate(0, 0, obj.Duration-1)
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
		return errors.New("Starting date is after end date, bye")
	}

	de := dateStart.AddDate(1, 0, 0)
	if de.Before(dateEnd) {
		return errors.New("No more than a year is allowed, bye")
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

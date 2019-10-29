//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

import (
	"fmt"
	"strings"
)

//______________________________________________________________________________

// createDailyNames crates daily names according to user request
func createDailyNames(obj *UserReq) ([]string, error) {
	var jd, dn, mn string
	var arrstr []string

	dateStart, _ := checkStringDate(obj.DateFrom)
	dateEnd, _ := checkStringDate(obj.DateTo)

	i := dateStart
	yf := "2006"
	mf := "01"
	df := "02"
	cn := deleteEmpty(strings.Split(obj.Content, ","))

	for i.Before(dateEnd) || i.Equal(dateEnd) {
		var names []string

		if obj.DoY {
			jd = fmt.Sprintf("%03d", i.YearDay())
		}

		dn = defaultDaysName[locale(obj.Lang)][nameformat(obj.DoW)][i.Weekday()]
		mn = defaultMonthsName[locale(obj.Lang)][nameformat(obj.DoM)][i.Month()]

		if obj.LoneOrSub {
			// create a single name for a date
			s := []string{obj.Prefix, i.Format("2006-01-02"), jd, dn, obj.Suffix}
			s1 := strings.Join(deleteEmpty(s), "_")

			names = []string{i.Format(yf), s1}
		} else {
			// create a single name for a month and subfolders for its days
			n := []string{i.Format(mf), mn} // Month format (01...12)_(Jen...Dec)
			n1 := strings.Join(deleteEmpty(n), "_")

			s := []string{obj.Prefix, jd, i.Format(df), dn, obj.Suffix} // Day format (01...31)_(Mon...Sun)
			s1 := strings.Join(deleteEmpty(s), "_")

			names = []string{i.Format(yf), n1, s1}
		}

		if len(cn) > 0 {
			for _, el := range cn {
				arrstr = append(arrstr, strings.Join(append(names, el), obj.PathSep))
			}
		} else {
			arrstr = append(arrstr, strings.Join(names, obj.PathSep))
		}

		i = i.AddDate(0, 0, 1)
	}

	return arrstr, nil
}

//______________________________________________________________________________

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

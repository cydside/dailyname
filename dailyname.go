//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

import (
	"fmt"
	"strings"
)

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
			s := []string{obj.Prefix, i.Format("2006-01-02"), jd, dn, obj.Suffix}
			s1 := strings.Join(deleteEmpty(s), "_")

			names = []string{i.Format(yf), s1}
		} else {
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

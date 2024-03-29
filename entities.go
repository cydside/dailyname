//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

import (
	"time"
)

//______________________________________________________________________________

// UserReq descrive gli
type UserReq struct {
	Lang      string `json:"lang"`          // Language selection: Italian('it_IT') and English('en_US' default)
	PathSep   string `json:"pathsep"`       // Path separator for the desired OS
	DateFrom  string `json:"datefrom"`      // The starting date in format YYYY-MM-DD(eg.: 2019-11-08 or 2019-11 or 2019) on missing day or month will be selected the first avaiable. If the final date won't be supplied the last day of the year or month will be the selected
	DateTo    string `json:"dateto"`        // The final date in format YYYY-MM-DD(eg.: 2020-01-16 or 2020-01) on missing day will be selected the last avaiable
	Suffix    string `json:"suffix"`        // Add suffix to day folder's name
	Prefix    string `json:"prefix"`        // Add prefix to day folder's name
	LoneOrSub bool   `json:"loneorsub"`     // True: A long forlder name per day, eg: 2017-01-22 (default); False: a forlder for the year, subfolders for the months and subfolders for the days.
	Duration  int    `json:"duration"`      // Duration in number of the days(1 to 366)
	Content   string `json:"content"`       // Last names to add to daily name separated by comma
	DoW       int    `json:"dayoftheweek"`  // Day's name added: Monday(1=long format), Mon(0=short format)
	DoM       int    `json:"dayofthemonth"` // Month's name added: Jenuary(1=long format), Jen(0=short format)
	DoY       bool   `json:"dayoftheyear"`  // Julian date added: 001 to 365/366
}

//______________________________________________________________________________

// locale define language codes(cat /usr/share/i18n/SUPPORTED)
type locale string

//______________________________________________________________________________

// nameformat define the format type to use for folder's name building.
type nameformat int

//______________________________________________________________________________

// daysname Define a type
type daysname map[locale]map[nameformat]map[time.Weekday]string

//______________________________________________________________________________

// monthsname Define a type
type monthsname map[locale]map[nameformat]map[time.Month]string

//______________________________________________________________________________

const (
	// short indicates a three-letter day's name abbreviation.
	short nameformat = iota

	// long indicates a day's name in the original format.
	long
)

//______________________________________________________________________________

var dateStart, dateEnd time.Time

//______________________________________________________________________________

var defaultDaysName = daysname{
	"en_US": {
		short: {
			time.Sunday:    "Sun",
			time.Monday:    "Mon",
			time.Tuesday:   "Tue",
			time.Wednesday: "Wed",
			time.Thursday:  "Thu",
			time.Friday:    "Fri",
			time.Saturday:  "Sat",
		},
		long: {
			time.Sunday:    "Sunday",
			time.Monday:    "Monday",
			time.Tuesday:   "Tuesday",
			time.Wednesday: "Wednesday",
			time.Thursday:  "Thursday",
			time.Friday:    "Friday",
			time.Saturday:  "Saturday",
		},
	},
	"it_IT": {
		short: {
			time.Sunday:    "Dom",
			time.Monday:    "Lun",
			time.Tuesday:   "Mar",
			time.Wednesday: "Mer",
			time.Thursday:  "Gio",
			time.Friday:    "Ven",
			time.Saturday:  "Sab",
		},
		long: {
			time.Sunday:    "Domenica",
			time.Monday:    "Lunedì",
			time.Tuesday:   "Martedì",
			time.Wednesday: "Mercoledì",
			time.Thursday:  "Giovedì",
			time.Friday:    "Venerdì",
			time.Saturday:  "Sabato",
		},
	},
}

var defaultMonthsName = monthsname{
	"en_US": {
		short: {
			time.January:   "Jan",
			time.February:  "Feb",
			time.March:     "Mar",
			time.April:     "Apr",
			time.May:       "May",
			time.June:      "Jun",
			time.July:      "Jul",
			time.August:    "Aug",
			time.September: "Sep",
			time.October:   "Oct",
			time.November:  "Nov",
			time.December:  "Dec",
		},
		long: {
			time.January:   "January",
			time.February:  "February",
			time.March:     "March",
			time.April:     "April",
			time.May:       "May",
			time.June:      "June",
			time.July:      "July",
			time.August:    "August",
			time.September: "September",
			time.October:   "October",
			time.November:  "November",
			time.December:  "December",
		},
	},
	"it_IT": {
		short: {
			time.January:   "Gen",
			time.February:  "Feb",
			time.March:     "Mar",
			time.April:     "Apr",
			time.May:       "Mag",
			time.June:      "Giu",
			time.July:      "Lug",
			time.August:    "Aug",
			time.September: "Set",
			time.October:   "Ott",
			time.November:  "Nov",
			time.December:  "Dic",
		},
		long: {
			time.January:   "Gennaio",
			time.February:  "Febbraio",
			time.March:     "Marzo",
			time.April:     "Aprile",
			time.May:       "Maggio",
			time.June:      "Giugno",
			time.July:      "Luglio",
			time.August:    "Agosto",
			time.September: "Settembre",
			time.October:   "Ottobre",
			time.November:  "Novembre",
			time.December:  "Dicembre",
		},
	},
}

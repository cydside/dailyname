//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

import (
	"time"
)

//______________________________________________________________________________

// UserReq descrive gli
type UserReq struct {
	Lang      string `json:"lang"`          // Language selected(Italian('it_IT' default) and English('en_US') are pre-loaded).
	PathSep   string `json:"pathsep"`       // Path separator for the desired OS.
	DateFrom  string `json:"datefrom"`      // The starting date in format YYYYMMDD(eg.: 2019-11-08 or 2019-11 or 2019) on missing data will be selected the first avaiable. If the final date won't be supplied the last day of the year will be the selected.
	DateTo    string `json:"dateto"`        // The final datein format YYYYMMDD(eg.: 2020-01-16 or  2020-01 or  2020) on missing data will be selected the last avaiable without crossing year, if the starting date won't be supplied the first day of the year will be the selected.
	Suffix    string `json:"suffix"`        // Add suffix to day folder's name.
	Prefix    string `json:"prefix"`        // Add prefix to day folder's name.
	LoneOrSub bool   `json:"loneorsub"`     // True: A long forlder name per day(eg: 2017-01-22); False: a forlder for the year, subfolders for the months and subfolders for the days(default).
	Duration  int    `json:"duration"`      // Duration in number of the days(1 to 366)
	DoW       int    `json:"dayoftheweek"`  // Day's name added: Monday(long format), Mon(short format)
	DoM       int    `json:"dayofthemonth"` // Month's name added: Jenuary(long format), Jen(short format)
	DoY       bool   `json:"dayoftheyear"`  // Julian date added: 001 for jenuary first
}

//______________________________________________________________________________

// locale define language codes(cat /usr/share/i18n/SUPPORTED)
type locale string

//______________________________________________________________________________

// nameformat define the format type to use for folder's name building.
type nameformat uint8

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

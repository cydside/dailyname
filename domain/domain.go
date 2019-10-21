//------------------------------------------------------------------------------

package domain

//------------------------------------------------------------------------------

import (
	// "errors"
	// "flag"
	// "fmt"
	// "io/ioutil"
	// "os"
	// "path/filepath"
	// "strconv"
	// "strings"
	"time"
)

//------------------------------------------------------------------------------

// Locale define language codes as in /usr/share/i18n/SUPPORTED
type Locale string

// NameFormat define the format type to use for folder's name building.
type NameFormat uint8

// DaysName Define a type
type DaysName map[Locale]map[NameFormat]map[time.Weekday]string

// MonthsName Define a type
type MonthsName map[Locale]map[NameFormat]map[time.Month]string

const (
	// Short indicates a three-letter day's name abbreviation.
	Short NameFormat = iota

	// Long indicates a day's name in the original format.
	Long
)

var defaultDaysName = DaysName{
	"en_US": {
		Short: {
			time.Sunday:    "Sun",
			time.Monday:    "Mon",
			time.Tuesday:   "Tue",
			time.Wednesday: "Wed",
			time.Thursday:  "Thu",
			time.Friday:    "Fri",
			time.Saturday:  "Sat",
		},
		Long: {
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
		Short: {
			time.Sunday:    "Dom",
			time.Monday:    "Lun",
			time.Tuesday:   "Mar",
			time.Wednesday: "Mer",
			time.Thursday:  "Gio",
			time.Friday:    "Ven",
			time.Saturday:  "Sab",
		},
		Long: {
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

var defaultMonthsName = MonthsName{
	"en_US": {
		Short: {
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
		Long: {
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
		Short: {
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
		Long: {
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

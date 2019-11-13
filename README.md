# dailyname
Package to create folders' name path. Take a look to [dailyfolders](https://github.com/cydside/dayfolders) to see it in action.

## Install

```
go get github.com/cydside/dailyname
```

## Usage
Set the options according to the following struct:

```
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
```

and get an array of strings just calling `GetDailyNames`.

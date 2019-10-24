//______________________________________________________________________________

package usecases

//______________________________________________________________________________

// UserReqLay interfaccia
type UserReqLay interface {
	Create(obj *UserReq) ([]string, error)
}

//______________________________________________________________________________

// UserReq descrive gli
type UserReq struct {
	ID        string `json:"id"`
	Lang      string `json:"lang"`          // Language selected(Italian('it_IT' default) and English('en_US') are pre-loaded).
	DateFrom  string `json:"datefrom"`      // The starting date in format YYYYMMDD(eg.: 2019-11-08 or 2019-11 or 2019) on missing data will be selected the first avaiable. If the final date won't be supplied the last day of the year will be the selected.
	DateTo    string `json:"dateto"`        // The final datein format YYYYMMDD(eg.: 2020-01-16 or  2020-01 or  2020) on missing data will be selected the last avaiable without crossing year, if the starting date won't be supplied the first day of the year will be the selected.
	Suffix    string `json:"suffix"`        // Add suffix to day folder's name.
	Prefix    string `json:"prefix"`        // Add prefix to day folder's name.
	LoneOrSub bool   `json:"loneorsub"`     // True: A long forlder name per day(eg: 2017-01-22); False: a forlder for the year, subfolders for the months and subfolders for the days(default).
	Duration  uint16 `json:"duration"`      // Duration in number of the days(1 to 366)
	DoW       uint8  `json:"dayoftheweek"`  // Day's name added: Monday(Long format), Mon(Short format)
	DoM       uint8  `json:"dayofthemonth"` // Month's name added: Jenuary(Long format), Jen(Short format)
	DoY       bool   `json:"dayoftheyear"`  // Julian date added: 001 for jenuary first
}

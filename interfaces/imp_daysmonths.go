package interfaces

//______________________________________________________________________________

import (
	// "log"

	"github.com/cydside/dailyname/domain"
)

//______________________________________________________________________________

// DaysNameImp implementazione dell'interfaccia DaysNameLay
type DaysNameImp struct {
}

//______________________________________________________________________________

// NewDaysNameImp new
func NewDaysNameImp() *DaysNameImp {
	obj := new(DaysNameImp)

	return obj
}

//______________________________________________________________________________

// Create implementa
func (p *DaysNameImp) Create(obj *domain.DaysName) error {
	var err error

	return err
}

//______________________________________________________________________________

// Save implementa
func (p *DaysNameImp) Save(key string, obj *domain.DaysName) error {
	var err error

	return err
}

//______________________________________________________________________________

// Fetch implementa
func (p *DaysNameImp) Fetch() []domain.DaysName {
	var objarr []domain.DaysName

	return objarr
}

//______________________________________________________________________________

// FetchByKey implementa
func (p *DaysNameImp) FetchByKey(key string) domain.DaysName {
	var obj domain.DaysName

	return obj
}

//______________________________________________________________________________

// EraseByKey implementa
func (p *DaysNameImp) EraseByKey(key string) error {
	var err error

	return err
}

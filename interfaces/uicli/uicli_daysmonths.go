package uicli

//______________________________________________________________________________

import (
	// "log"

	"github.com/cydside/dailyname/domain"
)

//______________________________________________________________________________

// MonthsDaysUscLay implementa MonthsDaysUscImp
type MonthsDaysUscLay interface {
	Add(objMonths *domain.MonthsName, objDays *domain.DaysName) error
	// Edit(item *domain.Account) error
	// Get(key string) domain.Account
	// GetAll() []domain.Account
	// Find(text string) []domain.Account
	// Remove(key string) error
}

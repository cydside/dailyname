//______________________________________________________________________________

package usecases

//______________________________________________________________________________

import (
	// "log"

	// "github.com/pkg/errors"

	"github.com/cydside/dailyname/domain"
)

//______________________________________________________________________________

// MonthsDaysUscImp implementa MonthsDaysUscLay
type MonthsDaysUscImp struct {
	DaysNameLay   domain.DaysNameLay
	MonthsNameLay domain.MonthsNameLay
}

//______________________________________________________________________________

// Add Aggiunge un nuovo set di nomi e mesi dei giorni
func (p *MonthsDaysUscImp) Add(objMonths *domain.MonthsName, objDays *domain.DaysName) error {
	// var m string
	// var err error

	// if err = p.ruleEmpty(item); err != nil {
	// 	return err
	// }

	// if err = p.chechIfKeyDuplicate(item.Tipo); err != nil {
	// 	return err
	// }

	// if err = p.ruleDecimal(item); err != nil {
	// 	return err
	// }

	// if err = p.ServicesArcLay.Create(item); err != nil {
	// 	m = "Errore nel salvataggio dei dati"
	// 	err = errors.Wrap(err, m)

	// 	return err
	// }

	return nil
}

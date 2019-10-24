//______________________________________________________________________________

package usecases

//______________________________________________________________________________

import (
// "log"

// "github.com/pkg/errors"
// "github.com/cydside/dailyname/domain"
)

//______________________________________________________________________________

// UserUscImp implementa
type UserUscImp struct {
	UserReqLay UserReqLay
}

//______________________________________________________________________________

// Get Utente
func (p *UserUscImp) Get(obj *UserReq) ([]string, error) {
	var objarr []string
	var err error

	// TODO: effettuare i controlli sulla richiesta

	objarr, err = p.UserReqLay.Create(obj)

	return objarr, err
}

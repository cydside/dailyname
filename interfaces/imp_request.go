package interfaces

//______________________________________________________________________________

import (
	// "log"

	"github.com/cydside/dailyname/usecases"
)

//______________________________________________________________________________

// UserReqImp implementazione dell'interfaccia UserReqLay
type UserReqImp struct {
}

//______________________________________________________________________________

// NewUserReqImp new
func NewUserReqImp() *UserReqImp {
	obj := new(UserReqImp)

	return obj
}

//______________________________________________________________________________

// Create implementa
func (p *UserReqImp) Create(obj *usecases.UserReq) ([]string, error) {
	var objarr []string
	var err error

	return objarr, err
}

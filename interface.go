//______________________________________________________________________________

package dailyname

//______________________________________________________________________________

// SetCommands imposta le opzioni disponibili per la crazione dei nomi
func SetCommands(obj *UserReq) error {
	var err error

	err = validateCommands(obj)

	return err
}

//______________________________________________________________________________

func validateCommands(obj *UserReq) error {
	var err error

	return err
}

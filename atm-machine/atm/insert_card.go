package atm

type InsertCard struct {
	atm *ATM
	ATMAbstract
}

func (d *InsertCard) InsertCard() error {
	d.atm.SetState(d.atm.readCard)
	return nil
}

func (d *InsertCard) StateName() string {
	return "InsertCard"
}

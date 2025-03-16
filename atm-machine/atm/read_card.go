package atm

type ReadCard struct {
	atm *ATM
	ATMAbstract
}

func (d *ReadCard) GetCardDetail() error {
	card := NewCardReader().ReadCard("valid")
	if card.CardNo == "" {
		d.atm.SetState(d.atm.insertCard)
	}
	d.atm.card = card
	d.atm.SetState(d.atm.selectAccount)
	return nil
}

func (d *ReadCard) StateName() string {
	return "ReadCard"
}

package atm

type Withdraw interface {
	ProcessAmount(atm *ATM, atmamount float64)
}

func NewWithDrawPipeline() Withdraw {
	return fiveHundreadWithdraw{
		next: hundreadHundreadWithdraw{},
	}
}

type fiveHundreadWithdraw struct {
	next Withdraw
}

func (f fiveHundreadWithdraw) ProcessAmount(atm *ATM, atmamount float64) {
	div := atmamount / 500
	rem := int(atmamount) % 500
	atm.WithdrawAs.FiveHundread = int(div)
	f.next.ProcessAmount(atm, float64(rem))
}

type hundreadHundreadWithdraw struct {
	next Withdraw
}

func (h hundreadHundreadWithdraw) ProcessAmount(atm *ATM, atmamount float64) {
	div := atmamount / 100
	rem := int(atmamount) % 100
	atm.WithdrawAs.Hundread = int(div)
	atm.WithdrawAs.Left = rem

}

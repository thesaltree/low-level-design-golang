package atm

type Withdraw interface {
	ProcessAmount(atm *ATM, atmamount float64)
}

func NewWithDrawPipeline() Withdraw {
	return fiveHundredWithdraw{
		next: oneHundredWithdraw{},
	}
}

type fiveHundredWithdraw struct {
	next Withdraw
}

func (f fiveHundredWithdraw) ProcessAmount(atm *ATM, atmamount float64) {
	div := atmamount / 500
	rem := int(atmamount) % 500
	atm.WithdrawAs.FiveHundred = int(div)
	f.next.ProcessAmount(atm, float64(rem))
}

type oneHundredWithdraw struct {
	next Withdraw
}

func (h oneHundredWithdraw) ProcessAmount(atm *ATM, atmamount float64) {
	div := atmamount / 100
	rem := int(atmamount) % 100
	atm.WithdrawAs.Hundred = int(div)
	atm.WithdrawAs.Left = rem

}

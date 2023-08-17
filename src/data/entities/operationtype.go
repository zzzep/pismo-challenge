package entities

type OperationType int

const (
	Purchase OperationType = iota + 1
	Installment
	Withdrawal
	Payment
)

func (o OperationType) String() string {
	switch o {
	case Purchase:
		return "COMPRA A VISTA"
	case Installment:
		return "COMPRA PARCELADA"
	case Withdrawal:
		return "SAQUE"
	case Payment:
		return "PAGAMENTO"
	}
	return "unknown"
}

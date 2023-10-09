package entities

type OperationType int

const (
	Purchase OperationType = iota + 1
	Installment
	Withdrawal
	Payment
)

// String returns a string representation of the OperationType.
//
// It returns the corresponding string representation for each OperationType constant.
// The possible values are:
// - Purchase: "COMPRA A VISTA"
// - Installment: "COMPRA PARCELADA"
// - Withdrawal: "SAQUE"
// - Payment: "PAGAMENTO"
// If the OperationType is not recognized, it returns "unknown".
// The returned string can be used for display or serialization purposes.
//
// Return:
// - string: The string representation of the OperationType.
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

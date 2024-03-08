package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSuportedCurrency returns true if the currency is suported
func IsSuportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}

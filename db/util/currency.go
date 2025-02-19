package util

const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	GBP = "GBP"
	AUD = "AUD"
	CAD = "CAD"
	RMB = "RMB"
)

func IsSupportedCurrency(currency string) bool {
	// Check if the currency is supported
	switch currency {
	case USD, EUR, JPY, GBP, AUD, CAD, RMB:
		return true
	}
	return false
}

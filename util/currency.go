package util

// Constantes para todo tipo de monedas
const (
	USD = "USD"
	EUR = "EUR"
	COP = "COP"
)

// IsSupportedCurrency devuelve true si la moneda es soportada
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, COP:
		return true
	}
	return false
}

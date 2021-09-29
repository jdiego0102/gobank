package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz" // Sin la letra ñ, falta configurar codificación UTF8

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt genera un entero aleatorio entre el mínimo y el máximo
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min->max
}

// RandomString genera una cadena alatoria de la longitud n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner genera un nombre propierario aleatorio
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney genera un monto de dinero aleatorio
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency genera un código de cambio de divisas aleatorio.
func RandomCurrency() string {
	currencies := []string{"COP", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

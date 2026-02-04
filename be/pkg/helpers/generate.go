package helpers

import (
	"math/rand"
	"rialfu/wallet/pkg/constants"
)

func GenerateRandomString(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = constants.Charset[rand.Intn(len(constants.Charset))]
	}
	return string(b)
}

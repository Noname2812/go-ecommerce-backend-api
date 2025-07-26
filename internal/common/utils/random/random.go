package random

import (
	"encoding/hex"
	"time"

	crypto "crypto/rand"
	"math/rand"
)

func GenerateSixDigitOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000) // 100000 - 999999
	return otp
}

func GenarateToken(length int) (string, error) {
	token := make([]byte, length)
	if _, err := crypto.Read(token); err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

package utils

import (
	"encoding/hex"
	"fmt"

	crypto "github.com/CityOfZion/neo-go/pkg/crypto"
)

func ToHash160(address string) (string, error) {
	u, err := crypto.Uint160DecodeAddress(address)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	rs := hex.EncodeToString(u.BytesReverse())
	return rs, nil
}

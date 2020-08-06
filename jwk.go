package main

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"

	"github.com/lestrrat-go/jwx/jwk"
)

func PrinkJwk(pubKey rsa.PublicKey) {
	jwkKey, err := jwk.New(pubKey)
	if err != nil {
		fmt.Println(err)
	}
	jwkKey.Set(jwk.AlgorithmKey, "RS256")
	jwkKey.Set(jwk.KeyUsageKey, "sig")
	jwkKey.Set(jwk.KeyIDKey, "asdf")
	fmt.Println(jwkKey)
	jsonbuf, err := json.MarshalIndent(jwkKey, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonbuf))
	thumbprint, _ := jwkKey.Thumbprint(crypto.SHA256)
	fmt.Println(string(thumbprint))
}

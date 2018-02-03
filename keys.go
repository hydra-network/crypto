package crypto

import (
    "os"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/x509"
    "fmt"
    "math/big"
    "encoding/hex"
)

func Keys() (*big.Int, string) {
 	pubkeyCurve := elliptic.P256()

 	privatekey := new(ecdsa.PrivateKey)
 	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	var pubkey ecdsa.PublicKey
 	pubkey = privatekey.PublicKey

    pubASN1, err := x509.MarshalPKIXPublicKey(&pubkey)
    if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
    }

    publickey := hex.EncodeToString(pubASN1)

    return privatekey.D, publickey
 }

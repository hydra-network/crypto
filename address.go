package crypto

import (
    "crypto/elliptic"
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/x509"
    "math/big"
    "encoding/hex"
)

type Address struct {
    Privatekey   *big.Int
    Publickey    string
}

func GenerateAddress() (Address, error) {
    var newAddress Address

    pubkeyCurve := elliptic.P256()

    privatekey := new(ecdsa.PrivateKey)
    privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
    if err != nil {
        return newAddress, err
    }

    var pubkey ecdsa.PublicKey
    pubkey = privatekey.PublicKey

    pubASN1, err := x509.MarshalPKIXPublicKey(&pubkey)
    if err != nil {
        return newAddress, err
    }

    publickey := hex.EncodeToString(pubASN1)

    newAddress.Privatekey = privatekey.D
    newAddress.Publickey = publickey

    return  newAddress, nil
 }

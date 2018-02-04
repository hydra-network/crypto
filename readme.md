# Crypto-functions

## Example

How to generate a private and public keys:

```
package main

import (
    "github.com/montana-network/crypto"
    "fmt"
)

func main() {
    address, addrErr := crypto.GenerateAddress()

    if addrErr == nil {
        fmt.Println("Private Key :")
        fmt.Printf("%x \n", address.Privatekey)

        fmt.Println("Public Key :")
        fmt.Printf("%x \n", address.Publickey)
    }
}

```

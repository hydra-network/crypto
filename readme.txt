## Crypto-functions

# Example

Generate a private and public keys

```
package main

import (
    "github.com/montana-network/crypto"
)

privatekey, publickey := crypto.Keys()
fmt.Println("Private Key :")
fmt.Printf("%x \n", privatekey)

fmt.Println("Public Key :")
fmt.Printf("%x \n", publickey)
```
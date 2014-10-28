package main

import (
  "fmt"
  "github.com/gillesdemey/go-rsa-example/crypto"
)

func main() {

  kp, err := crypto.CreateKeyPair();

  if err != nil {
    fmt.Printf("%v", err)
  }

  msg := []byte("Hello, NSA!")

  c, err := crypto.Encrypt(msg, kp.PublicKey)

  if err != nil {
    fmt.Printf("%v", err)
  }

  fmt.Printf("encrypted: %v\n", c)

  m, err := crypto.Decrypt(c, kp)

  if err != nil {
    fmt.Printf("%v", err)
  }

  fmt.Printf("decrypted: %s\n", m)
}


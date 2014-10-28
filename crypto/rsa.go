package crypto

import (
  "log"
  "crypto/sha1"
  "crypto/rand"
  "crypto/rsa"
)

func CreateKeyPair() (*rsa.PrivateKey, error) {

  size := 1024

  priv, err := rsa.GenerateKey(rand.Reader, size)

  if err != nil {
    log.Fatalf("Failed to generate %d-bit key", size)
    return nil, err
  }

  return priv, err

}

func Encrypt(in []byte, pub rsa.PublicKey) ([]byte, error) {

  sha1 := sha1.New()

  out, err := rsa.EncryptOAEP(sha1, rand.Reader, &pub, in, nil)

  if err != nil {
    log.Fatalf("Failed to encrypt message %v", err)
    return nil, err
  }

  return out, nil
}

func Decrypt(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {

  sha1 := sha1.New()

  out, err := rsa.DecryptOAEP(sha1, rand.Reader, priv, ciphertext, nil)

  if err != nil {
    log.Fatalf("Failed to decrypt message %v", err)
    return nil, err
  }

  return out, nil
}
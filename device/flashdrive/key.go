package flashdrive

import "golang.org/x/crypto/argon2"

func deriveKey(password, salt []byte) []byte {
  return argon2.IDKey(password, salt, 3, 64*1024, 4, 32)
}

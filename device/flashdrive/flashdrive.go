package flashdrive

import (
  "fmt"
  "path/filepath"
)

type FlashDrive struct {
  Path string
  Name string
}

func NewFlashDrive(path string) *FlashDrive {
  return &FlashDrive {
    Path: path,
    Name: filepath.Base(path),
  }
}

func (f *FlashDrive) Info() string {
  return fmt.Sprintf("Drive: %s at %s", f.Name, f.Path)
}

func (f *FlashDrive) Encrypt(crypt string) error {
  return f.EncryptContainer([]byte(crypt))
}

func (f *FlashDrive) Decrypt(crypt string) error {
  return f.DecryptContainer([]byte(crypt))
}

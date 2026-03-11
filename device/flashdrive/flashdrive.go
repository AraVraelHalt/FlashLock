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

func (f *FlashDrive) Encrypt() error {
  //TODO: implement encryption logic
  return fmt.Errorf("Encryption not implemeneted yet.")
}

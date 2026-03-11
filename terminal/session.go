package terminal

import (
  "flashlock/device/flashdrive"
)

type Session struct {
  SelectedDevice *flashdrive.FlashDrive
}

func NewSession() *Session {
  return &Session{}
}

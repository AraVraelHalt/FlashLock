package device

import "runtime"

type DeviceScanner interface {
  Scan() []string
}

func NewScanner() DeviceScanner {
  switch runtime.GOOS {
    case "windows":
      return &WindowsScanner{}
    case "darwin":
      return &MacScanner{}
    case "linux":
      return &LinuxScanner{}
    default:
      return nil
  }
}

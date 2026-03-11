package device

import (
  "runtime"
  "fmt"
  "path/filepath"
  "flashlock/device/flashdrive"
)

var FoundDevicesPaths []string

type Device interface {
  Scan() []string
  Eject(path string) error
}

func NewDevice() Device {
  switch runtime.GOOS {
    case "windows":
      return &Windows{}
    case "darwin":
      return &Mac{}
    case "linux":
      return &Linux{}
    default:
      return nil
  }
}

func ScanForDevices() {
  drives := NewDevice()
  FoundDevicesPaths = drives.Scan()

  if len(FoundDevicesPaths) == 0 {
    fmt.Println("\nNo drives detected.\n")
  } else {
      fmt.Println("\nDetected drives:")

      for i, d := range FoundDevicesPaths {
        name := filepath.Base(d)
        fmt.Println(i+1, "-", name)
      }

      fmt.Println()
  }
}

func SelectDevice(index int) (*flashdrive.FlashDrive, error) { 
  if !IsDeviceIndexInRange(index) {
    return nil, fmt.Errorf("Not a valid option in found devices") 
  }
  
  path := FoundDevicesPaths[index]

  return flashdrive.NewFlashDrive(path), nil
}

func ContainsDevices() bool {
  return len(FoundDevicesPaths) > 0
}

func IsDeviceIndexInRange(index int) bool {
  return index >= 0 && index < len(FoundDevicesPaths) 
}

func EjectDevice(path string) {
  drives := NewDevice()

  drives.Eject(path)
}

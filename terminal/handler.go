package terminal

import (
  "fmt"
  "strings"
  "strconv"
  "flashlock/device"
)

func ValidateSelect(input string) bool {  
  status := false
  parts := strings.Split(input, " ")

  if len(parts) < 2 {
    fmt.Println("\nUsage 'select <index>'\n")
    return status
  }
  
  if len(parts) > 2 {
    fmt.Println("\nToo many arguments for 'select <index>'\n")
    return status
  }

  index, err := strconv.Atoi(parts[1])
  if err != nil {
    fmt.Println("\nIndex must be a number\n")
    return status
  }
 
  if !device.ContainsDevices() {
    fmt.Println("\nRun 'scan' and see options first\n")
    return status
  }

  if !device.IsDeviceIndexInRange(index) {
    fmt.Println("\nNot a valid option in found devices\n")
    return status
  }  
  
  status = true
  return status
}

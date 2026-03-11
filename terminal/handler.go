package terminal

import (
  "fmt"
  "strings"
  "strconv"
  "flashlock/device"
  "flashlock/device/flashdrive"
)

func ValidateSelect(input string) (int, bool) {  
  invalid := -1
  status := false
  parts := strings.Split(input, " ")

  if len(parts) < 2 {
    fmt.Println("\nUsage 'select <index>'\n")
    return invalid, status
  }
  
  if len(parts) > 2 {
    fmt.Println("\nToo many arguments for 'select <index>'\n")
    return invalid, status
  }

  index, err := strconv.Atoi(parts[1])

  if err != nil {
    fmt.Println("\nIndex must be a number\n")
    return invalid, status
  }
 
  if !device.ContainsDevices() {
    fmt.Println("\nRun 'scan' and see options first\n")
    return invalid, status
  }
  
  status = true
  return index, status
}

func SelectDevice(input string) (*flashdrive.FlashDrive) {
  index, valid := ValidateSelect(input)

  if !valid {
    return nil
  }

  drive, err := device.SelectDevice(index-1) 

  if err != nil {
    fmt.Println("\n", err, "\n")
    return nil
  } 

  return drive
}

func ValidateCryption(input string) (string, bool) {
  parts := strings.Split(input, " ")
  invalid := ""

  if len(parts) < 2 {
    fmt.Println("\nUsage '",parts[0]," <psswd>'\n")
    return invalid, false 
  }

  if len(parts) > 2 {
    fmt.Println("\nToo many arguments for '",parts[0]," <psswd>'\n")
    return invalid, false
  }

  return parts[1], true
}

func PrintHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  scan             - List all connected drives")
  fmt.Println("  select <index>   - Selects a flashdrive from list") 
	fmt.Println("  encrypt          - Encrypts selected flash drive")
	fmt.Println("  decrypt          - Decrypts selected flash drive")
	fmt.Println("  eject            - Safely ejects drive before removing")
  fmt.Println("  clear            - Clean slate for terminal")
	fmt.Println("  help             - Shows list of commands")
  fmt.Println("  exit/quit        - Escape out of program\n")
}

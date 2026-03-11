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
    FormatPrintln("Usage 'select <index>'")
    return invalid, status
  }
  
  if len(parts) > 2 {
    FormatPrintln("Too many arguments for 'select <index'")
    return invalid, status
  }

  index, err := strconv.Atoi(parts[1])

  if err != nil {
    FormatPrintln("Index must be a number")
    return invalid, status
  }
 
  if !device.ContainsDevices() {
    FormatPrintln("Run 'scan' and see options first")
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
    FormatPrintln(err.Error())
    return nil
  } 

  return drive
}

func ValidateCryption(input string) (string, bool) {
  parts := strings.Split(input, " ")
  invalid := ""

  if len(parts) < 2 {
    fmt.Println()
    fmt.Println("Usage '",parts[0]," <psswd>'")
    fmt.Println()
    return invalid, false 
  }

  if len(parts) > 2 {
    fmt.Println()
    fmt.Println("Too many arguments for '",parts[0]," <psswd>'")
    fmt.Println()
    return invalid, false
  }

  return parts[1], true
}

func PrintHelp() {
  fmt.Println(`
Available commands:
  scan             - List all connected drives
  select <index>   - Selects a flashdrive from list 
  encrypt          - Encrypts selected flash drive
  decrypt          - Decrypts selected flash drive
  eject            - Safely ejects drive before removing
  clear            - Clean slate for terminal
  help             - Shows list of commands
  exit/quit        - Escape out of program
  `)
}

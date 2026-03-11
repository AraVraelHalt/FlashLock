package terminal

import ( 
  "bufio"
  "fmt"
  "os"
  "strings"
  "flashlock/device"
)

func Listen() {
  reader := bufio.NewReader(os.Stdin)
  session := NewSession()

  for {
    if session.SelectedDevice != nil {
      fmt.Print(session.SelectedDevice.Name, " > ")
    } else {
      fmt.Print("> ")
    }
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    switch {
      case input=="quit" || input=="exit":
        FormatPrintln("Exiting...")
        return
      case input=="clear":
        fmt.Print("\033[H\033[2J")
      case input=="scan":
        device.ScanForDevices()
      case strings.HasPrefix(input, "select "):
        session.SelectedDevice = SelectDevice(input)
        
        if session.SelectedDevice != nil {
          fmt.Println()
          fmt.Println("Selected: ",session.SelectedDevice.Info())
          fmt.Println()
        }
      case strings.HasPrefix(input, "encrypt "):
        if session.SelectedDevice != nil {
          crypt, valid := ValidateCryption(input)

          if valid {
            err := session.SelectedDevice.Encrypt(crypt)

            if err != nil {
              FormatPrintln(err.Error())
            } else {
              FormatPrintln("Successfully encrypted drive...")
            }
          }
        } else {
          FormatPrintln("Please select a device first")
        }
      case strings.HasPrefix(input, "decrypt "):
        if session.SelectedDevice != nil {
          crypt, valid := ValidateCryption(input)
  
          if valid {
            err := session.SelectedDevice.Decrypt(crypt)
  
            if err != nil {
              FormatPrintln(err.Error())
            } else {
              FormatPrintln("Successfully decrypted drive...")
            }
          }
        } else {
          FormatPrintln("Please select a device first")
        }
      case input=="eject":
        if session.SelectedDevice != nil {
          device.EjectDevice(session.SelectedDevice.Path)
          session.SelectedDevice = nil

          FormatPrintln("Drive ejected...")
        } else {
          FormatPrintln("Please select a device first")
        }
      case input=="help":
        PrintHelp()
      default:
        fmt.Println()
        fmt.Println("Unknown command:", input)
        fmt.Println()
    }
  }
}

func FormatPrintln (message string) {
  fmt.Println()
  fmt.Println(message)
  fmt.Println()
}

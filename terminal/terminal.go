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
        fmt.Println("\nExiting...\n")
        return
      case input=="clear":
        fmt.Print("\033[H\033[2J")
      case input=="scan":
        device.ScanForDevices()
      case strings.HasPrefix(input, "select"):
        session.SelectedDevice = SelectDevice(input)
        
        if session.SelectedDevice != nil {
          fmt.Println("\nSelected: ",session.SelectedDevice.Info(),"\n")
        }
      case input=="encrypt":
        if session.SelectedDevice != nil {
          session.SelectedDevice.Encrypt()
        } else {
          fmt.Println("\nPlease select a device first\n")
        }
      case input=="decrypt":
        if session.SelectedDevice != nil {
          session.SelectedDevice.Decrypt()
        } else {
          fmt.Println("\nPlease select a device first\n")
        }
      case input=="eject":
        if session.SelectedDevice != nil {
          device.EjectDevice(session.SelectedDevice.Path)
          session.SelectedDevice = nil

          fmt.Println("\nDrive ejected...\n")
        } else {
          fmt.Println("\nPlease select a device first\n")
        }
      case input=="help":
        PrintHelp()
      default:
        fmt.Println("\nUnknown command:", input, "\n")
    }
  }
}

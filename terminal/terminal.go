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

  for {
    fmt.Print("> ")
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
        ValidateSelect(input) 
      default:
        fmt.Println("\nUnknown command:", input, "\n")
    }
  }
}

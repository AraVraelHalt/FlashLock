package terminal

import ( 
  "bufio"
  "fmt"
  "os"
  "strings"
)

func Listen() {
  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("> ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    switch input {
      case "quit", "exit":
        fmt.Println("\nExiting...\n")
        return
      case "scan":
        fmt.Println("\nScanning...\n")
      default:
        fmt.Println("\nUnknown command:", input, "\n")
    }
  }
}

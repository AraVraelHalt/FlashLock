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

    switch input {
      case "quit", "exit":
        fmt.Println("\nExiting...\n")
        return
      case "clear":
        fmt.Print("\033[H\033[2J")
      case "scan":
        fmt.Println("\nScanning...\n")
        scanner := device.NewScanner()
        drives := scanner.Scan()

        if len(drives) == 0 {
          fmt.Println("No drives detected.\n")
        } else {
            fmt.Println("Detected drives:")
          
            for _, d := range drives {
              fmt.Println(" -", d)
            }
            fmt.Println()
        } 
      default:
        fmt.Println("Unknown command:", input, "\n")
    }
  }
}

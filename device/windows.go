package device

import (
  "fmt"
  "os/exec"
  "strings"
)

type WindowsScanner struct{}

func (w *WindowsScanner) Scan() []string {
  out, err := exec.Command("wmic", "logicaldisk", "get", "name").Output()

  if err != nil {
    fmt.Println("Error scanning drives:", err)
    return nil
  }

  lines := strings.Split(string(out), "\n")
  drives := []string{}

  for _, line := range lines {
    line = strings.TrimSpace(line)
    
    if len(line) > 0 && line != "Name" {
      drives = append(drives, line+"\\")
    }
  }

  return drives
}

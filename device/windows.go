package device

import (
  "fmt"
  "os/exec"
  "strings"
)

type Windows struct{}

func (w *Windows) Scan() []string {
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

func (w *Windows) Eject(driveLetter string) error {
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Remove-Item -Path %s -Recurse -Force", driveLetter))
	return cmd.Run()
}

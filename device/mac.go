package device

import (
  "io/ioutil"
  "log"
)

type MacScanner struct{}

func (m *MacScanner) Scan() []string {
  mountPath := "/Volumes/"
  entries, err := ioutil.ReadDir(mountPath)

  if err != nil {
    log.Println("Error reading mounts:", err)
    return nil
  }

  drives := []string{}
  
  for _, e := range entries {
    if e.IsDir() && e.Name() != "Macintosh HD" {
      drives = append(drives, mountPath+e.Name())
    }
  }
  
  return drives
}

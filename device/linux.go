package device

import (
  "io/ioutil"
  "log"
)

type LinuxScanner struct{}

func (l *LinuxScanner) Scan() []string {
  mountPath := "/media/"
  entries, err := ioutil.ReadDir(mountPath)

  if err != nil {
    log.Println("Error reading mounts:", err)
    return nil
  }

  drives := []string{}

  for _, e := range entries {
    if e.IsDir() {
      drives = append(drives, mountPath+e.Name())
    }
  }

  return drives
}

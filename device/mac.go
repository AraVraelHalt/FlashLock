package device

import (
  "io/ioutil"
  "os/exec"
  "log"
)

type Mac struct{}

func (m *Mac) Scan() []string {
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

func (m *Mac) Eject(mountPath string) error {
	cmd := exec.Command("diskutil", "unmount", mountPath)
	return cmd.Run()
}

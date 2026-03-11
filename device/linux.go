package device

import (
  "io/ioutil"
  "os/exec"
  "log"
)

type Linux struct{}

func (l *Linux) Scan() []string {
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

func (l *Linux) Eject(mountPath string) error {
	cmd := exec.Command("umount", mountPath)
	return cmd.Run()
}

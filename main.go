package main

import (
  "flashlock/banner"
  "flashlock/terminal"
)

func main() {
  banner.Print()
  terminal.Listen()
}

package main

import (
    "os"
    "fmt"
    "github.com/jacobsa/go-serial/serial"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Fprintf(os.Stderr, "Usage: serial-reader portname\n")
    os.Exit(1)
  }
  portname := os.Args[1]
  // /dev/ttyACM0
  options := serial.OpenOptions{
    PortName:        portname,
    BaudRate:        9600,
    DataBits:        8,
    StopBits:        1,
    MinimumReadSize: 4,
  }

  port, _ := serial.Open(options)
  for {
    buf := make([]byte, 1000)
    port.Read(buf)
    fmt.Printf("%s", string(buf))
  }
}

package main

import (
	// "fmt"
	"os"
	"syscall"
)

func main() {
	v, _ := syscall.Getenv("PROGRAM")
	f, _ := os.Create(v + "-out")
	defer f.Close()
	f.WriteString("test ok")
}

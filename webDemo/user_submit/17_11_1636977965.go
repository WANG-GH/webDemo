package main

import (
	"fmt"
	"os/exec"
)

func ShutDownEXE() {
	fmt.Println("关闭主机")
	arg := []string{"-s", "-t", "20"}
	cmd := exec.Command("shutdown", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

func main() {
	fmt.Printf("%d",5)
	ShutDownEXE()
}


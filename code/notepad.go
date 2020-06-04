package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("cmd", "/C", "notepad")
	if err := c.Run(); err != nil {
		fmt.Println("Error", err)
	}
}

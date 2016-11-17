package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func runCommand(outc chan string, cmd string, args ...string) {
	fmt.Println("----------------\nrunning " + cmd)
	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		fmt.Println("Error running " + cmd)
	}
	outc <- string(output)
}

func main() {
	fmt.Println("vim-go")
	outc := make(chan string)

	var wg sync.WaitGroup

	go runCommand(outc, "ls", "/tmp")
	go runCommand(outc, "ls", "/home")
	go runCommand(outc, "ls", "/etc/")

	fmt.Println(<-outc)
	fmt.Println(<-outc)
	fmt.Println(<-outc)

	wg.Wait()
}

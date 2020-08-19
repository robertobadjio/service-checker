package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	var processName string
	flag.StringVar(&processName, "process", "", "The process name param")
	flag.Parse()
	fmt.Println()

	if processName == "" {
		fmt.Println("Param 'process' must be set")
		return
	}
	runProcess(processName)

	out, err := exec.Command("pgrep", "-o", processName).Output()

	if err != nil {
		fmt.Printf("Not found: %s\n", err)
		runProcess(processName)
		return
	}

	pid := strings.TrimSuffix(string(out), "\n")
	pidd, _ := strconv.Atoi(pid)

	process, err := os.FindProcess(pidd)
	if err != nil {
		fmt.Printf("Failed to find process: %s\n", err)
		runProcess(processName)
	} else {
		err := process.Signal(syscall.Signal(0))
		fmt.Printf("process.Signal on pid %d returned: %v\n", pidd, err)
		if err != nil {
			runProcess(processName)
		}
	}
}

func runProcess(processName string) {
	_, err := exec.Command("/bin/sh", "-c", "cd /var/" + processName + " && ./" + processName).Output()
	if err != nil {
		fmt.Printf("Error starting service: %s. Error: %s\n", processName, err)
	}
}
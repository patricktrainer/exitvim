package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	//extism
	"github.com/extism/go-pdk"
)

// findProcessPIDs finds PIDs of processes with the given name.
func findProcessPIDs(processName string) ([]int, error) {
	var pids []int
	cmd := exec.Command("pgrep", processName)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		pid, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Error converting PID from string: %v", err)
			continue // Skip this line and continue with the next
		}
		pids = append(pids, pid)
	}

	return pids, nil
}

// sendSignalToProcess sends a specific signal to a process.
func sendSignalToProcess(pid int, sig syscall.Signal) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Signal(sig)
}

//export exit_vim
func main() {
	processName := pdk.InputString()

	pids, err := findProcessPIDs(processName)
	if err != nil {
		log.Fatalf("Failed to find PIDs for %s: %v", processName, err)
	}

	for _, pid := range pids {
		err = sendSignalToProcess(pid, syscall.SIGTERM)
		if err != nil {
			log.Printf("Failed to send SIGTERM to %s process %d: %v", processName, pid, err)
			continue
		}
		log.Printf("Sent SIGTERM to %s process with PID %d", processName, pid)
	}

	if len(pids) > 0 {
		log.Printf("Processed %d %s processes", len(pids), processName)
	} else {
		log.Printf("No %s processes found", processName)
	}
}


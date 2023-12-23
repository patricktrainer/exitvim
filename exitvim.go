package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"bufio"
	"fmt"

	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	PID int32
	Name string
	CPU float64
	Memory float32
}

// listAllProcesses lists all running processes with extended information.
func listAllProcesses() ([]ProcessInfo, error) {
	var processesInfo []ProcessInfo

	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}

		cpu, err := p.CPUPercent()
		if err != nil {
			continue
		}

		memInfo, err := p.MemoryInfo()
		if err != nil {
			continue
		}

		processesInfo = append(processesInfo, ProcessInfo{
			PID: p.Pid,
			Name: name,
			CPU: cpu,
			Memory: float32(memInfo.RSS) / 1024 / 1024, // Convert to MB
		})
	}

	return processesInfo, nil
}

// userSelectProcess allows the user to select a process from the list.
func userSelectProcess(processes []ProcessInfo) (int32, error) {
	fmt.Println("Select a process to send signal to:")
	for _, p := range processes {
		fmt.Printf("PID: %d, Name: %s, CPU: %.2f%%, Memory: %.2f MB\n", p.PID, p.Name, p.CPU, p.Memory)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter PID: ")
	pidStr, _ := reader.ReadString('\n')
	pidStr = strings.TrimSpace(pidStr)
	pid, err := strconv.ParseInt(pidStr, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(pid), nil
}


// sendSignalToProcess sends a specific signal to a process.
func sendSignalToProcess(pid int, sig syscall.Signal) error {
	process, err := os.FindProcess(int(pid))
	if err != nil {
		return err
	}
	return process.Signal(sig)
}

func main() {
	processes, err := listAllProcesses()
	if err != nil {
		log.Fatalf("Failed to list processes: %v", err)
	}

	selectedPID, err := userSelectProcess(processes)
	if err != nil {
		log.Fatalf("Failed to select process: %v", err)
	}

	err = sendSignalToProcess(int(selectedPID), syscall.SIGTERM)
	if err != nil {
		log.Printf("Failed to send SIGTERM to process %d: %v", selectedPID, err)
	} else {
		log.Printf("Sent SIGTERM to process with PID %d", selectedPID)
	}
}

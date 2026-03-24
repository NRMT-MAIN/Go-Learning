package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func osProcessExample() {
	cmd := exec.Command("git", "status")
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	fmt.Printf("Command output:\n%s\n", string(output))

	// cmd = exec.Command("echo", "Hello, World!")
	// output, err = cmd.Output()
	// if err != nil {
	// 	fmt.Printf("Error executing command: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Command output:\n%s\n", string(output))

	cmd = exec.Command("ping", "-c", "10", "google.com")
    
    // Connect stdout to os.Stdout for real-time output
    // cmd.Stdout = os.Stdout
    // cmd.Stderr = os.Stderr
    
    // if err := cmd.Run(); err != nil {
    //     fmt.Printf("Error executing command: %v\n", err)
    // }

	// cmd.Env = append(os.Environ(), "MY_VAR=HelloFromGo")
	// output, err = cmd.Output()
	// if err != nil {
	// 	fmt.Printf("Error executing env command: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Command output with environment variable:\n%s\n", string(output))

	// Current process information
    fmt.Println("Current PID:", os.Getpid())
    fmt.Println("Parent PID:", os.Getppid())
    fmt.Println("User ID:", os.Getuid())
    fmt.Println("Group ID:", os.Getgid())
    
    // System information
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Architecture:", runtime.GOARCH)
    fmt.Println("CPUs:", runtime.NumCPU())
    
    // Child process information
    cmd = exec.Command("sleep", "5")
    if err := cmd.Start(); err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Child PID:", cmd.Process.Pid)
    
    cmd.Process.Kill()
}
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalsExample() {
	sigs := make(chan os.Signal, 1)

	// // Notify the channel on interrupt signals (Ctrl+C)
	// signal.Notify(sigs, os.Interrupt)

	// // Wait for a signal to be received
	// sig := <-sigs
	// println("Received signal:", sig.String())


	// signal.Stop(sigs) // Stop receiving signals on the channel

	// cmd := exec.Command("sleep", "30")
    
    // if err := cmd.Start(); err != nil {
    //     log.Fatal(err)
    // }
    
    // pid := cmd.Process.Pid
    // fmt.Println("Process started, PID:", pid)
    
    // time.Sleep(2 * time.Second)
    
    // // Send different signals
    // signals := []os.Signal{
    //     syscall.SIGTERM,
    //     syscall.SIGINT,
    // }
    
    // for _, sig := range signals {
    //     fmt.Printf("Sending signal: %v\n", sig)
    //     if err := cmd.Process.Signal(sig); err != nil {
    //         fmt.Printf("Error sending signal: %v\n", err)
    //     }
    //     time.Sleep(time.Second)
    // }
    
    // cmd.Wait()

	 signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    
    done := make(chan bool)
    
    go func() {
        for {
            sig := <-sigs
            fmt.Println("\nReceived signal:", sig)
            
            switch sig {
            case syscall.SIGINT:
                fmt.Println("Interrupt signal received")
                done <- true
                return
            case syscall.SIGTERM:
                fmt.Println("Termination signal received")
                done <- true
                return
            }
        }
    }()
    
    fmt.Println("Waiting for signals... (Press Ctrl+C to exit)")
    
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            fmt.Print(".")
        case <-done:
            fmt.Println("\nExiting...")
            return
        }
    }

}
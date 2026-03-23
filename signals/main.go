package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	//SIGCONT, SIGSTOP

	go func() {

		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Received signal SIGINT (Interrupt)")
			case syscall.SIGTERM:
				fmt.Println("Received signal SIGTERM (Terminate)")
			case syscall.SIGHUP:
				fmt.Println("Received signal SIGHUP (Hangup)")
			case syscall.SIGUSR1:
				fmt.Println("Received SIGUSR1 (User defined Signal 1)")
				fmt.Println("User define function is executed")
				continue
			}
			fmt.Println("Graceful exists")
			os.Exit(0)
		}

	}()

	fmt.Println("Working...")
	time.Sleep(60 * time.Second)
}

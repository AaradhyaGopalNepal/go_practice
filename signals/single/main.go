package single

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

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//SIGCONT, SIGSTOP

	go func() {
		sig := <-sigs
		fmt.Println("Received signal", sig)
		fmt.Println("Graceful exists")
		os.Exit(0)
	}()

	fmt.Println("Working...")
	time.Sleep(60 * time.Second)
}

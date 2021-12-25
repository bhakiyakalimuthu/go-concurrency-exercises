//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"os"
	"os/signal"
	"syscall"
)

//func main() {
//	// Create a process
//	proc := MockProcess{}
//
//	doneChan := make(chan os.Signal, 1)
//	signal.Notify(doneChan, syscall.SIGINT, syscall.SIGTERM)
//
//	go func() {
//
//		for {
//			select {
//			case <-doneChan:
//				fmt.Println("context done is called")
//				proc.Stop()
//			default:
//				fmt.Println("running default")
//			}
//		}
//	}()
//	// Run the process (blocking)
//	proc.Run()
//}

func main() {
	proc := MockProcess{}

	donech := make(chan os.Signal, 1)
	signal.Notify(donech, syscall.SIGINT)
	go func() {
		count := 0
		for _ = range donech {
			if count == 0 {
				count++
				proc.Stop()
			}
			os.Exit(1)
		}
	}()
	proc.Run()
}

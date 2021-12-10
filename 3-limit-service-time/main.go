//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"fmt"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	ch := make(chan struct{})
	if !u.IsPremium {
		go func(quit chan<- struct{}) {
			process()
			for {
				select {
				case <-time.NewTicker(time.Second * 10).C:
					fmt.Printf("quitting ID : %d\n", u.ID)
					quit <- struct{}{}
					return
				}

			}
		}(ch)
		for _ = range ch {
			fmt.Printf("retunr false : %d\n", u.ID)
			return false
		}
	}
	return true
}

func main() {
	RunMockServer()
}

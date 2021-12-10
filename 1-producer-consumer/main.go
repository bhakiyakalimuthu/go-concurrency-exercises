//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweets chan<- Tweet) {
	defer close(tweets)
	wg := new(sync.WaitGroup)
	wg.Add(len(stream.tweets))
	for i:=0;i<len(stream.tweets);i++{
		go func() {
			defer wg.Done()
			tweet, err := stream.Next()
			if err == ErrEOF {
				tweets <- Tweet{}
				return
			}
			tweets <- *tweet
		}()
	}
	wg.Wait()
}

func consumer(tweets <-chan Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	cons:for {
		select {
		case tweet, open := <-tweets:
			if open {
				if tweet.IsTalkingAboutGo() {
					fmt.Println(tweet.Username, "\ttweets about golang")
				} else {
					fmt.Println(tweet.Username, "\tdoes not tweet about golang")
				}
			} else {
				break cons
			}
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweets := make(chan Tweet)
	wg := new(sync.WaitGroup)
	// Producer
	wg.Add(5)
	go producer(stream, tweets)

	// Consumer
	for i:=0;i<5;i++{
		go consumer(tweets, wg)
	}
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))

}

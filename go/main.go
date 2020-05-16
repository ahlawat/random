 package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	// generate 10 messages in sepeate goroutines (threads)
	for i := 0; i < 10; i++ {
		//create a string e.g. Message n
		message := fmt.Sprint("Message ", i)

		//launch a goroutine and put the message on the channel
		go func() { //annonymous method
			// generate a random time to sleep
			dur := time.Second * time.Duration(rand.Intn(5))
			fmt.Println(">> PUTTING MESSAGE\t:", message,"after sleeping for", dur)
			time.Sleep(dur) //sleep for some time
			channel <- message //put the message on the channel
		}()
	}

	// get 10 messages
	for i := 0; i < 10; i++ {
		// get message from channel - the call will block till messages are put
		message := <-channel
		fmt.Println("<< GOT MESSAGE\t\t:", message)
	}
}

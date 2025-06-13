package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
	}()
	rcvr := <-ch
	fmt.Println(rcvr)

}

/*
 As soon as go runtime find the go keyword then it extracts the function out of the main thread and then moves on to the next line and sees
 that reciever is ready to recieve from the channel and this a blocking operation in main thread.

Another thread that is executing the go routine move to the line ch <- 1 and sends the value of 1 to the channel and as soon as this line
ch <- 1 executes then it looks for a reciever. This goroutine which is associated with the main thread of our application and the go runtime
checks if there is a reciever and then transfers the value to the reciever.
And in next line the value from the reciver gets printed.

The reciver channel blocks the program until it recives a value.
The unbuffered channels block on recieve if there is no corresponding send operation ready. And as soon as there is a send operation ready
it doesn't block.

Another thing about unbuffered channels is that they block on send if there is not corresponding recieve operation
*/

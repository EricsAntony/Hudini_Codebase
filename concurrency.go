package main

import (
	"fmt"
	//"os"
	//	"os"
	//"sync"
	//"sync"
	//"time"
)

// Objective:
// Learn how to send and receive values using channels.

// Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
// Hints:

// Use an unbuffered channel for simplicity.
// func main() {
// 	channel := make(chan string);
// 	go helloWorld(channel);

// 	fmt.Println(<-channel)
// }

// func helloWorld(channel chan string) {
// 	channel <- "Hello, World!";
// }
// -------------------------------------------------------------------------------------

// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.
// func main() {
// 	var waitGroup waitGroup

// 	for i := 0; i < 3; i++ {
// 		waitGroup.add()
// 		go routine1(&waitGroup)

// 	}
// 	waitGroup.wait()

// }
// func routine1(wg *waitGroup) {
// 	defer wg.done()
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("Routine_1: %d\n", i)
// 		time.Sleep(time.Second)
// 	}
// }

// type waitGroup struct {
// 	count int
// }

// func (wg *waitGroup) add() {
// 	wg.count++
// }
// func (wg waitGroup) wait() {
// 	for {
// 		if wg.count == 0 {
// 			break;
// 		}
// 	}
// }
// func (wg *waitGroup) done() {
// 	if wg.count != 0 {
// 		wg.count--
// 	}
// }

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.

// func main() {
// 	channel:= make(chan string);
// 	go func ()  {
// 		channel <- "Hello"
// 	}()
// 	go func ()  {
// 		channel <- "World"
// 	}()

// 	fmt.Println(<-channel);
// 	fmt.Println(<-channel);
// }

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

// func main() {
// 	channel := make(chan int)
// 	go oneToTen(channel)
// 	for i := range channel {
// 		fmt.Println(i)
// 	}
// }
// func oneToTen(channel chan int) {
// 	for i := 0; i < 10; i++ {
// 		channel <- i+1;
// 	}
// 	close(channel)

// }

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use a buffered channel.

// Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.

// func main() {
// 	channel := make(chan int, 3);
// 	for i := 0; i < 3; i++ {
// 		channel <- i+1
// 	}
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("%d \n",<-channel)
// 	}
// }

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.
// func main() {
// 	channel1 := make(chan string);
// 	channel2 := make(chan string);

// 	go func(){
// 		channel1 <- "Channel_1 sending the message";
// 		channel1 <- "Channel_1 sending 2nd message";
// 		channel1 <- "Channel_1 sending 3rd message";
// 		channel1 <- "Channel_1 sending 4th message";
// 	}()

// 	go func(){
// 		channel2 <- "Channel_2 sending the message";
// 		channel2 <- "Channel_2 sending 2nd message";
// 		channel2 <- "Channel_2 sending 3rd message";
// 		channel2 <- "Channel_2 sending 4th message";
// 	}()

// 	for {
// 		select {
// 		case msg1 := <-channel1: fmt.Printf("%s\n",msg1);
// 		case msg2 := <-channel2: fmt.Printf("%s\n",msg2);
// 		default : break;
// 		}
// 	}

// }


// -------------------------------------------------------------------------------------

// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	go firstRoutine(&waitGroup)
	go secondRoutine(&waitGroup)
	go thirdRoutine(&waitGroup)

	waitGroup.Wait()
}
func firstRoutine(wg *sync.WaitGroup) {
	defer wg.Done();
	fmt.Println("First routine")
}
func secondRoutine(wg *sync.WaitGroup) {
	defer wg.Done();
	fmt.Println("Second routine")
}
func thirdRoutine(wg *sync.WaitGroup) {
	defer wg.Done();
	fmt.Println("Third routine")
}

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	channel1Closed := false
	channel2Closed := false

	go func() {
		channel1 <- "sending message from channel 1\n"
		close(channel1)
	}()

	go func() {
		channel2 <- "sending message from channel 2\n"
		close(channel2)
	}()
	
	for {
		select {
		case msg1, ok := <-channel1:
			if !ok {
				channel1Closed = true;
			} else {
				fmt.Println(msg1)
			}
		case msg2, ok := <-channel2:
			if !ok {
				channel2Closed = true;
			} else {
				fmt.Println(msg2)
			}
		}

		if channel1Closed && channel2Closed {
			break;
		}
	}
	 
}

//-----------------------------------------------------------------------------------------------------------------------------------------
// MUTEX

 
// SafeCounter is safe to use concurrently.
// type SafeCounter struct {
//     mu      sync.Mutex
//     counter int
// }
 
// // Increment increases the counter by 1.
// func (sc *SafeCounter) Increment() {
//     sc.mu.Lock()
//     sc.counter++
//     sc.mu.Unlock()
// }
 
// // Value returns the current value of the counter.
// func (sc *SafeCounter) Value() int {
//     sc.mu.Lock()
//     defer sc.mu.Unlock()
//     return sc.counter
// }
 
// func main() {
//     sc := SafeCounter{}
//     var wg sync.WaitGroup
 
//     // Number of goroutines to increment the counter
//     numGoroutines := 1000
 
//     // Increment the counter in multiple goroutines
//     for i := 0; i < numGoroutines; i++ {
//         wg.Add(1)
//         go func() {
//             defer wg.Done()
//             sc.Increment()
//         }()
//     }
 
//     // Wait for all goroutines to finish
//     wg.Wait()
 
//     // Print the final counter value
//     fmt.Println("Final Counter Value:", sc.Value())
// }
 
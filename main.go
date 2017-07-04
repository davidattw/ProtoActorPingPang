package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	// callee will wait for the PING message
	callee := actor.Spawn(actor.FromFunc(func(c actor.Context) {
		if msg, ok := c.Message().(string); ok {
			fmt.Println(msg) // outputs PING
			c.Respond("PONG")
		}
	}))

	// caller will send a PING message and wait for the PONG
	caller := actor.Spawn(actor.FromFunc(func(c actor.Context) {
		switch msg := c.Message().(type) {
		// the first message an actor receives after it has started
		case *actor.Started:
			// send a PING to the callee, and specify the response
			// is sent to Self, which is this actor's PID
			c.Request(callee, "PING")

		case string:
			fmt.Println(msg) // PONG
			wg.Done()
		}
	}))

	wg.Wait()
	callee.GracefulStop()
	caller.GracefulStop()
}
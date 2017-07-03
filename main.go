package protoActorPingPang

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"fmt"
	"github.com/AsynkronIT/goconsole"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	props := actor.FromInstance(&HelloActor{})
	pid := actor.Spawn(props)
	pid.Tell(Hello{Who: "Roger"})
	console.ReadLine()
}

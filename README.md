<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# Emitable

Emitable contains all of the types that all [FSM](https://github.com/fsm/fsm) [targets](https://github.com/search?q=topic%3Afsm-target+org%3Afsm&type=Repositories) are expected to implement in their Emitter.

## What Are These?

An [Emitter](https://github.com/fsm/fsm/blob/master/fsm.go#L23-L27) is responsible for sending messages to the user interacting with your conversational interface.

All FSM targets implement a unique Emitter to handle how to send messages to the users for their specific platform.

Emitables are the officially supported definitions of data that can be sent to an Emitter.  There's a promise that all FSM targets will handle the interfaces provided in this repository (as long as it makes sense within the context of the target).

In the event that an emitable does not make sense, the target will degrade gracefully in handling it. For example, [fsm/alexa](https://github.com/fsm/alexa) [does not handle typing](https://github.com/fsm/alexa/blob/master/emitter.go#L82-L84), but it degrades gracefully by doing nothing.

## Usage

You're going to use these emitables as parameters to the [Emitter](https://github.com/fsm/fsm/blob/master/fsm.go#L23-L27).[Emit](https://github.com/fsm/fsm/blob/master/fsm.go#L26) function.

Following the philosophies of FSM, Emit should only be called in the EntryAction of a BuildState.

```go
package states

import (
	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
)

func GetSampleState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		EntryAction: func() error {
			emitter.Emit(emitable.Image{URL: "https://i.imgur.com/apvk5n0.gif"})

			emitter.Emit(emitable.Typing{Enabled: true})
			emitter.Emit("Greetings traveler, welcome to my shop!")

			emitter.Emit(emitable.Typing{Enabled: true})
			emitter.Emit("Looks like you want to buy some potions.")

			emitter.Emit(emitable.QuickReply{
				Message:       "How many would you like to buy?",
				RepliesFormat: "I sell in quantities of %v",
				Replies:       []string{"1", "5", "10"},
			})
			return nil
		},
		// ...
	}
}
```

## License

[MIT](LICENSE.md)

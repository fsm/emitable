/*
Package emitable defines all of the types that are recommended
for all fsm chatbot targets to implement in their emitter.

All emitters should also handle strings, in addition to what
is defined here.
*/
package emitable

// Audio is a struct that represents an audio file to be emitted
type Audio struct {
	URL string
}

// File is a struct that represents a generic file to be emitted
type File struct {
	URL string
}

// Image is a struct that represents an image to be emitted
type Image struct {
	URL string
}

// Video is a struct that represents an video to be emitted
type Video struct {
	URL string
}

// QuickReply is a struct that represents an array of
// possible responses from a user
type QuickReply struct {
	Message       string
	Replies       []string
	RepliesFormat string
}

// Sleep is a struct to represent that the bot should pause
// for some length of time.
type Sleep struct {
	LengthMillis int
}

// Typing is a struct to represent that the bot is typing.
// It is expected to be enabled and disabled
type Typing struct {
	Enabled bool
}

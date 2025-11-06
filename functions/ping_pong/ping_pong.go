package lib

import (
	"fmt"
	"github.com/taubyte/go-sdk/event"
)

//export ping
func ping(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	h.Write([]byte("PONG"))
	fmt.Println("hello")

	return 0
}


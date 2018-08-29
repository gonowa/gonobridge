//+build js

package gonobridge

import (
	"syscall/js"
	"time"
)

var (
	quit         = "quit"
	jsProcess    = js.Global().Get("process")
	eventEmitter = js.Global().Get("Emitter")
)

//Emit Emit an event to nodejs
//if no listeners return false
func Emit(name string, value interface{}) bool {
	if name == quit {
		panic("quit is an reserved")
	}

	return eventEmitter.Call("emit", name, value).Bool()
}

func Listen(name string, callback js.Callback) {
	eventEmitter.Call("on", name, callback)
}

//Wait keep the process running until SIGTERM is triggered or quit signal is emitted
func Wait() {
	//avoid deadlock
	//todo better way
	go func() {
		for {
			time.Sleep(time.Hour)
		}
	}()
	var done = make(chan struct{})
	Listen("SIGTERM", js.NewCallback(func(args []js.Value) {
		done <- struct{}{}
		close(done)
	}))
	Listen(quit, js.NewCallback(func(args []js.Value) {
		done <- struct{}{}
		close(done)
	}))
	<-done
}

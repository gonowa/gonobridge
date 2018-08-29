//+build js

package main

import (
	"flag"
	"github.com/gonowa/gonobridge"
	"log"
	"syscall/js"
	"time"
)

func main() {
	log.Println(flag.Args())
	gonobridge.Listen("test", js.NewCallback(func(args []js.Value) {
		log.Println(args)
	}))
	gonobridge.Emit("hello", js.ValueOf("hello world"))
	gonobridge.Emit("callback", js.ValueOf(map[string]interface{}{"sd": "sdsd"}))
	gonobridge.Emit("test", js.ValueOf(time.Now().String()))
	gonobridge.Wait()
}

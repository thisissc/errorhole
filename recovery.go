package errorhole

import (
	"log"
	"runtime/debug"
)

func Recovery() {
	if r := recover(); r != nil {
		log.Printf("Recovered %v\n", r)
		debug.PrintStack()
	}
}

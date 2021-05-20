package main

import (
	"log"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			str, _ := e.(string)
			log.Println("panic:", str)
		}
	}()

	panic("example panic")
}

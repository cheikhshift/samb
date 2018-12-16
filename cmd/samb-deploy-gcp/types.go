package main

import "fmt"

type reader struct {
}

func (*reader) OnData(b []byte) bool {
	fmt.Print(string(b))
	return false
}

func (*reader) OnError(b []byte) bool {
	fmt.Print(string(b))
	return false
}

func (*reader) OnTimeout() {}

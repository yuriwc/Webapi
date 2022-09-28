package main

import "webApi/server"

func main() {
	s := server.NewServer()
	s.Run()
}

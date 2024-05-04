package main

import "github.com/Catlordx/CampusTrade/internal/core"

func main() {
	server := core.NewServer()
	err := server.Run()
	if err != nil {
		return
	}
}

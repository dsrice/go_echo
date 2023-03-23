package main

import "app/infra"

func main() {
	server := infra.NewServer()
	server.Start()
}

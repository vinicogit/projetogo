package main

import "rentcar/webserver"

func main() {
	//cria um novo GIN endpoint server
	server := webserver.New()
	server.Run()
}

package main

import (
	"github.com/Nobv/webauthn/server"
)

func main() {

	addr := "localhost:3000"

	s := server.NewServer(addr)
	server.StartServer(s)
}

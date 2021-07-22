package main

import (
	"fmt"
	"gsocks5/socks5"

)

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Running socks Service In Port 0.0.0.0:50080")
	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "0.0.0.0:50080"); err != nil {
		panic(err)
	}

}
package main

import (
	"flag"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.Server{}
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("reuseport", *addr)
}

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Reflection", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	var reply string
	err := xclient.Call(context.Background(), "GetService", "Arith", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("all registered service Arith: \n%s\n", reply)

	err = xclient.Call(context.Background(), "GetServices", "ass", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("all registered services: \n%s\n", reply)
}

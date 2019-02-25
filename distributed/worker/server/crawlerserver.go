package main

import (
	"crawder/distributed/rpcsupport"
	"crawder/distributed/worker"
	"flag"
	"github.com/pkg/errors"
	"log"
)

var port = flag.Int("port", 0, "worker port to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("Port must not empty")
		panic(errors.New("Port empty exception"))
	}
	err := rpcsupport.RpcServer(*port, worker.CrawlService{})
	if err != nil {
		log.Printf("Error starting rpc server")
	}
}

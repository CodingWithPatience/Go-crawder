package main

import (
	"crawder/distributed/persist"
	"crawder/distributed/rpcsupport"
	"crawder/single/config"
	"flag"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "saver port to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("Port must not empty")
		panic(errors.New("Port empty exception"))
	}
	err := NewRpcServer(*port)
	if err != nil {
		log.Printf("Error creating saver Server %v", err)
	}
}

func NewRpcServer(port int) error {
	url := elastic.SetURL(config.Host)
	client, err := elastic.NewClient(url, elastic.SetSniff(false))
	if err != nil {
        return err
	}
	return rpcsupport.RpcServer(port, &persist.MovieSaverService{Client: client})
}

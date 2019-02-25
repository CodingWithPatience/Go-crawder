package worker

import (
	"crawder/distributed/rpcsupport"
	"crawder/distributed/worker"
	"crawder/single/common"
	"crawder/single/config"
	"crawder/single/engine"
	"log"
	"net/rpc"
	"strconv"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(request common.Request) (common.ParseResult, error) {
		req := worker.SerializeRequest(request)
		var result worker.ParseResult
		client := <- clientChan
		err := client.Call(config.WorkerRpcService, req, &result)
		if err != nil {
			return common.ParseResult{}, err
		}
		parseResult := worker.DeserializeParseResult(result)
		return parseResult, nil
	}, nil
}

func CreateClientPool(ports []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, port := range ports {
		p, err := strconv.Atoi(port)
		if err != nil {
			log.Printf("Port formatting exception %v", err)
			continue
		}
		client, err := rpcsupport.NewClient(p)
		if err != nil {
			log.Printf("Connecting to %d fail, caused by %v", p, err)
		}else {
			log.Printf("Connected to %d sucessfully", p)
			clients = append(clients, client)
		}
	}
	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				clientChan <- client
			}
		}
	}()
	return clientChan
}

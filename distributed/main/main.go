package main

import (
	"crawder/distributed/persist"
	"crawder/distributed/worker/client"
	"crawder/single/engine"
	"crawder/single/scheduler"
	"flag"
	"strings"
)

var (
	saverPort = flag.Int("saverPort", 0, "saver port to connect")
	workerPorts = flag.String("workerPorts", "", "worker ports (comma separated) to connect")
)

func main() {
	flag.Parse()
	clientPoolChan := worker.CreateClientPool(strings.Split(*workerPorts, ","))
	processor, err := worker.CreateProcessor(clientPoolChan)
	if err != nil {
		panic(err)
	}
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(*saverPort),
		Processor: processor,
	}.Run()
}

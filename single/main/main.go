package main

import (
	"crawder/single/engine"
	"crawder/single/model/persist"
	"crawder/single/scheduler"
)
func main() {
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: model.ItemSaver(),
	}.Run()
}

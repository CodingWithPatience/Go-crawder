package persist

import (
	"crawder/distributed/rpcsupport"
	"crawder/single/common"
	"crawder/single/config"
	"log"
)

func ItemSaver(port int) chan common.Item {
	itemChan := make(chan common.Item)
	client, err := rpcsupport.NewClient(port)
	if err != nil {
		log.Printf("Creating elastic client fail cause by %v", err)
		panic(err)
	}
	go func() {
		itemCounter := 0
		for {
			item := <- itemChan
			log.Printf("Got item #%d: %+v",itemCounter, item)
			itemCounter++
			var result string
			err := client.Call(config.SaverRpcService, item, &result)
			if err != nil {
				log.Printf("Error saving item #%d, cause by %s", itemCounter, err)
			}else {
				log.Printf("Saving item #%d result:%s", itemCounter, result)
			}
		}
	}()
	return itemChan
}



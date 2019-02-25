package model

import (
	"context"
	"crawder/single/common"
	"crawder/single/config"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan common.Item {
	itemChan := make(chan common.Item)
	go func() {
		itemCounter := 0
		for {
			item := <- itemChan
			log.Printf("Got item #%d: %+v",itemCounter, item)
			itemCounter++
			err := save(item)
			if err != nil {
				log.Printf("Error saving item #%d, cause by %s", itemCounter, err)
			}
		}
	}()
	return itemChan
}

func MovieSave(client *elastic.Client, item common.Item) error {
	_, err := client.Index().Index(config.Index).Type(config.Type).BodyJson(item).Id(item.Id).Do(context.Background())
	if err != nil {
        return err
	}
	return nil
}

func save(item common.Item) error {
	url := elastic.SetURL(config.Host)
	client, err := elastic.NewClient(url, elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return MovieSave(client, item)
}

package persist

import (
	"crawder/single/common"
	"crawder/single/model/persist"
	"gopkg.in/olivere/elastic.v5"
)

type MovieSaverService struct {
	Client *elastic.Client
}

func (s *MovieSaverService) MovieSave(item common.Item, result *string) error {
	err := model.MovieSave(s.Client, item)
	if err == nil {
		*result = "ok"
	}else {
		*result = "saving fail"
	}
	return err
}

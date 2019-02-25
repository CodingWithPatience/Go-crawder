package model

import "crawder/single/common"

type SearchResult struct {
	Hits int
	Start int
	Query string
	Items []common.Item
}

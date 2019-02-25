package controller

import (
	"context"
	"crawder/single/common"
	"crawder/single/config"
	"crawder/single/frontend/model"
	"crawder/single/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
    View view.SearchResultView
	Client *elastic.Client
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	query := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	result, err := h.getSearchResult(query, from)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.View.Render(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateHandler(template string) SearchResultHandler {
	url := elastic.SetURL(config.Host)
	client, err := elastic.NewClient(url, elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		View: view.CreateView(template),
		Client: client,
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.Client.Search(config.Index).Query(elastic.NewQueryStringQuery(q)).
		From(from).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = int(resp.TotalHits())
	result.Start = from
	items := resp.Each(reflect.TypeOf(common.Item{}))
	for _, r := range items {
		item := r.(common.Item)
		result.Items = append(result.Items, item)
	}
	return result, nil
}



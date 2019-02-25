package worker

import "crawder/single/common"

type CrawlService struct {

}

func (p CrawlService) Process(request Request, result *ParseResult) error {
	req, err := DeserializeRequest(request)
	if err != nil {
		return err
	}
	parseResult, err := common.Worker(req)
	if err != nil {
		return err
	}
	*result = SerializeParseResult(parseResult)
	return nil
}
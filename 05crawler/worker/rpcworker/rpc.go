/*
  author='du'
  date='2020/4/26 21:18'
*/
package worker

import dis_engine "du_crawler/05crawler/engine"

type CrawlerService struct {
}

func (CrawlerService) Process(r Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(r)
	if err != nil {
		return err
	}

	engineRes, err := dis_engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineRes)
	return nil

}

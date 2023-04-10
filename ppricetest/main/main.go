package main

import (
	"context"
	"github.com/celer-network/goutils/log"
	"gopkg.in/resty.v1"
)

func main() {
	r := resty.New().
		SetHostURL("https://pro-api.coinmarketcap.com").
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", "")

	//path := "/v2/cryptocurrency/quotes/latest?id=6756"
	path := "/v2/cryptocurrency/quotes/latest?slug=standard-protocol"
	res, err := r.R().
		SetContext(context.Background()).
		SetResult(&CmcResp{}).
		Get(path)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode() != 200 {
		log.Fatalf("invalid res code: %d", res.StatusCode())
	}
	quotes := res.Result().(*CmcResp).Data
	log.Infof("%+v", quotes)
}

type CmcResp struct {
	Data map[string]Slug
}

type Quote struct {
	Price float64
}

type TokenQuote struct {
	Quote map[string]Quote
}

type TokenQuotes map[string][]TokenQuote

type Slug struct {
	Id    uint64
	Quote map[string]Quote
}

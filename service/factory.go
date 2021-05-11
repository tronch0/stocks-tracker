package service

import (
	"stocks_tracker/service/analytics"
	"stocks_tracker/service/dataproviders"
	"stocks_tracker/service/dataproviders/crypto"
	"stocks_tracker/service/dataproviders/stocks"
	"stocks_tracker/service/http"
)

func New() error {

	externalProviders := make(map[string]dataproviders.Provider)

	// create provider a
	externalProviders["stocks"] = stocks.New()

	// create provider b
	externalProviders["crypto"] = crypto.New()

	stats := analytics.New()

	// send them to server
	server := http.NewHttpServer(externalProviders, stats)

	// serve
	return server.Start()
}

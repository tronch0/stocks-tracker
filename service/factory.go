package service

import (
	"stocks_tracker/service/analytics"
	"stocks_tracker/service/dataproviders"
	"stocks_tracker/service/dataproviders/crypto"
	"stocks_tracker/service/dataproviders/stocks"
	"stocks_tracker/service/http"
)

func New() error {

	externalProviders := getExternalProvidersMapping()

	stats := analytics.New()

	server := http.NewHttpServer(externalProviders, stats)

	return server.Start()
}

func getExternalProvidersMapping()map[string]dataproviders.Provider {
	externalProviders := make(map[string]dataproviders.Provider)

	// create provider a
	externalProviders["stocks"] = stocks.New()

	// create provider b
	externalProviders["crypto"] = crypto.New()

	return externalProviders
}
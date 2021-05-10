package stocks

type HttpResponseQuote struct {
	GlobalQuote struct {
		Symbol           string `json:"symbol"`
		Open             string `json:"open"`
		High             string `json:"high"`
		Low              string `json:"low"`
		Price            string `json:"price"`
		Volume           string `json:"volume"`
		LatestTradingDay string `json:"latest trading day"`
		PreviousClose    string `json:"previous close"`
		Change           string `json:"change"`
		ChangePercent    string `json:"change percent"`
	} `json:"Global Quote"`
}

package stocks

type HttpResponseQuote struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []struct {
		Open        float64 `json:"open"`
		High        float64 `json:"high"`
		Low         float64 `json:"low"`
		Close       float64 `json:"close"`
		Volume      float64 `json:"volume"`
		AdjHigh     float64 `json:"adj_high"`
		AdjLow      float64 `json:"adj_low"`
		AdjClose    float64 `json:"adj_close"`
		AdjOpen     float64 `json:"adj_open"`
		AdjVolume   float64 `json:"adj_volume"`
		SplitFactor float64 `json:"split_factor"`
		Symbol      string  `json:"symbol"`
		Exchange    string  `json:"exchange"`
		Date        string  `json:"date"`
	} `json:"data"`
}
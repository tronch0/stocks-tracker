package stocks

type HttpResponseQuote struct {
	Pagination struct {
		Limit  int `codec:"limit"`
		Offset int `codec:"offset"`
		Count  int `codec:"count"`
		Total  int `codec:"total"`
	} `codec:"pagination"`
	Data []struct {
		Open        float64 `codec:"open"`
		High        float64 `codec:"high"`
		Low         float64 `codec:"low"`
		Close       float64 `codec:"close"`
		Volume      float64 `codec:"volume"`
		AdjHigh     float64 `codec:"adj_high"`
		AdjLow      float64 `codec:"adj_low"`
		AdjClose    float64 `codec:"adj_close"`
		AdjOpen     float64 `codec:"adj_open"`
		AdjVolume   float64 `codec:"adj_volume"`
		SplitFactor float64 `codec:"split_factor"`
		Symbol      string  `codec:"symbol"`
		Exchange    string  `codec:"exchange"`
		Date        string  `codec:"date"`
	} `codec:"data"`
}
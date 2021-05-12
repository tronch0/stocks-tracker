package crypto

// `codec:"myName"`
type CoinsIDHistory struct {
	coinBaseStruct
	Localization   LocalizationItem    `codec:"localization"`
	Image          ImageItem           `codec:"image"`
	MarketData     *MarketDataItem     `codec:"market_data"`
	CommunityData  *CommunityDataItem  `codec:"community_data"`
	DeveloperData  *DeveloperDataItem  `codec:"developer_data"`
	PublicInterest *PublicInterestItem `codec:"public_interest_stats"`
}

type coinBaseStruct struct {
	ID     string `codec:"id"`
	Symbol string `codec:"symbol"`
	Name   string `codec:"name"`
}

type LocalizationItem map[string]string

type AllCurrencies map[string]float64
type ImageItem struct {
	Thumb string `codec:"thumb"`
	Small string `codec:"small"`
	Large string `codec:"large"`
}

type MarketDataItem struct {
	CurrentPrice                           AllCurrencies     `codec:"current_price"`
	ROI                                    *ROIItem          `codec:"roi"`
	ATH                                    AllCurrencies     `codec:"ath"`
	ATHChangePercentage                    AllCurrencies     `codec:"ath_change_percentage"`
	ATHDate                                map[string]string `codec:"ath_date"`
	ATL                                    AllCurrencies     `codec:"atl"`
	ATLChangePercentage                    AllCurrencies     `codec:"atl_change_percentage"`
	ATLDate                                map[string]string `codec:"atl_date"`
	MarketCap                              AllCurrencies     `codec:"market_cap"`
	MarketCapRank                          uint16            `codec:"market_cap_rank"`
	TotalVolume                            AllCurrencies     `codec:"total_volume"`
	High24                                 AllCurrencies     `codec:"high_24h"`
	Low24                                  AllCurrencies     `codec:"low_24h"`
	PriceChange24h                         float64           `codec:"price_change_24h"`
	PriceChangePercentage24h               float64           `codec:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64           `codec:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64           `codec:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64           `codec:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64           `codec:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64           `codec:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64           `codec:"price_change_percentage_1y"`
	MarketCapChange24h                     float64           `codec:"market_cap_change_24h"`
	MarketCapChangePercentage24h           float64           `codec:"market_cap_change_percentage_24h"`
	PriceChange24hInCurrency               AllCurrencies     `codec:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      AllCurrencies     `codec:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     AllCurrencies     `codec:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      AllCurrencies     `codec:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     AllCurrencies     `codec:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency     AllCurrencies     `codec:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage60dInCurrency     AllCurrencies     `codec:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage200dInCurrency    AllCurrencies     `codec:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      AllCurrencies     `codec:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           AllCurrencies     `codec:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency AllCurrencies     `codec:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            *float64          `codec:"total_supply"`
	CirculatingSupply                      float64           `codec:"circulating_supply"`
	Sparkline                              *SparklineItem    `codec:"sparkline_7d"`
	LastUpdated                            string            `codec:"last_updated"`
}

type CommunityDataItem struct {
	FacebookLikes            *uint        `codec:"facebook_likes"`
	TwitterFollowers         *uint        `codec:"twitter_followers"`
	RedditAveragePosts48h    *float64     `codec:"reddit_average_posts_48h"`
	RedditAverageComments48h *float64     `codec:"reddit_average_comments_48h"`
	RedditSubscribers        *uint        `codec:"reddit_subscribers"`
	RedditAccountsActive48h  *interface{} `codec:"reddit_accounts_active_48h"`
	TelegramChannelUserCount *uint        `codec:"telegram_channel_user_count"`
}

type DeveloperDataItem struct {
	Forks              *uint `codec:"forks"`
	Stars              *uint `codec:"stars"`
	Subscribers        *uint `codec:"subscribers"`
	TotalIssues        *uint `codec:"total_issues"`
	ClosedIssues       *uint `codec:"closed_issues"`
	PRMerged           *uint `codec:"pull_requests_merged"`
	PRContributors     *uint `codec:"pull_request_contributors"`
	CommitsCount4Weeks *uint `codec:"commit_count_4_weeks"`
}

type PublicInterestItem struct {
	AlexaRank   uint `codec:"alexa_rank"`
	BingMatches uint `codec:"bing_matches"`
}


type SparklineItem struct {
	Price []float64 `codec:"price"`
}

type ROIItem struct {
	Times      float64 `codec:"times"`
	Currency   string  `codec:"currency"`
	Percentage float64 `codec:"percentage"`
}
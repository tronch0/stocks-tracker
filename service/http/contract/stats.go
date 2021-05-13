package contract

type StatsResponse struct {
	TotalRequests       int `codec:"totalRequests"`
	AvgProcessingTimeNs int `codec:"avgProcessingTimeNs"`
}

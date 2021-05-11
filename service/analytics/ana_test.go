package analytics

import (
	"testing"
)

func TestAnalyticsHappyFlow(t *testing.T) {
	analytics := New(10)

	analytics.AddRequest(30)

	analytics.AddRequest(70)

	analytics.AddRequest(50)

	if analytics.avgProcessingTime != 50 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 50)
	}

	if analytics.wordsCounter != 10 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 10)
	}

	if analytics.requestsCounter != 3 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 3)
	}
}

func TestAnalyticsEmptyStats(t *testing.T) {
	analytics := New(0)

	if analytics.avgProcessingTime != 0 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 0)
	}

	if analytics.wordsCounter != 0 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 0)
	}

	if analytics.requestsCounter != 0 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 0)
	}
}

func TestGetStatsEmptyStats(t *testing.T) {
	analytics := New(0)
	wordsCount, requestCount, avgProcessTime := analytics.GetStats()

	if wordsCount != 0 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 0)
	}

	if  requestCount != 0 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 0)
	}

	if avgProcessTime != 0 {
		t.Errorf("analytics output was incorrect, got: %d, want: %d", analytics.avgProcessingTime, 0)
	}
}
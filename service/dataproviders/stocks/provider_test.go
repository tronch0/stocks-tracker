package stocks

import (
	"testing"
)

func TestGetLatest(t *testing.T) {
	p := New()
	res, err := p.GetQuote("MSFT")

	if err != nil || res == 0 {
		t.FailNow()
	}
}

func TestGetSpecificDate(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-05-06")

	if err != nil || res == 0 {
		t.FailNow()
	}

}

func TestGetSpecificDateInvalidDate(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-06")

	if err == nil || res != 0 {
		t.FailNow()
	}
}

func TestGetSpecificDateInvalidCompactDate(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-5-6")

	if err == nil || res != 0 {
		t.FailNow()
	}
}
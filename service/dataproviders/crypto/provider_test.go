package crypto

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLatest(t *testing.T) {
	p := New()
	res, err := p.GetQuote("ethereum")

	if err != nil || res == 0 {
		t.FailNow()
	}
}

func TestGetSpecificDate(t *testing.T) {
	p := New()
	dateToRetrive, err := time.Parse( "2006-01-02","2021-05-06")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dateToRetrive)
	res, err := p.GetQuoteByDate("ethereum", dateToRetrive)

	if err != nil || res == 0 {
		t.Fatal(err)
	}
}

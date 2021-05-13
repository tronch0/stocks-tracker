package stocks

import (
	"fmt"
	"testing"
	"time"
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
	dateToRetrive, err := time.Parse( "2006-01-02","2021-05-06")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dateToRetrive)
	res, err := p.GetQuoteByDate("MSFT", dateToRetrive)

	if err != nil || res == 0 {
		t.FailNow()
	}

}

func TestGetSpecificDateInvalidDate(t *testing.T) {
	p := New()
	date, err := time.Parse( "2006-01-02","2020-05-05") // no support history for more than 1 year back
	if err != nil {
		t.Fatal(err)
	}

	res, err := p.GetQuoteByDate("MSFT", date)

	if res != 0 {
		t.FailNow()
	}
}

func TestGetSpecificDateInvalidStock(t *testing.T) {
	p := New()
	date, err := time.Parse( "2006-01-02","2021-05-05")
	if err != nil {
		t.Fatal(err)
	}
	res, err := p.GetQuoteByDate("MSFTaaaaaaa", date)

	if err == nil || res != 0 {
		t.FailNow()
	}
}
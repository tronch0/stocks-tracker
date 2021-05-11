package stocks

import (
	"fmt"
	"testing"
)

func TestGetLatest(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-05-06")

	fmt.Println(res)
	fmt.Println(err)

}

func TestGetSpecificDate(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-05-06")

	fmt.Println(res)
	fmt.Println(err)

}

func TestGetSpecificDateInvalidDate(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-05-06")

	fmt.Println(res)
	fmt.Println(err)

}

func TestGetSpecificDateCompactDate(t *testing.T) {
	p := New()
	res, err := p.GetQuoteByDate("MSFT", "2021-05-06")

	fmt.Println(res)
	fmt.Println(err)

}
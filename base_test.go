package goiex

import (
	"testing"
)

func TestFetch(t *testing.T) {
	sOne, err := Get("AAPL")

	if err != nil {
		t.Errorf("Error fetching tickers!")
	}

	sTwo, err := Get("AMZN")

	if err != nil {
		t.Errorf("Error fetching tickers!")
	}

	if sOne == sTwo {
		t.Errorf("Comparison between AAPL and AMZN is incorrect!")
	} else if sOne.Symbol == "" {
		t.Errorf("Valid ticker has zero value!")
	}

	cOne, err := sOne.GetCompany()

	if err != nil {
		t.Errorf("Error fetching company information!")
	}

	if cOne == (Company{}) {
		t.Errorf("Valid ticker returns empty company struct")
	}
}

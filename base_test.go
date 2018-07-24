package goiex

import (
	"testing"
)

func TestFetch(t *testing.T) {
	s_one, err := Get("AAPL")
   
   if err != nil {
      t.Errorf("Error fetching tickers!")
   }
   
	s_two, err := Get("AMZN")

   if err != nil {
      t.Errorf("Error fetching tickers!")
   }

	if s_one == s_two {
		t.Errorf("Comparison between AAPL and AMZN is incorrect!")
	} else if s_one.Symbol == "" {
		t.Errorf("Valid ticker has zero value!")
	}
}

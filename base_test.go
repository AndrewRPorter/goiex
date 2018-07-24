package goiex

import (
	"github.com/AndrewRPorter/goiex"
	"testing"
)

func TestFetch(t *testing.T) {
	s_one, err := goiex.Get("AAPL")
   
   if err != nil {
      t.Errorf("Error fetching tickers!")
   }
   
	s_two, err := goiex.Get("AMZN")

   if err != nil {
      t.Errorf("Error fetching tickers!")
   }

	if s_one == s_two {
		t.Errorf("Comparison between AAPL and AMZN is incorrect!")
	} else if s_one.Symbol == "" {
		t.Errorf("Valid ticker has zero value!")
	}
}

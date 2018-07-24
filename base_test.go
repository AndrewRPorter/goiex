package goiex

import (
   "testing"
   "github.com/AndrewRPorter/goiex"
)

func TestFetch(t *testing.T) {   
   s_one := goiex.Get("AAPL")
   s_two := goiex.Get("AMZN")
   
   if s_one == s_two {
      t.Errorf("Comparison between AAPL and AMZN is incorrect!")
   } else if s_one.Symbol == "" {
      t.Errorf("Valid ticker has zero value!")
   }
}

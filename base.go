package goiex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Symbol          string
	CompanyName     string
	PrimaryExchange string
	Sector          string
	Open            float64
	Close           float64
	High            float64
	Low             float64
	LatestVolume    int
	LatestPrice     float64
	Change          float64
	ChangePercent   float64
	AvgTotalVolume  int
	MarketCap       int
	PeRatio         float64
	Week52High      float64
	Week52Low       float64
	YtdChange       float64
}

func (s *Stock) GetSymbol() string {
	return s.Symbol
}

func (s *Stock) GetName() string {
	return s.CompanyName
}

func (s *Stock) GetExchange() string {
	return s.PrimaryExchange
}

func (s *Stock) GetSector() string {
	return s.Sector
}

func (s *Stock) GetOpen() float64 {
	return s.Open
}

func (s *Stock) GetClose() float64 {
	return s.Close
}

func (s *Stock) GetVolume() int {
	return s.LatestVolume
}

func (s *Stock) GetPrice() float64 {
	return s.LatestPrice
}

func (s *Stock) GetChange() float64 {
	return s.Change
}

func (s *Stock) GetPercentChange() float64 {
	return s.ChangePercent
}

func (s *Stock) GetAverageVolume() int {
	return s.AvgTotalVolume
}

func (s *Stock) GetMarketCap() int {
	return s.MarketCap
}

func (s *Stock) GetRatio() float64 {
	return s.PeRatio
}

func (s *Stock) GetYearHigh() float64 {
	return s.Week52High
}

func (s *Stock) GetYearLow() float64 {
	return s.Week52Low
}

func (s *Stock) GetYearChange() float64 {
	return s.YtdChange
}

func Get(ticker string) (Stock, error) {
	url := fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/quote", ticker)
	resp, err := http.Get(url)

	if err != nil {
		return Stock{}, err
	}

	s := Stock{}

	err = json.NewDecoder(resp.Body).Decode(&s)

	if err != nil {
      return Stock{}, fmt.Errorf("Unknown ticker %s", ticker)
	}

	return s, nil
}

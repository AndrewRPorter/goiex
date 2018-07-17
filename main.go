package main

import (
   "fmt"
   "log"
   "net/http"
   "encoding/json"
)

type Stock struct{
   Symbol string
   CompanyName string
   PrimaryExchange string
   Sector string
   Open float64
   OpenTime int
   Close float64
   CloseTime int
   High float64
   Low float64
   LatestVolume int
   IexRealtimePrice float64
   Change float64
   ChangePercent float64
   AvgTotalVolume int
   MarketCap int
   PeRatio float64
   Week52High float64
   Week52Low float64
   YtdChange float64
}

func get(ticker string) Stock {
   url := fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/quote", ticker)
   resp, err := http.Get(url)

   if err != nil {
      log.Fatalln("Unable to make request: ", err)
   }

   s := Stock{}

   err = json.NewDecoder(resp.Body).Decode(&s)

   if err != nil {
      log.Fatalln("Unable to decode JSON: ", err)
   }

   return s
}

func (s *Stock) get_symbol() string {
   return s.Symbol
}

func (s *Stock) get_company_name() string {
   return s.CompanyName
}

func (s *Stock) get_primary_exchange() string {
   return s.PrimaryExchange
}

func (s *Stock) get_sector() string {
   return s.Sector
}

func (s *Stock) get_open() float64 {
   return s.Open
}

func (s *Stock) get_open_time() int {
   return s.OpenTime
}

func (s *Stock) get_close() float64 {
   return s.Close
}

func (s *Stock) get_close_time() int {
   return s.CloseTime
}

func (s *Stock) get_volume() int {
   return s.LatestVolume
}

func (s *Stock) get_price() float64 {
   return s.IexRealtimePrice
}

func (s *Stock) get_change() float64 {
   return s.Change
}

func (s *Stock) get_percent_change() float64 {
   return s.ChangePercent
}

func (s *Stock) get_average_volume() int {
   return s.AvgTotalVolume
}

func (s *Stock) get_market_cap() int {
   return s.MarketCap
}

func (s *Stock) get_ratio() float64 {
   return s.PeRatio
}

func (s *Stock) get_52_high() float64 {
   return s.Week52High
}

func (s *Stock) get_52_low() float64 {
   return s.Week52Low
}

func (s *Stock) get_year_change() float64 {
   return s.YtdChange
}

func main() {
   s := get("GEVO")

   fmt.Println(s.get_symbol())
   fmt.Println(s.get_price())
}

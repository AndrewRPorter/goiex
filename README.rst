=====
goiex
=====

A go interface for accessing IEX finanical information. Data provided for free 
by `IEX <https://iextrading.com/developer/>`_. View `IEX’s Terms of Use 
<https://iextrading.com/api-exhibit-a/>`_.


Installation
------------

.. code::
   
      go get github.com/AndrewRPorter/goiex

Usage
-----

.. code:: go

   package main
   
   import (
      "fmt"
      "github.com/AndrewRPorter/goiex"
   )

   func main() {   
      s, err := goiex.Get("AAPL")
      
      if err != nil {
         fmt.Errorf("Unable to fetch ticker: %s", "AAPL")
      }
      
      fmt.Println(s.GetPrice())
   }
   
**Available Methods**

- ``Get()``
- ``GetSymbol()``
- ``GetName()``
- ``GetExchange()``
- ``GetSector()``
- ``GetOpen()``
- ``GetClose()``
- ``GetVolume()``
- ``GetPrice()``
- ``GetChange()``
- ``GetPercentChange()``
- ``GetAverageVolume()``
- ``GetMarketCap()``
- ``GetRatio()``
- ``GetYearHigh()``
- ``GetYearLow()``
- ``GetYearChange()``

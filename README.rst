.. image:: https://travis-ci.com/AndrewRPorter/goiex.svg?branch=master
    :target: https://travis-ci.com/AndrewRPorter/goiex
    
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
      
      fmt.Println(s.Price)
   }
   
**Available Methods**

- ``Get()``

**Available Fields**

- ``Symbol``
- ``CompanyName``
- ``Exchange``
- ``Sector``
- ``Open``
- ``Close``
- ``High``
- ``Low``
- ``Volume``
- ``Price``
- ``Change``
- ``ChangePercent``
- ``AverageVolume``
- ``MarketCap``
- ``PeRatio``
- ``YearHigh``
- ``YearLow``
- ``YtdChange``

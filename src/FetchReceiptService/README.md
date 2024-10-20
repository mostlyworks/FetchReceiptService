# FetchReceiptService

Created for [fetch](https://github.com/fetch-rewards/receipt-processor-challenge) backend coding challenge.


## Running

Assuming the Go sdk is installed:
```bash
go run main.go
```
Should get everything running.

Assuming the Go sdk isn't installed:

Go [here](https://go.dev/doc/install) install it and go back to the other method.

## Configuration
Contained in the /config directory are configuration json. The databaseConfig is currently ignored as no database is setup to run with this service (in Memory Mode!).

Why did you add configuration?

Future proofing.

It made sense for long term use if used to be able to manipulate the point calculations with out having to put out a new release with changes all over. Rather just modify the json.

### Points Configuration

Configuration of points can be done via editing the json example below with defaults:
```json
{
    "totalRoundedPoints": 50,
    "totalMutiplePoints": 25,
    "totalMutiple": 0.25,
    "totalRoundMod": 1.00,
    "itemCountPoints": 5,
    "itemCountDivsor": 2,
    "itemDescriptionMutiple": 3,
    "itemDescriptionPriceMutiplier": 0.2,
    "priceMutiplierRoundingPoints": 0,
    "purchaseDatePoints": 6,
    "purchaseDateCheckMod": 2,
    "purchaseTimeLowerBound": 14,
    "purchaseTimeUpperBound": 16,
    "purchaseTimePoints": 10,
    "dateExpectedFormat": "2006-01-02",
    "timeExpectedFormat": "15:04",
    "retailerNamePointMutiplier": 1,
    "defaultPointReturn": 0
}
```
If missing or failed to be read in it will automatically default the configuration to these values. 

***NOTE:*** Some tests will not pass if not using default values (Mainly due to hardcoded assertions)

### HTTP(s) Configuration

Configuration of the HTTP server can be done via the http config json, example below with defaults:
```json
{
    "useHttps": false,
    "port": 8080,
    "certFile": "",
    "keyFile": ""
}
```
If missing or failed to be read in it will automatically default the configuration to these values.

Validation around this isn't super strong, So don't try to break it or you will.
Expected inputs for cert and key are filenames to be loaded.
Port changes should cause little issue.
If Https is true and no cert or key are provided, it will default to useing http.

## Point Calcuation
I set this up to calcuate points on submission with the mindset that depending on the point offers this could change over time. So it would make more sense to run it when recieved rather than retrieval as that could cause it to change later (increase or decrease point counts) if things were later modified.

## Tests
Unit tests are provided for point calculations. End to end testing of the full expected cycle would have been good to have, but is not currently impossible to manually test.

## Fun Notes
I haven't touched Go in about 3 years so this was a fun shake off the dust project. It would probably be cleaner and better in Java or node but this was fun to re-learn some stuff.
## JSON Data

Typically JSON is used when ease-of-use is the primary goal of data interchange. Because JSON is human-readable, it is easy to debug if something breaks.  Many APIs and datastores/caches represent data in JSON, and, without a doubt, you will have to work with JSON as a data scientist.  

## Notes

* Support for Decoding and Encoding JSON are provided by the standard libary.
* This package gets better and better with every release.

## Links

[Decode JSON documents in Go](http://www.goinggo.net/2014/01/decode-json-documents-in-go.html)  
[Go walkthrough encoding/json](https://medium.com/@benbjohnson/go-walkthrough-encoding-json-package-9681d1d37a8f#.22rr9e3w4)

## Code Review

[Unmarshal JSON data from the Citibike API](example1/example1.go)   
[Marshal data and save a JSON file](example2/example2.go)   

## Exercises

### Exercise 1

Convert the station data in from the Citibike statin status API from JSON to CSV.  Unmarshal the station statuses from the API and save the station ID along with corresponding integer counts to a CSV file with a header.

[Template](exercises/template1/template1.go) | 
[Answer](exercises/exercise1/exercise1.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

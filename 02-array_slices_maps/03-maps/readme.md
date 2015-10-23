## Maps - Arrays, Slices and Maps

Maps provide a data structure that allow for the storage and management of key/value pair data.

## Notes

* Maps provide a way to store and retrieve key/value pairs.
* The map key must be a value that can be used in an assignment statement.
* Iterating over a map is always random.

## Links

http://blog.golang.org/go-maps-in-action

http://www.goinggo.net/2013/12/macro-view-of-map-internals-in-go.html

## Code Review

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](https://play.golang.org/p/wVgTXEVimA))

[Map literal initialization](example2/example2.go) ([Go Playground](https://play.golang.org/p/C0RJU7WUca))

[Map key restrictions](example3/example3.go) ([Go Playground](http://play.golang.org/p/FcY_0ckwOZ))

## Advanced Code Review

[Composing maps of maps](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/mycosI0zpN))

[Properties of nil maps](advanced/example2/example2.go) ([Go Playground](http://play.golang.org/p/GF0gbY4SvN))

## Exercises

### Exercise 1

Declare and make a map of integer values with a string as the key. Populate the map with five values and iterate over the map to display the key/value pairs.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/-JBSUoux-v)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/9DDe_wFFYi))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

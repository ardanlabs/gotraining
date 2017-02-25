## Packaging Design

Packages contain the basic unit of compiled code. They define a scope for the identifiers that are declared within them. Learning how to package our code is vital because exported identifiers become part of the packages API. Stable and usable API's are incredibly important.

## Links

https://www.goinggo.net/2017/02/design-philosophy-on-packaging.html  
https://www.goinggo.net/2017/02/package-oriented-design.html  

## An Interview with Brian Kernighan

http://www.cs.cmu.edu/~mihaib/kernighan-interview/index.html

_Can you tell us about the worse features of C, from your point of view?_

_I think that the real problem with C is that it doesn't give you enough mechanisms for structuring really big programs, for creating ``firewalls'' within programs so you can keep the various pieces apart. It's not that you can't do all of these things, that you can't simulate object-oriented programming or other methodology you want in C. You can simulate it, but the compiler, the language itself isn't giving you any help._ - July 2000

## Design Review

* Learn about the [design guidelines](../../#package-oriented-design) for packaging.

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

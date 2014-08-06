## Go Course Outline
This is a 3 day, 21 hour bootcamp style course for existing developers who are looking to gain a working understanding of Go.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

[Click Here To Review The Course Material](00-slides/readme.md)

__Minimal Qualified Student:__

* Programming
	* Studied CS in school or has a minimum of two years of experience programming full time professionally.
	* Familiar with structural and object oriented programming styles.
	* Has worked with arrays, lists, queues and stacks.
	* Understands processes, threads and synchronization at a high level.
	* Operating Systems
    	* Has worked with a command shell.
    	* Knows how to maneuver around the file system.
    	* Understands what environment variables are.

__What a student is expected to learn:__

* Understand what is available in the language.
    * Get a feel for writing code in an idiomatic style and syntax.
    * Learn how to write maintainable and solid programs that are production ready.
    * Walk away with patterns and techniques for solving common problems. 

### Day 1
On this day we take our initial tour of the language. We learn about variables, types, data structures, interfaces and concurrency. We also explore what is idiomatic and how the language is very orthogonal. This includes following the community standards for coding and style.

* Setting Up Your Environment ( Pre-Hour )
	* Installing Go
	* Installing Editor
	* Github Account
	* Git, Hg
* Language and Syntax
	* Variables
	* Struct Types
	* Constants
	* Pointers
	* Named Types
	* Functions
* Arrays, Slices and Maps
	* Arrays
	* Slices
	* Maps
* Methods, Interfaces and Embedding
	* Methods
	* Interfaces
	* Embedding Types
* Concurrency and Channels
	* Scheduler and Goroutines
	* Race Conditions
	* Channels

### Day 2
On this day we take go deeper into Go. We learn about testing, packaging, logging, encoding, io and reflection.

* Testing
	* Standard testing
	* Table tests
	* Benchmarking
* Packaging
	* Naming conventions
	* Go tooling
	* Exporting / Unexporting identifiers
* Standard Library
	* Logging
	* Encoding
	* Readers and Writers
* Reflection
	* Empty Interface
	* Reflect struct types and tags
	* Decoding values into types

### Day 3
On this day we talk about more advanced topics and learn how to use Go to build a full application. The program implements functionality that can be found in many Go programs being developed today. The program pulls different data feeds from the web and compares the content against a search term. The content that matches is then displayed to the terminal window. The program reads text files, makes web calls, decodes both XML and JSON into struct type values and finally does all of this using Go concurrency to make things fast.

* Project Setup
	* Version Control
	* Workspace
* Program Architecture
	* Review and Discuss
* DVCS, Packaging and Go Toolset
	* Repositories
		* Creating and naming
	    * Licensing
	    * Readme files
	* Packages
	    * Organizing and Naming
		* Import Form/Syntax
	* Go Tools
	    * Get, Vet, Format, Doc
* Program
	* Main Package
		* Provides the initialization and entry point for the program. 
	* Data Package
		* Contains a JSON data file that provides the program the set of URL’s to feeds that will be retrieved and processed.
	* Search Package
		* The core frameworks for the search engine are implemented in the package. The framework leverages interfaces to provide a plugin infrastructure to extend the types of content that can be retrieved and searched.
	* RSS Package
		* Implements a plugin that can retrieve and search RSS content.
* Tests
	* Search Package
		* Tests for validating the search frameworks and plugin environment.
	* RSS Package
		* Tests for validating the plugin can retrieve and search RSS content.
* Profiling
	* Run basic profiling to test the performance.
* Race Detection
	* Run the race detector.

___
[![GoingGo Training](00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](00-slides/images/ggb_logo.png)](http://www.goinggo.net)
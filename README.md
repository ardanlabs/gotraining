## Go Course Outline

This is a 2 day, 15 hour bootcamp style course for existing developers who are looking to gain a working understanding of Go. 

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

On this day we take a tour of the language. We learn the keywords, built-in functions and syntax. We also explore what is idiomatic and how the language is very orthogonal. This includes following the community standards for coding and style.

* Setting Up Your Environment ( Pre-Hour )
	* Installing Go
	* Installing Editor
	* Github Account
	* Git, Hg, Bzr
* Language and Syntax
	* Variables
	* Keyword var
	* Short variable declaration operator
	* Conversions
	* Exercises
* Struct Types
	* Declare and initialize
	* Exercises
* Constants
	* Declare and initialize
	* Parallel type system
	* Exercises
* Pointers
	* Pass by value
	* Sharing data
	* Exercises
* Named Types
	* Declare and initialize
	* Support with constants
	* Exercises
* Functions
	* Return multiple values
	* Blank identifiers
	* Trapping panics
	* Variadic functions
	* Exercises
* Packaging
	* Naming conventions
	* Exporting / Unexporting identifiers
	* Unexported value access
	* Exporting / Unexporting fields
	* Embedded types
	* Exercises
* Arrays, Slices and Maps
	* Arrays
		* Declaring, initializing and assignments
		* Unique types
		* Iteration
		* Exercises
	* Slices
		* Declaring, initializing and assignments
		* Length vs. Capacity
		* Reference type
		* Slicing
		* Appending values
		* Iteration
		* Three index slices
		* Using with functions
		* Exercises
	* Maps
		* Declaring, initializing and assigning
		* Restrictions on key values
		* Composition
		* Exercises
* Methods, Interfaces and Embedding
	* Methods
		* Declaring methods
		* Receivers
		* Exercises
	* Interfaces
		* Declaring
		* Implementing
		* Exercises
	* Embedding Types
		* Declaring
		* Use with interface
		* Promotion
		* Exercises
* Standard Library
	* Logging
		* Configuration
		* Writing messages
		* Exercises
	* Encoding
		* Unmarshaling JSON
		* Working with files
		* Marshaling
		* Exercises
	* Readers and Writers
		* Writing buffers
		* Web requests
		* Multi Writers
		* Exercises
* Reflection
	* Empty Interface
	* Reflect struct types and tags
	* Decoding values into types
	* Exercises
* Concurrency and Channels
	* Scheduler and Goroutines
		* Create goroutines
		* Concurrency
		* Parallelism
		* Exercises
	* Race Conditions
		* Atomic Operations
		* Mutexes
		* Exercises
	* Channels
		* Unbuffered
		* Buffered
		* Exercises
	* Testing
		* Standard tests
		* Table tests
		* Benchmarking
		* Exercises
* Advanced Topics
	* Advanced code samples for above topics

### Day II

On this day we learn how to use Go to build a full application. The program implements functionality that can be found in many Go programs being developed today. The program pulls different data feeds from the web and compares the content against a search term. The content that matches is then displayed to the terminal window. The program reads text files, makes web calls, decodes both XML and JSON into struct type values and finally does all of this using Go concurrency to make things fast.

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

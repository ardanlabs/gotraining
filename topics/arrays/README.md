## Arrays

Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory. Arrays have some special features in Go related to how they are declared and viewed as types.

## Notes

* If you don't understand the data, you don't understand the problem.
* If you don't understand the cost of solving the problem, you can't reason about the problem.
* If you don't understand the hardware, you can't reason about the cost of solving the problem.
* Arrays are fixed length data structures that can't change.
* Arrays of different sizes are considered to be of different types.
* Memory is allocated as a contiguous block.

## CPU Caches
This content is provided by Scott Meyers from his talk in 2014 at Dive:

[CPU Caches and Why You Care (18:50-20:30)](https://youtu.be/WDIkqP4JbkE?t=1129)  
[CPU Caches and Why You Care (44:36-45:40)](https://youtu.be/WDIkqP4JbkE?t=2676)

## CPU Cache Notes

* CPU caches works by caching main memory on cache lines.
* Cache lines today are either 32 or 64 bytes wide depending on the hardware.
* Cores do not access main memory directly. They tend to only have access their local caches.
* Both data and instructions are stored in the caches.
* Cache lines are shuffled down L1->L2->L3 as new cache lines need to be stored in the caches.
* Hardware likes to traverse data and instructions linearly along cache lines.
* Main memory is built on relatively fast cheap memory. Caches are built on very fast expensive memory.

* Access to main memory is incredibly slow, we need the cache.
	* Accessing one byte from main memory will cause an entire cache line to be read and cached.
	* Writes to one byte in a cache line requires the entire cache line to be written.

* Small = Fast
	* Compact, well localized code that fits in cache is fastest.
	* Compact data structures that fit in cache are fastest.
	* Traversals touching only cached data is the fastest.

* Predictable access patterns matter.
	* Whenever it is practical, you want to employ a linear array traversal.
	* Provide regular patterns of memory access.
	* Hardware can make better predictions about required memory.

* Cache misses can result in TLB cache misses as well.
	* Cache of translations of a virtual address to a physical address.
	* Waiting on the OS to tell us where the memory is.

### Cache Hierarchies
This is a diagram showing the relationship of the cache hierarchy for the 4 Core i7-9xx processor. The caches in the diagram are not to scale. This processor has four cores and each core has two hardware threads. The hardware threads per core share the Level 1 caches. The cores have individual Level 1 and Level 2 caches. All cores for all the processor share the L3 cache.

![figure1](figure1.png)

This is subject to be different in different processors. For this content, the following is the multi-levels of cache associated with the Intel 4 Core i7-9xx processor:

	L1 - 64KB Cache (Per Core)
		32KB I-Cache
		32KB D-Cache
		2 HW Threads
		4 cycles of latency

	L2 - 256KB Cache (Per Core)
		Holds both Instructions and Data
		2 HW Threads
		11 cycles of latency

	L3 - 8MB Cache
		Holds both Instructions and Data
		Shared across all 4 cores
		8 HW Threads
		39 cycles of latency

	Main Memory
		107 cycle of latency

## Links

[Scott Meyers: CPU Caches and Why You Care - Video](https://www.youtube.com/watch?v=WDIkqP4JbkE)

[Scott Meyers: CPU Caches and Why You Care - Deck](http://www.aristeia.com/TalkNotes/codedive-CPUCachesHandouts.pdf)

[What Every Programmer Should Know About Memory](http://www.akkadia.org/drepper/cpumemory.pdf)

[How CPU Caches Work and Why](http://www.extremetech.com/extreme/188776-how-l1-and-l2-cpu-caches-work-and-why-theyre-an-essential-part-of-modern-chips)

[Ulrich Drepper - Memory part 2: CPU caches](http://lwn.net/Articles/252125)

[The Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm)

[Dick Sites - "Data Center Computers: Modern Challenges in CPU Design](https://m.youtube.com/watch?feature=youtu.be&v=QBu2Ae8-8LM)

[Wirth's Law](https://en.wikipedia.org/wiki/Wirth%27s_law)

[Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc)

[Pitfalls of OOP](http://harmful.cat-v.org/software/OO_programming/_pdf/Pitfalls_of_Object_Oriented_Programming_GCAP_09.pdf)

[Why you should avoid Linked Lists](https://www.youtube.com/watch?v=YQs6IC-vgmo)

## Code Review

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](http://play.golang.org/p/2D24t6fbW_))

[Different type arrays](example2/example2.go) ([Go Playground](http://play.golang.org/p/nYgwqqOctt))

[Contiguous memory allocations](example3/example3.go) ([Go Playground](http://play.golang.org/p/L-SmdGfUcP))

## Advanced Code Review

[Contiguous Backing Array](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/IIubKW34GA))

## Exercises

### Exercise 1

Declare an array of 5 strings with each element initialized to its zero value. Declare a second array of 5 strings and initialize this array with literal string values. Assign the second array to the first and display the results of the first array. Display the string value and address of each element.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/qKUNW0FSgC)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/px3X_qw_SX))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
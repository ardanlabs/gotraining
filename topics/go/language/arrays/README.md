## Arrays

Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory. Arrays have some special features in Go related to how they are declared and viewed as types.

## Notes

* If you don't understand the data, you don't understand the problem.
* If you don't understand the cost of solving the problem, you can't reason about the problem.
* If you don't understand the hardware, you can't reason about the cost of solving the problem.
* Arrays are fixed length data structures that can't change.
* Arrays of different sizes are considered to be of different types.
* Memory is allocated as a contiguous block.
* Go gives you control over spacial locality.

## Design Guidelines

* Learn about the [design guidelines](../../#data-oriented-design) for data oriented design.

## CPU Caches

[CPU Caches and Why You Care (18:50-20:30)](https://youtu.be/WDIkqP4JbkE?t=1129) - Scott Meyers  
[CPU Caches and Why You Care (44:36-45:40)](https://youtu.be/WDIkqP4JbkE?t=2676) - Scott Meyers   
[Performance Through Cache-Friendliness (4:25-5:48)](https://youtu.be/jEG4Qyo_4Bc?t=266) - Damian Gryski  

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

	3GHz(3 clock cycles/ns) * 4 instructions per cycle = 12 instructions per ns!

	L1 - 64KB Cache (Per Core)
		32KB I-Cache
		32KB D-Cache
		2 HW Threads
		4 cycles of latency
		Stalls for 16 instructions or 1.3 ns

	L2 - 256KB Cache (Per Core)
		Holds both Instructions and Data
		2 HW Threads
		11 cycles of latency
		Stalls for 44 instructions or 3.6 ns

	L3 - 8MB Cache
		Holds both Instructions and Data
		Shared across all 4 cores
		8 HW Threads
		39 cycles of latency
		Stalls for 156 instructions or 13 ns

	Main Memory
		107 cycle of latency
		Stalled for 428 instructions or 35.6 ns
		27 times slower!

### Latencies

```
L1 cache reference ......................... 0.5 ns
Branch mispredict ............................ 5 ns
L2 cache reference ........................... 7 ns
Mutex lock/unlock ........................... 25 ns
Main memory reference ...................... 100 ns             
Compress 1K bytes with Zippy ............. 3,000 ns  =   3 µs
Send 2K bytes over 1 Gbps network ....... 20,000 ns  =  20 µs
SSD random read ........................ 150,000 ns  = 150 µs
Read 1 MB sequentially from memory ..... 250,000 ns  = 250 µs
Round trip within same datacenter ...... 500,000 ns  = 0.5 ms
Read 1 MB sequentially from SSD* ..... 1,000,000 ns  =   1 ms
Disk seek ........................... 10,000,000 ns  =  10 ms
Read 1 MB sequentially from disk .... 20,000,000 ns  =  20 ms
Send packet CA->Netherlands->CA .... 150,000,000 ns  = 150 ms
```

![Cache Latencies Image](cache_latencies_graph.png)

#### CPU Caches / Memory

[CPU Caches and Why You Care - Video](https://www.youtube.com/watch?v=WDIkqP4JbkE) - Scott Meyers  
[CPU Caches and Why You Care - Deck](http://www.aristeia.com/TalkNotes/codedive-CPUCachesHandouts.pdf) - Scott Meyers  
[Mythbusting Modern Hardware to Gain 'Mechanical Sympathy'](https://www.youtube.com/watch?v=MC1EKLQ2Wmg) - Martin Thompson  
[What Every Programmer Should Know About Memory](http://www.akkadia.org/drepper/cpumemory.pdf) - Ulrich Drepper  
[How CPU Caches Work and Why](http://www.extremetech.com/extreme/188776-how-l1-and-l2-cpu-caches-work-and-why-theyre-an-essential-part-of-modern-chips) - Joel Hruska  
[Modern Microprocessors A 90 Minute Guide](http://www.lighterra.com/papers/modernmicroprocessors) - Jason Robert Carey Patterson  
[Memory part 2: CPU caches](http://lwn.net/Articles/252125) - Ulrich Drepper  
[The Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm) - Herb Sutter  
[Data Center Computers: Modern Challenges in CPU Design](https://m.youtube.com/watch?feature=youtu.be&v=QBu2Ae8-8LM) - Dick Sites  
[Wirth's Law](https://en.wikipedia.org/wiki/Wirth%27s_law) - Wikipedia  
[Eliminate False Sharing](http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206) - Herb Sutter  
[NUMA Deep Dive Series](http://frankdenneman.nl/2016/07/06/introduction-2016-numa-deep-dive-series/) - Frank Denneman    
[The Myth Of Ram](http://www.ilikebigbits.com/blog/2014/4/21/the-myth-of-ram-part-i) - Emil Ernerfeldt  
[Understanding Transaction Hardware Memory](https://www.infoq.com/presentations/hardware-transactional-memory) - Gil Gene  
[Performance Through Cache-Friendliness (4:25-5:48)](https://youtu.be/jEG4Qyo_4Bc?t=266) - Damian Gryski  
[Going Nowhere Faster](https://www.youtube.com/watch?v=2EWejmkKlxs) - Chandler Carruth  

#### Data-Oriented Design

[Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc) - Mike Acton  
[Efficiency with Algorithms, Performance with Data Structures](https://www.youtube.com/watch?v=fHNmRkzxHWs) - Chandler Carruth  
[Taming the performance Beast](https://www.youtube.com/watch?v=LrVi9LHP8Bk) - Klaus Iglberger  

[Pitfalls of OOP](http://harmful.cat-v.org/software/OO_programming/_pdf/Pitfalls_of_Object_Oriented_Programming_GCAP_09.pdf) - Tony Albrecht  
[Why you should avoid Linked Lists](https://www.youtube.com/watch?v=YQs6IC-vgmo) - Bjarne Stroustrup  
[Data-Oriented Design (Or Why You Might Be Shooting Yourself in The Foot With OOP)](http://gamesfromwithin.com/data-oriented-design) - Noel    

## Code Review

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](https://play.golang.org/p/1KLhI3EZ7Id))  
[Different type arrays](example2/example2.go) ([Go Playground](https://play.golang.org/p/KmXH8Z9F3Rh))  
[Contiguous memory allocations](example3/example3.go) ([Go Playground](https://play.golang.org/p/QxnU1fO_f1e))  
[Range mechanics](example4/example4.go) ([Go Playground](https://play.golang.org/p/5n3bjld-Y0T))  

## Exercises

### Exercise 1

Declare an array of 5 strings with each element initialized to its zero value. Declare a second array of 5 strings and initialize this array with literal string values. Assign the second array to the first and display the results of the first array. Display the string value and address of each element.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/yURUYyZ5_hA)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/VtuGj8S-K2l))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

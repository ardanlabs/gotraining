## Design Guidelines

You must develop a design philosophy that establishes a set of guidelines. This is more important than developing a set of rules or patterns you apply blindly. Guidelines help to formulate, drive and validate decisions. You can't begin to make the best decisions without understanding the impact of your decisions. Every decision you make, every line of code you write comes with trade-offs.

### Prepare Your Mind

**Somewhere Along The Line**  
* We became impressed with programs that contain large amounts of code.
* We strived to create large abstractions in our code base.
* We forgot that the hardware is the platform.
* We lost the understanding that every decision comes with a cost.

**These Days Are Gone**  
* We can throw more hardware at the problem.
* We can throw more developers at the problem.

**Aspire To**  
* Be a champion for quality and efficiency.
* Be pro-active.
* Be fearless in the face of hard problems.
* Be responsible with expectations, time and resources.
* Have an insatiable curiosity.
* Have a point of view.
* Value good communication.
* Value introspection and self-review.
* Value simplicity.
* Give and receive feedback well.

**Open Your Mind**  
* Technology changes quickly but people's minds change slowly.
* Easy to adopt new technology but hard to adopt new ways of thinking.

### Performance vs Productivity

Performance and productivity both matter, but in the past you couldn’t have both. You needed to choose one over the other. We naturally gravitated to productivity, with the idea or hope that the hardware would resolve our performance problems for free. This movement towards productivity has resulted in the design of programming languages that produce sluggish software that is out pacing the hardware’s ability to make them faster.

By following Go’s idioms and a few guidelines, we can write code that can be reasoned about by anyone who looks at it. We can write software that simplifies, minimizes and reduces the amount of code we need to solve the problems we are working on. We don’t have to choose productivity over performance or performance over productivity anymore. We can have both.

**Quotes**

_"The hope is that the progress in hardware will cure all software ills. However, a critical observer may observe that software manages to outgrow hardware in size and sluggishness. Other observers had noted this for some time before, indeed the trend was becoming obvious as early as 1987." - Niklaus Wirth_

_"The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry." - Henry Petroski_

_"The hardware folks will not put more cores into their hardware if the software isn’t going to use them, so, it is this balancing act of each other staring at each other, and we are hoping that Go is going to break through on the software side.” - Rick Hudson_

_"C is the best balance I've ever seen between power and expressiveness. You can do almost anything you want to do by programming fairly straightforwardly and you will have a very good mental model of what's going to happen on the machine; you can predict reasonably well how quickly it's going to run, you understand what's going on .... - Brian Kernighan_

### Guidelines, Decision Making and Trade-Offs

Develop your design philosophy around these three major categories in this order: Integrity, Readability, Simplicity and then Performance. You must consciously and with great reason be able to explain the category you are choosing.

**_Note: There are exceptions to everything but when you are not sure an exception applies, follow the guidelines presented the best you can._** 

**1) Integrity**  
**_We need to become very serious about reliability._**

This is the accuracy and consistency of your code performing every read, write and the execution of every instruction. Just as important, it's knowing the error handling code is the main code. Nothing trumps integrity - EVER.

Case Study on Postmortems:  
48 critical failures were found in a study looking at a couple hundred bugs in Cassandra, HBase, HDFS, MapReduce, and Redis.  
* 92% : Failures from bad error handling
    * 35% : Incorrect handling
        * 25% : Simply ignoring an error
        * 8% : Catching the wrong exception
        * 2% : Incomplete TODOs
    * 57% System specific
        * 23% : Easily detectable
        * 34% : Complex bugs
* 8% : Failures from latent human errors

Ignorance vs Carelessness:  
Anytime we identify an integrity issue we need to ask ourselves why it happened. 
```
                    Not Deliberate               Deliberate
              ------------------------------------------------------
              |                          |                         |
              |                          |                         |
   Ignorance  |  Learning / Prototyping  |    Hacking / Guessing   |
              |                          |                         |
              |                          |                         |
              |-----------------------------------------------------
              |                          |                         |
              |                          |                         |
Carelessness  |        Education         |   Dangerous Situation   |
              |                          |                         |
              |                          |                         |
              ------------------------------------------------------
```
[Software Development for Infrastructure](http://www.stroustrup.com/Software-for-infrastructure.pdf) - Bjarne Stroustrup  
[Normalization of Deviance in Software](http://danluu.com/wat/) - danluu.com  
[Lessons learned from reading postmortems](http://danluu.com/postmortem-lessons/) - danluu.com  
[Technical Debt Quadrant](https://martinfowler.com/bliki/TechnicalDebtQuadrant.html) - Martin Fowler  

**2) Readability**   
**_We must structure our systems to be more comprehensible._**  
**_Readability means reliability._**

This is about writing simple code that is easy to read and understand without the need of mental exhaustion. Just as important, it's about not hiding the cost/impact of the code per line, function, package and the overall ecosystem it runs in.

[Example of classic readability issues](http://codepad.org/Xw7eUSSA)  

**3) Simplicity**  
**_We must understand that simplicity is hard to design and complicated to build._**  

This is about hiding complexity. A lot of care and design must go into simplicity because this can cause more problems then good. It can create issues with readability and it can cause issues with performance. Validate that abstractions are not generalized or even too concise. You might think you are helping the programmer or the code but validate things are still easy to use, understand, debug and maintain.

[Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM) - Rob Pike  

**4) Performance**  
**_We must compute less to get the results we need._**

This is about not wasting effort and achieving execution efficiency. Writing code that is mechanically sympathetic with the runtime, operating system and hardware. Achieving performance by writing less and more efficient code but staying within the idioms and framework of the language.

Rules of Performance:   
    * Never guess about performance.  
    * Measurements must be relevant.  
    * Profile before you decide something is performance critical.  
    * Test to know you are correct. 

Micro-Optimizations are about squeezing every ounce of performance as possible. When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic. You are writing clever code that may require the unsafe package or you may need to drop into assembly.

[Example of a micro optimization](https://play.golang.org/p/D_bImirgXL)

### Data-Oriented Design

[Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc) - Mike Acton  
[Efficiency with Algorithms, Performance with Data Structures](https://www.youtube.com/watch?v=fHNmRkzxHWs) - Chandler Carruth

* If you don't understand the data, you don't understand the problem.
* All problems are unique and specific to the data you are working with.
* Data transformations are at the heart of solving problems. Each function, method and work-flow must focus on implementing the specific data transformations required to solve the problems.
* If your data is changing, your problems are changing. When your problems are changing, the data transformations needs to change with it.
* Uncertainty about the data is not a license to guess but a directive to STOP and learn more.
* Solving problems you don't have, creates more problems you now do.
* If performance matters, you must have mechanical sympathy for how the hardware and operating system work.
* Minimize, simplify and REDUCE the amount of code required to solve each problem. Do less work by not wasting effort.
* Code that can be reasoned about and does not hide execution costs can be better understood, debugged and performance tuned.
* Coupling data together and writing code that produces predictable access patterns to the data will be the most performant.
* Changing data layouts can yield more significant performance improvements than changing just the algorithms.
* Efficiency is obtained through algorithms but performance is obtained through data structures and layouts.

### Interface and Composition Design

_With help from [Sandi Metz](https://twitter.com/sandimetz). and [Rob Pike](https://twitter.com/rob_pike)_

* Interfaces give programs structure.
* Interfaces encourage design by composition.
* Interfaces enable and enforce clean divisions between components.
    * The standardization of interfaces can set clear and consistent expectations.
* Interfaces allow you to group concrete types by what they do.
    * Don't group types by a common DNA but by a common behavior.
    * Everyone can work together when we focus on what we do and not who we are.
* Interfaces provide the highest form of decoupling when the concrete types used to implement them can remain opaque.
    * Decoupling means reducing the amount of intimate knowledge code must have about concrete types.
    * When dependencies are weakened and the coupling loosened, cascading changes are minimized and stability is improved.
* Interfaces help your code decouple itself from change.
    * You must do your best to understand what could change and use interfaces to decouple.
    * Interfaces with more than one method have more than one reason to change.
    * Uncertainty about change is not a license to guess but a directive to STOP and learn more.
    * Recognizing and minimizing cascading changes across the code is a way to architect adaptability and stability in your software.

### Interface Pollution

_With help from [Sarah Mei](https://twitter.com/sarahmei) and [Burcu Dogan](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275)_

* Don't use an interface for the sake of using an interface.
* Don't use an interface to generalize an algorithm.
* Unless the user needs to provide an implementation or you have multiple implementations, question.
* Don’t export any interface unless your user needs it. This includes interfaces for internal testing. Users can declare their own interfaces.
* If it's not clear how an abstraction makes the code better, it probably doesn't.

### Package-Oriented Design

* Start with a Project that contains all the source code you need to build the products and services the Project owns.
* Maintain each Project in a single repo.
* Only break a Project up when developer productivity is a cost.
* Breaking a Project into multiple projects comes with extra dependency costs.
* In many languages folders are used to organize code, in Go folders are used to organize API's (packages).
* Packages in Go provide API boundaries that should focus on solving one specific problem or a highly focused group of problems.
* You must understand how changes to the API for a particular package affects the other packages that depend on it.
* Recognizing and minimizing cascading changes across different packages is a way to architect adaptability and stability in your software.
* When dependencies between packages are weakened and the coupling loosened, cascading changes are minimized and stability is improved.

### Writing Concurrent Software

**_Note: This material is covered in detail in the classroom. This is a summary of the guidelines that are discussed._**

Concurrency is about managing multiple things at once. Like one person washing the dishes while they are also cooking dinner. You're making progress on both but you're only ever doing one of those things at the same time. Parallelism is about doing multiple things at once. Like one person cooking and placing dirty dishes in the sink, while another washes the dishes. They are happening at the same time.

Both you and the runtime have a responsibility in managing the concurrency of the application. You are responsible for managing these three things when writing concurrent software:

* The application must startup and shutdown with integrity.
    * Know how and when every goroutine you create terminates.
    * All goroutines you create should terminate before main returns.
    * Applications should be capable of shutting down on demand, even under load, in a controlled way.
        * You want to stop accepting new requests and finish the requests you have (load shedding).
* Identify and monitor critical points of back pressure that can exist inside your application.
    * Channels, mutexes and atomic functions can create back pressure when goroutines are required to wait.
    * A little back pressure is good, it means there is a good balance of concerns.
    * A lot of back pressure is bad, it means things are imbalanced.
    * Back pressure that is imbalanced will cause:
        * Failures inside the software and across the entire platform.
        * Your application to collapse, implode or freeze.
    * Measuring back pressure is a way to measure the health of the application.
* Rate limit to prevent overwhelming back pressure inside your application.
    * Every system has a breaking point, you must know what it is for your application.
    * Applications should reject new requests as early as possible once they are overloaded.
        * Don’t take in more work than you can reasonably work on at a time.
        * Push back when you are at critical mass. Create your own external back pressure.
    * Use an external system for rate limiting when it is reasonable and practical.
* Use timeouts to release the back pressure inside your application.
    * No request or task is allowed to take forever.
    * Identify how long users are willing to wait.
    * Higher-level calls should tell lower-level calls how long they have to run.
    * At the top level, the user should decide how long they are willing to wait.
    * Use the `Context` package.
        * Functions that users wait for should take a `Context`.
            * These functions should select on <-ctx.Done() when they would otherwise block indefinitely.
        * Set a timeout on a `Context` only when you have good reason to expect that a function's execution has a real time limit.
        * Allow the upstream caller to decide when the `Context` should be canceled.
        * Cancel a `Context` whenever the user abandons or explicitly aborts a call.
* Architect applications to:
    * Identify problems when they are happening.
    * Stop the bleeding.
    * Return the system back to a normal state.

### Channels

**_Note: This material is covered in detail in the classroom. This is a summary of the guidelines that are discussed._**

Channels allow goroutines to communicate with each other through the use of signaling semantics. Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.

* Use channels to orchestrate and coordinate goroutines.
    * Focus on the signaling semantics and not the sharing of data.
    * Signal by passing data or by closing the channel.
    * Question their use for synchronizing access to shared state.
        * _There are cases where channels can be simpler for this but initially question._
* Unbuffered channels:
    * Blocks the sending goroutine from moving forward until a different goroutine has received the data signal.
        * The sending goroutine gains immediate knowledge their data signal has been received.
    * Both goroutines involved must be at the channel at the same time.
        * More important: The Receive happens before the Send.
    * Trade-offs:
        * We take the benefit of knowing the data signal has been received for the cost of higher blocking latency.
* Buffered channels:
    * Does NOT block the sending goroutine from moving forward until a different goroutine has received the data signal.
        * The sending goroutine gains no knowledge that their data signal has been received.
    * Both goroutines involved don't need to be at the channel at the same time.
        * More important: The Send happens before the Receive.
    * Trade-offs:
        * We take the benefit of reducing blocking latency for the cost of not knowing if/when the data signal is received.
* Closing channels:
    * Signaling without the need for data passing.
    * Perfect for signaling cancellations and deadlines.
* NIL channels:
    * Turn off signaling
    * Perfect for rate limiting or short term stoppages.

### Channel Semantics

**_Note: This material is covered in detail in the classroom. This is a summary of the guidelines that are discussed._**

Depending on the problem you are solving, you may require different channel semantics. Depending on the semantics you need, different architectural choices must be taken.

* If any given Send on a channel `CAN` cause the goroutine to block:
    * You have less buffers compared to the number of goroutines that will perform a Send at any given time.
    * An example would be a resource pool of database connections.
    * This requires knowing what happens when the Send blocks because this will create a situation of back pressure inside the application in front of the channel.
    * The things discussed above about [writing concurrent software](https://github.com/ardanlabs/gotraining/blob/master/reading/design_guidelines.md#writing-concurrent-software) must be taken into account for this channel.
    * Not knowing what happens when the Send blocks on the channel is not a license to guess but a directive to STOP, understand and take appropriate measures.
* If any given Send on a channel `WON'T` cause the goroutine to block:
    * You have a buffer for every goroutine that will perform a Send.
    * You will abandon the Send immediately if it can't be performed.
    * An example would be a fan out pattern or pipelining.
    * This requires knowing if the size of the buffer is appropriate and if it is acceptable to abandon the Send.
* Less is more with buffers.
    * Don’t think about performance when thinking about buffers.
    * Buffers can help to reduce blocking latency between signaling.
        * Reducing blocking latency towards zero does not necessarily mean better throughput.
        * If a buffer of one is giving you good enough throughput then keep it.
        * Question buffers that are larger than one and measure for size.
        * Find the smallest buffer possible that provides good enough throughput.

### Code Reviews

I teach a lot about the things I look for in code reviews. I am slowly attempting to document this.

* Try to use functions over methods when it is practical. Functions allow for better readability and reusability because all the input is passed in and the output is returned out. No information is lost or abstracted.
* Eliminate the use of the else statements when it is practical. Do not attempt to push code paths to the end of a function. Keep your positive path code in the first tabbed position and use the if statement to process negative path. Return from the function as part of error handling.
* Don't start off with pointer variables if it can be avoided. It is easier to work with variables that represent a value, even if that value is going to escape to the heap. The use of the & operator can go a long way to maintaining readability in your code.
* Use the keyword var to represent the declaration of a variable that is being set to its zero value. This helps with readability and can provide the basis for developing a consistent set of rules around variable declarations. One of Go's biggest warts is there are too many ways to declare and create variables.

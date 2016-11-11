## Design Guidelines

Performance and productivity both matter, but in the past you couldn’t have both. You needed to choose one over the other. We naturally gravitated to productivity, with the idea or hope that the hardware would resolve our performance problems for free. This movement towards productivity has resulted in the design of programming languages that produce sluggish software that is out pacing the hardware’s ability to make them faster.

By following Go’s idioms and a few guidelines, we can write code that can be reasoned about by anyone who looks at it. We can write software that simplifies, minimizes and reduces the amount of code we need to solve the problems we are working on. We don’t have to choose productivity over performance or performance over productivity anymore. We can have both.

#### Quotes

_"The hope is that the progress in hardware will cure all software ills. However, a critical observer may observe that software manages to outgrow hardware in size and sluggishness. Other observers had noted this for some time before, indeed the trend was becoming obvious as early as 1987." - Niklaus Wirth_

_"The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry." - Henry Petroski_

_"The hardware folks will not put more cores into their hardware if the software isn’t going to use them, so, it is this balancing act of each other staring at each other, and we are hoping that Go is going to break through on the software side.” - Rick Hudson_

_"C is the best balance I've ever seen between power and expressiveness. You can do almost anything you want to do by programming fairly straightforwardly and you will have a very good mental model of what's going to happen on the machine; you can predict reasonably well how quickly it's going to run, you understand what's going on .... - Brian Kernighan_

#### Somewhere Along The Line

* We became impressed with programs that contain large amounts of code.
* We strived to create large abstractions in our code base.
* We forgot that the hardware is the platform.
* We lost the understanding that every decision comes with a cost.

#### These Days Are Gone

* We can throw more hardware at the problem.
* We can throw more developers at the problem.

#### Remember

The **compiler** is a tool and it's not all knowing or perfect.  
You need to work with it and help it.

The **operating system** is doing its best to keep cores busy and leverage the right core at the right time.  
You need to work with it and help it.

The **hardware** is doing its best to execute as many instructions per clock cycle as possible.  
You need to work with it and help it.

Writing idiomatic code in Go gets you a long way. But you can do so much more when you know how these things work.

#### Decision Making and Trade-Offs

You can't make the best decisions without understanding the impact of your decisions. Every decision you make, every line of code you write comes with trade-offs. Evaluate the trade-offs of a decision in this order: Integrity, Readability, and then Performance. Before you allow a lower priority trade-off to trump you must consciously and with great reason be able to explain yourself.

* A design philosophy is far more important than a set of rules or patterns to apply like stamps.
* Write code for today, design code for tomorrow.
* Design principals are not rules, they're guidelines.

**1) Integrity**  
This is the accuracy and consistency of your code performing every read, write and the execution of every instruction. Nothing trumps integrity - EVER.

_We need to become very serious about reliability._  

**2) Readability**    
This is about writing simple code that is easy to read and understand without the need of mental exhaustion. Just as important, it's about not hiding the cost/impact of the code per line, function, package and the overall ecosystem it runs in.

_We must structure our systems to be more comprehensible._

**3) Performance**    
This is about not wasting effort and achieving execution efficiency. Writing code that is mechanically sympathetic with the runtime, operating system and hardware. Achieving performance by writing less and more efficient code but staying within the idioms and framework of the language.

_We must compute less to get the results we need._

**4) Micro-Optimization**    
This is about squeezing every ounce of performance as possible. When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic. You are writing clever code that may require the unsafe package or you may need to drop into assembly.

_You can choose to ignore everything above if there is a good reason, but be aware of the risk of ignoring them._

**_Note: There are exceptions to everything but as we teach, let your experience be your guide. When you are not sure an exception applies, follow the guidelines presented the best you can._** 

#### Data-Oriented Design
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

#### Interface and Composition Design

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

#### Interface Pollution

_With help from [Sarah Mei](https://twitter.com/sarahmei) and [Burcu Dogan](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275)_

* Don't use an interface for the sake of using an interface.
* Don't use an interface to generalize an algorithm.
* Unless the user needs to provide an implementation or you have multiple implementations, question.
* Don’t export any interface unless your user needs it. This includes interfaces for internal testing. Users can declare their own interfaces.
* If it's not clear how an abstraction makes the code better, it probably doesn't.

#### Package-Oriented Design

* Start with a Project that contains all the source code you need to build the products and services the Project owns.
* Maintain each Project in a single repo.
* Only break a Project up when developer productivity is a cost.
* Breaking a Project into multiple projects comes with extra dependency costs.
* In many languages folders are used to organize code, in Go folders are used to organize API's (packages).
* Packages in Go provide API boundaries that should focus on solving one specific problem or a highly focused group of problems.
* You must understand how changes to the API for a particular package affects the other packages that depend on it.
* Recognizing and minimizing cascading changes across different packages is a way to architect adaptability and stability in your software.
* When dependencies between packages are weakened and the coupling loosened, cascading changes are minimized and stability is improved.

#### Writing Concurrent Software

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

#### Channels

**_Note: This material is covered in detail in the classroom. This is a summary of the guidelines that are discussed._**

Channels are for the orchestration and coordinate of goroutines. Channels accomplish this by synchronizing the sending and receiving of data between goroutines or by providing mechanisms to signal goroutines about events.

* Use channels to orchestrate and coordinate goroutines.
    * They can help to simplify and make code more readable over using mutexes and atomic functions.
    * They provide support for cancellation and deadlines.
    * Though slower than mutexes and atomic functions, if the code is simpler and fast enough use them.
        * Mutexes and atomic functions can make accessing shared state simpler.
* Unbuffered channels:
    * Provides a full guarantee that information being exchanged has been delivered to another goroutine.
    * Both goroutines making the exchange must be at the channel at the same time.
        * The Receive happens before the Send.
    * Provides the highest level of guarantee within channel communication.
        * Responsibility for the information immediately changes hands.
        * The information is still visible and actively being worked on.
    * Trade-offs:
        * We take the benefit of the full guarantee for the cost of higher blocking latency.
* Buffered channels:
    * If the Send/Receive can block:
        * Provides a weaker guarantee that information being exchanged has been delivered to another goroutine.
        * Both goroutines making the exchange don’t need to be at the channel at the same time.
            * The Send happens before the Receive.
        * If the Sender / Receiver care about the information:
            * Provides a lower level of guarantee within channel communication.
                * Responsibility for the information does not immediately changes hands.
                * Creates risk for the sending goroutine because of the responsibility it maintains while losing control of the information.
            * The information is hidden and waiting to be rediscovered.
        * If the Sender / Receiver `don't` care about the information:
            * The lost of the guarantee is not an issue.
            * Useful when creating pipelines that may be canceled or making replicated requests to backends.
            * The information is hidden but we don't care if it is rediscovered.
        * Trade-offs:
            * We take the benefit of reducing blocking latency for the cost of weaker guarantees.
        * Notes:
            * Adding the buffer is a semantic_change, not an optimization.
            * More care must be taken when designing, configuring and balancing the application.
            * Don't use buffered channels when you need application-level queues.
    * If the Send/Receive `can't` block:
        * You get the full guarantee without the latency cost.
        * A Fan Out pattern is an example of this.
* Less is more with buffers.
    * Don’t think about performance when thinking about buffers.
    * Buffers can help to reduce blocking latency in the exchange of information.
        * Reducing blocking latency towards zero does not necessarily mean better throughput.
        * If a buffer of one is giving you good enough throughput then keep it.
        * Question buffers that are larger than one and measure for size.
        * Find the smallest buffer possible that provides good enough throughput.

#### Code Reviews

I teach a lot about the things I look for in code reviews. I am slowly attempting to document this.

* Try to use functions over methods when it is practical. Functions allow for better readability and reusability because all the input is passed in and the output is returned out. No information is lost or abstracted.
* Eliminate the use of the else statements when it is practical. Do not attempt to push code paths to the end of a function. Keep your positive path code in the first tabbed position and use the if statement to process negative path. Return from the function as part of error handling.
* Don't start off with pointer variables if it can be avoided. It is easier to work with variables that represent a value, even if that value is going to escape to the heap. The use of the & operator can go a long way to maintaining readability in your code.
* Use the keyword var to represent the declaration of a variable that is being set to its zero value. This helps with readability and can provide the basis for developing a consistent set of rules around variable declarations. One of Go's biggest warts is there are too many ways to declare and create variables.

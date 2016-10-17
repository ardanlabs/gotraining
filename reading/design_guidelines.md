## Design Guidelines

Performance and productivity both matter, but in the past you couldn’t have both. You needed to choose one over the other. We naturally gravitated to productivity, with the idea or hope that the hardware would resolve our performance problems for free. This movement towards productivity has resulted in the design of programming languages that produce sluggish software that is out pacing the hardware’s ability to make them faster.

By following Go’s idioms and a few guidelines, we can write code that can be reasoned about by anyone who looks at it. We can write software that simplifies, minimizes and reduces the amount of code we need to solve the problems we are working on. We don’t have to choose productivity over performance or performance over productivity anymore. We can have both.

#### Quotes

*"The hope is that the progress in hardware will cure all software ills. However, a critical observer may observe that software manages to outgrow hardware in size and sluggishness. Other observers had noted this for some time before, indeed the trend was becoming obvious as early as 1987." - Niklaus Wirth*

*"The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry." - Henry Petroski*

*"The hardware folks will not put more cores into their hardware if the software isn’t going to use them, so, it is this balancing act of each other staring at each other, and we are hoping that Go is going to break through on the software side.” - Rick Hudson*

*"C is the best balance I've ever seen between power and expressiveness. You can do almost anything you want to do by programming fairly straightforwardly and you will have a very good mental model of what's going to happen on the machine; you can predict reasonably well how quickly it's going to run, you understand what's going on .... - Brian Kernighan*

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

**1) Integrity**  
This is the accuracy and consistency of your code performing every read, write and the execution of every instruction. Nothing trumps integrity - EVER.

**2) Readability**  
This is about writing simple code that is easy to read and understand without the need of mental exhaustion. Just as important, it's about not hiding the cost/impact of the code per line, function, package and the overall ecosystem it runs in.

**3) Performance**  
This is about not wasting effort and achieving execution efficiency. Writing code that is mechanically sympathetic with the runtime, operating system and hardware. Achieving performance by writing less and more efficient code but staying within the idioms and framework of the language.

**4) Micro-Optimization**  
This is about squeezing every ounce of performance as possible. When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic. You are writing clever code that may require the unsafe package or you may need to drop into assembly.

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

_With help from [Sandi Metz](https://twitter.com/sandimetz). and Rob Pike_

* Interfaces give programs structure.
* Interfaces encourage design by composition.
* Interfaces enable and enforce clean divisions between components.
    * The standardization of interfaces can set clear and consistent expectations.
* Interfaces allow you to group concrete types by what they do.
    * Don't group types by a common DNA but by a common behavior.
    * Everyone can work together when we focus when we do and not what we are.
* Interfaces provide the highest form of decoupling when the concrete types used to implement them can remain opaque.
    * Decoupling means reducing the amount of intimate knowledge code must have about concrete types.
    * When dependencies are weakened and the coupling loosened, cascading changes are minimized and stability is improved.
* Interfaces with more than one method has more than one reason to change.
    * You must do your best to understand what could change and decouple those aspects of your code.
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

#### Code Reviews

I teach a lot about the things I look for in code reviews. I am slowly attempting to document this.

* Try to use functions over methods when it is practical. Functions allow for better readability and reusability because all the input is passed in and the output is returned out. No information is lost or abstracted.
* Eliminate the use of the else statements when it is practical. Do not attempt to push code paths to the end of a function. Keep your positive path code in the first tabbed position and use the if statement to process negative path. Return from the function as part of error handling.
* Don't start off with pointer variables if it can be avoided. It is easier to work with variables that represent a value, even if that value is going to escape to the heap. The use of the & operator can go a long way to maintaining readability in your code.
* Use the keyword var to represent the declaration of a variable that is being set to its zero value. This helps with readability and can provide the basis for developing a consistent set of rules around variable declarations. One of Go's biggest warts is there are too many ways to declare and create variables.



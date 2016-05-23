## Design Guidelines

These are a set of design guidelines for data, interfaces, composition and packages. Please consider these thoughts when designing your own software.

#### Mike Acton : Data-Oriented Design
https://www.youtube.com/watch?v=rX0ItVEVjHc

* If you don't understand the data, you don't understand the problem.

* All problems are unqiue and specific to the data you are working with.

* Data transformations are at the heart of solving problems. Each function, method and workflow must focus on implementing the specific data transformations required to solve the problems.

* If your data is changing, your problems are changing. When your problems are changing, the data transformations needs to change with it.

* Uncertainty about the data is not a license to guess but a directive to STOP and learn more.

* Solving problems you don't have, creates more problems you now do.

* Minimize, simplfy and reduce the amount of code required to solve each problem. Code that can be reasoned about and does not hide execution costs can be better understood, debugged and performance tuned.

* If performance matters, you must have mechanical sympathy for how the hardware and operating system work.

* Coupling data together and writing code that produces predictable access patterns to the data will be the most performant.

#### Scott Myers : The Most Important Design Guideline
https://www.youtube.com/watch?v=5tg1ONG18H8

* Make Interfaces easy to use correctly and hard to use incorrectly.

* Principle of least astonishment.
	* Keep the expectation clear, allows users to guess correctly.
	* Take advantage of what people already know.
	* Behavior should maintain a level of expectation.

* Choose good names.
	* Names are the interface.
	* Give a lot of thought to the names you use.

* Be consistent.

* Document before using.
	* Test driven design.
	* This is identify problems early on.

* Try to minimize user mistakes with the API.
	* Trying to constrain values can create readability issues.
	* Minimize choices.

#### Interface and Composition Design

_With help from [Sandi Metz](https://twitter.com/sandimetz)._

* Interfaces provide the highest form of decoupling when the concrete types used to implement them can remain opaque.

* Decoupling means reducing the amount of intimate knowledge code must have about concrete types.

* Interfaces with more than one method has more than one reason to change.

* You must do your best to understand what could change and decouple those aspects of your code.

* Uncertainty about change is not a license to guess but a directive to STOP and learn more.

* Recognizing and minimizing cascading changes across the code is a way to architect adaptability and stability in your software.

* When dependencies are weakened and the coupling loosened, cascading changes are minimized and stability is improved.

* The standardization of interfaces can set clear and consistent expectations.

#### Interface Polution

_With help from [Sarah Mei](https://twitter.com/sarahmei) and [Burcu Dogan](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275)_

* Don't use an interface for the sake of using an interface.

* Unless the user needs to provide an implementation or you have multiple implementations, question.

* Don’t export any interfaces until your user needs it. Users can declare their own interfaces.

* Don't add an interface just for the sake of testing. API's are for users not tests.

* If it's not clear how an abstraction makes the code better, it probably doesn't.

#### Package-Oriented Design

* Start with a Project that contains all the source code you need to build the products and services the Project owns.

* Maintain each Project in a single repo.

* Only break a Project up when developer productivity is a cost.

* Breaking a Project into multiple projects comes with dependency costs.

* In many languages folders are used to organize code, in Go folders are used to organize API's (packages).

* Packages in Go provide API boundaries that should focus on solving one specific problem or a highly focused group of problems.

* You must understand how changes to the API for a particular package affects the other packages that depend on it.

* Recognizing and minimizing cascading changes across different packages is a way to architect adaptability and stability in your software.

* When dependencies between packages are weakened and the coupling loosened, cascading changes are minimized and stability is improved.


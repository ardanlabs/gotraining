## Ultimate Data

This is material for developers who have some experience with Go and statistics and want to learn how to work with data to generate valuable insights. We believe these classes are perfect for data analysts/scientists/engineers interested in working in Go or Go programmers interested in doing data analysis.

[Ultimate Data](../courses/data/README.md)

## Design Guidelines

To be effective as a data analyst, you must develop a design philosophy that establishes a set of guidelines for your work. This is more than developing a set of rules or patterns you apply blindly. Guidelines help to formulate, drive and validate decisions. You can't expect to generate valuable insights without understanding the impact of your implementation decisions. Every decision you make, every line of code you write comes with trade-offs.

### What is data analysis?

Data analysis transforms Datasets into **Insights** that have corresponding **Actions** and **Consequences**.

### Prepare your mind

Every data analytics or data science project must begin by considering the:

1. Insights
2. Actions
3. Consequences

Before and during any data analytics project, you must be able to answer the following questions:

- What insights do I want to generate?
- What actions are triggered by the insights?
- What are the consequences of those actions?
- What do the results need to contain to best represent the desired insights?
- What is the data required to produce a valid result?
- How will I measure the validity of results?
- Can the results be effectively conveyed to decision makers as insights?
- Am I confident in the results?

Remember, uncertainty is not a license to guess but a directive to stop.

### Order of Operations

Data analytics projects should follow these steps in this order:

1. Understand the insights, actions and consequences involved.
2. Understand the relevant data to be gathered and analyzed.
3. Gather and organize the relevant data.
4. Understand the expectations for determining valid results.
5. Determine the most interpretable process that can produce valid results.
6. Determine how you will test the validity of results.
7. Develop the determined process and tests.
8. Test the results and evaluate against your expectations.
9. Refactor as necessary.
10. Looks for ways to simplify, minimize and reduce.

When the results don’t meet the expectations, ask yourself if modifying the determined process or data might improve the validity of the results?  

- If YES, you must re-evaluate:
    - Are such modifications warranted?
    - Can the modification be tested against the expectations?
    - Do I need to increase complexity?
    - Have I tested the most simplistic and interpretable solutions first?
- If NO, you must re-evaluate:
    - Am I using the best determined process?
    - Am I using the best data?
    - Are my expectations unrealistic?

### Guidelines, Decision Making and Trade-Offs

Develop your design philosophy around these major categories in this order: Integrity, Value, Readability/Interpretability, and Performance. You must consciously and with great reason be able to explain the category you are choosing.

**_Note: There are exceptions to everything but when you are not sure an exception applies, follow the guidelines presented the best you can._**

**1) Integrity** - If data science/analysis generates insights from Datasets, a breakdown in integrity results in bad insights. These types of results impact people, and therefore, generating bad insights may cause irreparable damage to real people. Nothing trumps integrity - EVER.

Rules of Integrity:
- Error handling code is the main code.
- You must understand the data.
- Control the input and output of your processes.
- You must be able to reproduce results.

**2) Value** - Effort without actionable results is not valuable. Just because you can produce a result, does not mean the result contains value.

Rules of Value:
- If an action can not be taken based on a result, the result does not have value.
- If the impact of a result can not be measured, the result does not have value.

**3) Readability and Interpretability** - This is about writing simple analyses that are easy to read and understand without mental exhaustion. However, this is also about avoiding unnecessary data transformations and analysis complexity that hides:

- The cost/impact of individual steps of the analyses.
- The underlying purpose of the data transformations and analyses.

**4) Performance** - This is about making your analyses run as fast as possible and produce results that minimize a given measure of error.  When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic.  If increasing the accuracy, e.g., of a given result by 0.001% takes a significant increase in effort/complexity and doesn’t produce more value or differing actions, the effort Optimization/Efficiency is not warranted.

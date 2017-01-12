## Hypothesis Testing

While working as a data scientist, you will encounter the need to test many hypotheses.  For example, you may want to know if a change to a website produced more email registrations, or you may want to know if a new model is actually performing better than an old model. These hypotheses must must be testable for you to actually come to some conclusion. In other words, there must be some way of determining if your hypothesis explains an observation.

## Notes

- A "null hypothesis," usually denoted by H0, is the hypothesis that observations result purely from chance, the opposite of the alternative hypothesis.
- An "alternative hypothesis," usually denoted by H1, is the hypothesis that observations caused by some non-random process.
- A "p-value" measures the probability of observing a test statistic as extreme as the current test statistic. 
- Hypothesis tests are carried out assuming the null hypothesis is true.

## Links

[Hypothesis Testing - Penn State](https://onlinecourses.science.psu.edu/statprogram/node/136)    
[The four steps of hypothesis testing](http://mathworld.wolfram.com/HypothesisTesting.html)

## Code Review

[Chi-Squared Table](http://sites.stat.psu.edu/~mga/401/tables/Chi-square-table.pdf)   
[Calculate Expected Frequencies](example1/example1.go)   
[Calculate Chi-squared Test Statistic](example2/example2.go)  
[Output the Result of the Test](example3/example3.go)   

## Exercises

### Exercise 1

Design a test to determine if die is fair. Use the Chi-squared test statistic and a 5% level of significance.  Answer the following:

- What would you measure?
- How would you calculate the test statistic?
- What range of Chi-squared values would cause you to abandon your null hypothesis?

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

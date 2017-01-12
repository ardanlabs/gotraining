## Summary Measures

Despite a huge focus on machine learning in the context of data science, solid statistical analysis (e.g., via summary and aggregation) needs to be part of every data science project and, in all honestly, can provide more value (and quicker value) than sophisticated ML.  However, even in cases where more sophisticated ML is justified, data scientists must understand the statistics of their data to determine the validity of modeling techniques and to develop intuition about the data. 

## Notes

- The goal of summary statistics is to communicate as much information as possible about a series of observations as simply as possible.
- Summary statistics often include a measure of central tendency (e.g., a mean or median) and a measure of "spread" (e.g., variance or standard deviation).
- ALL DATA SCIENCE PROJECTS MUST INCLUDE SUMMARY STATISTICS along with and before any more sophisticated modeling.

## Links

[Stat Trek](http://stattrek.com/)  
[Khan Academy - Statistics](https://www.khanacademy.org/math/statistics-probability)  
[Bayesian Statistics](http://hbanaszak.mjr.uw.edu.pl/StatRozw/Books/Bolstad_2007_Introduction%20to%20Bayesian%20Statistics.pdf)  
[Elements of Statistical Learning](http://statweb.stanford.edu/~tibs/ElemStatLearn/)  

## Code Review

[github.com/gonum/stat docs](https://godoc.org/github.com/gonum/stat)  
[github.com/montanaflynn/stats docs](https://godoc.org/github.com/montanaflynn/stats)  
[github.com/gonum/floats docs](https://godoc.org/github.com/gonum/floats)   
[Mean, Mode, Median](example1/example1.go)  
[Min, Max, Range](example2/example2.go)  
[Variance, Standard Deviation](example3/example3.go)    
[Quantiles](example4/example4.go)  

## Exercises

### Exercise 1

Output central tendency and statisitcal dispersion (or "spread") measures together for all numeric features of the iris data set.  Looking at these measure together gives a quick snapshot of "what the data looks like" numerically.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

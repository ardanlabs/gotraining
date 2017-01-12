## Visualizing Distributions

Although summary statistics are very useful and should be utilized to describe data, visualizing distributions of data can provide very quick intuition about how to proceed in an analysis.  Tools like histograms and box plots provide a visual representation about the distribution of variables.   

## Example Histogram

![alt tag](histogram.png)    
from [here](https://statistics.laerd.com/statistical-guides/understanding-histograms.php)

## Example Box Plot

![alt tag](box.gif)    
from [here](http://www.physics.csbsju.edu/stats/box2.html)

## Notes

- A histogram is an estimate of the probability distribution of a continuous variable created by "binning" data together into a certain number of bins defined by variable ranges.
- A box plot also displays the distribution of data.  In this case, based on a five number summary: minimum, first quartile, median, third quartile, and maximum

## Links

[Stat Trek](http://stattrek.com/)  
[Khan Academy - Statistics](https://www.khanacademy.org/math/statistics-probability)  
[Bayesian Statistics](http://hbanaszak.mjr.uw.edu.pl/StatRozw/Books/Bolstad_2007_Introduction%20to%20Bayesian%20Statistics.pdf)  
[Elements of Statistical Learning](http://statweb.stanford.edu/~tibs/ElemStatLearn/)  

## Code Review

[github.com/gonum/plot docs](https://godoc.org/github.com/gonum/plot)  
[Histogram of a normal distribution](example1/example1.go)  
[Histograms with real data](example2/example2.go)  
[Box Plot with various distribution](example3/example3.go)  
[Box Plots with real data](example4/example4.go)    

## Exercises

### Exercise 1

**Part A** Create a box plot of the values in the third column `bmi` of [diabetes.csv](data/diabetes.csv).

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

**Part B** By examining the box plot, draw (on paper) what you think a histogram of the bmi values would look like.

### Exercise 2

Create a histogram of the values in the third column `bmi` of [diabetes.csv](data/diabetes.csv).  Compare the histogram to your drawing created in **Part A**.

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

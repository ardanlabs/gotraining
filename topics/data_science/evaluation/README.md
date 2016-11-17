## Evaluation

Being able to properly evaluate a model is essential. Without evaluation, the model development process is just guess work.  Evaluation helps us find the best model to make a prediction and gives us an idea about how the model should perform in the future. represents our data and how well the chosen model will work in the future. A whole variety of evaluation metrics have been developed and not all evaluation metrics are relevant to all models.  We will explore a sampling of these metrics/scores here, but, as data scientists, it is very important to evaluate the metrics/scores you are using on a case by case basis.

## Notes

- "When building prediction models, the primary goal should be to make a model that most accurately predicts the desired target value for new data. The measure of model error that is used should be one that achieves this goal." from [here](http://scott.fortmann-roe.com/docs/MeasuringError.html)
- Let's say we are predicting if people have or don't have a disease (from [here](http://www.dataschool.io/simple-guide-to-confusion-matrix-terminology/)):
    - true positives (TP): These are cases in which we predicted yes (they have the disease), and they do have the disease.   
    - true negatives (TN): We predicted no, and they don't have the disease.   
    - false positives (FP): We predicted yes, but they don't actually have the disease. (Also known as a "Type I error.")    
    - false negatives (FN): We predicted no, but they actually do have the disease. (Also known as a "Type II error.")

## Links

[Comparison of Evaluation Measures - Wikipedia](https://en.wikipedia.org/wiki/Precision_and_recall)    
[Accurately Measuring Model Prediction Errors](http://scott.fortmann-roe.com/docs/MeasuringError.html)       
[Understanding the Bias-Variance Tradeoff](http://scott.fortmann-roe.com/docs/BiasVariance.html)    
[Simple guide to confusion matrix terminology](http://www.dataschool.io/simple-guide-to-confusion-matrix-terminology/)

## Code Review

[Calculate R2 (Coefficient of Determination)](example1/example1.go)  
[Calculate Mean Absolute Error](example2/example2.go)   
[Caclulate Accuracy](example3/example3.go)   
[Caclulate Precision](example4/example4.go)   
[Caclulate Recall](example5/example5.go) 

## Exercises

### Exercise 1

For the `labeled.csv` results, implement and calculate the evaluation metric called "specificity" as defined [here](https://en.wikipedia.org/wiki/Confusion_matrix).  Think about when we might want to use this as compared with accuracy, precision, or recall.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

### Exercise 2

For the `continuous.csv` results, implement and calculate the evaluation metric called "mean squared error" as defined [here](https://en.wikipedia.org/wiki/Mean_squared_error).  What advantages or disadvantages might this metric have as compared to mean absolute error?

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

## Data Pipelines 

Often we need to connect various stages of data processing.  For example, with machine learning workflows, we need to train our models, we need to pre-process data, we need to utilize a model to make predictions, etc.  In order to maintain reproducibility and deploy multi-stage processing, we need to create a Data pipeline that connects multiple stages and maintains a record of the data input/output to the various stages.

[Pachyderm](http://pachyderm.io/), an open source project written in Go, allows us to create containerized data pipelines where each stage of our data pipeline is defined by a Docker container.  In addition, Pachyderm versions the data input/output of each stage such that we exactly reproduce any processing and have "provenance" for results.

## Notes

- A Pachyderm pipeline is created based on a [JSON specification](http://docs.pachyderm.io/en/latest/reference/pipeline_spec.html).
- When a pipeline is created the Docker images reference are deployed to Kubernetes.
- Pachyderm controls what data is input to the running containers (or pods) on Kubernets, such that your data is processed in the manner described by the specification.
- You can scale each stage of the pipeline by launching multiple "workers." 

## Links

[Creating Pachyderm analysis pipelines](http://docs.pachyderm.io/en/latest/fundamentals/creating_analysis_pipelines.html)    
[Distributed computing on Pachyderm](http://docs.pachyderm.io/en/latest/fundamentals/distributed_computing.html)   
[Updating Pachyderm pipelines](http://docs.pachyderm.io/en/latest/fundamentals/updating_pipelines.html)     

## Code Review

[Create a linear regression model training pipeline](example1/train.json)  
[Create a linear regression prediction pipeline](example2/predict.json)
[Update the linear regression model to a multiple regression model](example3/train.json)  

## Exercises

### Exercise 1

With a local Pachyderm cluster running, create two data repositories for a linear regression pipeline:

```
$ pachctl create-repo training
$ pachctl create-repo attributes
```

Version our training data set in Pachyderm:

```
$ cd data
$ pachctl put-file training master -c -f diabetes.csv
```

Create the training data pipeline:

```
$ cd ../example1/
$ pachctl create-pipeline -f train.json
```

Make sure the training job completes:

```
$ pachctl list-job 
```

### Exercise 2

Put example "attribute" files into the `attributes` data repository:

```
$ cd ../data/test/
$ pachctl put-file attributes master -c -r -f .
```

Create the prediction pipeline:

```
$ cd ../../example2/
$ pachctl create-pipeline -f predict.json
```

Make sure the prediction job completes:

```
$ pachctl list-job 
```

### Exercise 3

Update the linear regresion pipeline to utlize a multiple regression model referenced in [exercises/template3](exercises/template3).

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

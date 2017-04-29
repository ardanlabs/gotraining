## Containers

Our data analyses and models are all but worthless if they stay on our laptops.  We need to be able to get our analyses out into the wild of our company's infrastructure, and we need to make sure that our applications behave as expected in that environment.  Containers allow us to package up our analysis code into a portable unit that will run on any machine (or at least any machine with Docker) and behave the exact same way as it behaved on our local machine.

## Notes

- A Docker "image" is a portable set of layers that includes our application and things that it needs to run properly.
- A Docker "container" is a running instance of a Docker image.
- You build a Docker image based on a Dockerfile.
- The best practice for Go devs is to build your Go binary outside of Docker and add this binary to the Docker image (assuming you aren't relying on cgo).

## Links

[Introduction to Docker](https://training.docker.com/introduction-to-docker)    
[Building minimal Go Docker images](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)  

## Code Review

[Containerize training of a linear regression model with a single ind. var.](example1)  
[Containerize linear regression prediction](example2)  

## Exercises
 
### Exercise 1

Containerize the multiple linear regression model implemented in [template1.go](exercises/template1/template1.go). Specifically, create a Dockerfile and Makefile that builds the go binary, puts it in a Docker image, and uploads the image to Docker Hub.   

[Template](exercises/template1) |
[Answer](exercises/exercise1)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).

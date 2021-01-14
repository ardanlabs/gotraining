# The Birthday Problem

The [birthday problem](https://en.wikipedia.org/wiki/Birthday_problem) asks what is the probability that in a group of people, at least two people will have the same birthday?

We're going to run a simulation where in each round we sample `n` (23 in our case) and for each round we check if we have at least two people with the same birthday.

At the end we divide the number of times we saw duplicate birthdays with the total number of rounds to get probability.

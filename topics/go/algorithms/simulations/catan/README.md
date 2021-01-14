# Catan Tiles

In the game of [Catan](https://en.wikipedia.org/wiki/Catan), you gain resources when a roll of two dice match a number of a tile on the board. At the beginning of the game, you place your settlements next to tiles with the idea of selecting titles that have a higher probability of being matched.

For each value of two-dice roll (2-12) we'll count how many times it came up in a simulation of `n` rounds.
At the end we'll divide the count we got for each number with the total number of runs to get probabilities.

Note that Catan does not have tiles with the number 7 on them.

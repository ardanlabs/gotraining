# Exercises

## Interfaces

### Exercise 1
Declare an interface called Speaker with a method called SayHello. Declare a struct that represents a person who speaks English and one who speaks Chinese. Implement the interface for each struct using a pointer receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type Speaker and assign the address of each value and call the method.

### Exercise 2
From exercise 1, add a new function called SayHello that accepts a value of type Speaker. From that function call the SayHello method. Then change the program to pass the address of each struct type to the function.
# **Question 1: Determine if Leap Year or not**  

## **Problem Statement**  

Develop a program to determine if a given year is a leap year using control flow.

## **Requirements**  

- The program should prompt the user to enter a year.  
- It should use conditional statements (e.g., `if`, `else if`, `else`) to determine if the year is a leap year.  
- The output should provide if the year is a leap year or not.  

## **Example Execution**  

```sh
Enter a year: 2024
2024 is a leap year.
```

## **Implementation Guidelines**  

- Use the `fmt` package for input and output.  
- Read user input using `fmt.Scanln`.  
- Implement conditional logic to do the decision branch and determine if leap year or not.  

## **Expected Output**  

For the input values `2017`, the output should be:  

```shell
2017 is not a leap year.
```

## **File Structure**  

```shell
question3/
├── main.go
├── go.mod
├── logic/
    ├── question3.go
    ├── question3_test.go
└── README.md
```

## **Testing**  

- A test file `question3_test.go` should be included to validate different cases.  
- Use Go's `testing` package to write unit tests.  

## **Run the Program**  

To execute the program, navigate to the `question3` directory and run:  

```sh
go run main.go
```

## **Run Tests**  

```sh
go test
```

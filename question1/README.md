# **Question 1: Find the Largest of Three Numbers**  

## **Problem Statement**  

Write a Go program that accepts user input for three numbers and determines the largest using conditional statements.  

## **Requirements**  

- The program should prompt the user to enter three numbers.  
- It should use conditional statements (e.g., `if`, `else if`, `else`) to determine the largest number.  
- The output should clearly indicate the largest number among the three.  

## **Example Execution**  

```sh
Enter first number: 12  
Enter second number: 45  
Enter third number: 23  
The largest number is: 45
```

## **Implementation Guidelines**  

- Use the `fmt` package for input and output.  
- Read user input using `fmt.Scanln`.  
- Implement conditional logic to compare the three numbers and determine the largest.  

## **Expected Output**  

For the input values `10`, `25`, and `17`, the output should be:  

```shell
The largest number is: 25
```

## **File Structure**  

```shell
question-1/
├── main.go
├── go.mod
├── logic/
    ├── question1.go
    ├── question1_test.go
└── README.md
```

## **Testing**  

- A test file `question1_test.go` should be included to validate different cases.  
- Use Go's `testing` package to write unit tests.  

## **Run the Program**  

To execute the program, navigate to the `question-1` directory and run:  

```sh
go run main.go
```

## **Run Tests**  

```sh
go test
```

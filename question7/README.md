# **Question 7: Calculate Factorial Using Recursion**  

## **Problem Statement**  

Develop a Go program that calculates the factorial of a number using recursion.

## **Requirements**  

- The program should prompt the user to enter a number.
- It must use recursion to calculate the factorial.
- The output should display the factorial of the entered number.

## **Example Execution**  

```sh
Enter a number: 5
The factorial of 5 is 120
```

## **Implementation Guidelines**  

- Use Go's built-in packages (e.g., `fmt` for input/output).  
- Read user input using `fmt.Scanln`.  
- Implement a recursive function to compute the factorial.  

## **Expected Output**  

For example, if the input is `5`, the output should be:  

```shell
The factorial of 5 is 120
```

## **File Structure**  

```shell
question7/
├── main.go
├── go.mod
├── logic/
│   ├── question7.go
│   └── question7_test.go
└── README.md
```

## **Testing**  

- A test file `question7_test.go` is included to validate different cases.  
- Use Go's `testing` package to write unit tests.  
- To run the tests, execute:  

```sh
go test
```

## **Run the Program**  

To run the program, navigate to the `question7` directory and execute:  

```sh
go run main.go
```


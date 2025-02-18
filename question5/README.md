# **Question 5: Check if a String is a Palindrome**  

## **Problem Statement**  

Create a Go program that reads a string input and determines if it is a palindrome.

## **Requirements**  

- The program should prompt the user to enter a string.
- It should check if the string is the same forward and backward.
- The output should indicate whether the input is a palindrome or not.

## **Example Execution**  

```sh
Enter the string to check: racecar
racecar is a palindrome.

Enter the string to check: hello
hello is not a palindrome.
```

## **Implementation Guidelines**  

- Use the `fmt` package for input and output.  
- Read user input using `fmt.Scanln`.  
- Implement logic to reverse the string and compare it with the original input.  

## **Expected Output**  

For the input value `racecar`, the output should be:  

```shell
racecar is a palindrome.
```

For the input value `hello`, the output should be:  

```shell
hello is not a palindrome.
```

## **File Structure**  

```shell
question5/
├── main.go
├── go.mod
├── logic/
    ├── question5.go
    ├── question5_test.go
└── README.md
```

## **Testing**  

- A test file `question5_test.go` should be included to validate different cases.  
- Use Go's `testing` package to write unit tests.  

## **Run the Program**  

To execute the program, navigate to the `question5` directory and run:  

```sh
go run main.go
```

## **Run Tests**  

```sh
go test
```

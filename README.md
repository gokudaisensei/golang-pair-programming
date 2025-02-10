# Pair-Programming Repository

This repository contains solutions for a series of pair-programming exercises in Golang. Each exercise (or question) is implemented with a focus on writing clean, testable, and modular code. The repository is organized to separate core logic from the application entry point, making it easier to maintain and extend.

## Repository Structure

```sh
pair-programming/
├── logic
│   ├── question1.go         # Implementation for Question 1: Finding the largest of three numbers
│   └── question1_test.go    # Unit tests for the Question 1 logic
├── main.go                  # Application entry point that utilizes the logic package
└── README.md                # This file
```

- **logic/**: Contains the core logic for the exercises. Each question's implementation and its corresponding tests are organized within this folder.
- **main.go**: Serves as the entry point for the application. It imports and uses the logic from the `logic` package to execute the exercises.
- **README.md**: Provides an overview of the repository, instructions for setup, usage, testing, and contribution guidelines.

## Getting Started

### Prerequisites

- **Go**: Ensure that Go (version 1.16 or later) is installed on your system.
- **Git**: Recommended for version control and collaboration.

### Installation

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/yourusername/pair-programming.git
   cd pair-programming
   ```

2. **Initialize the Go Module (if not already initialized):**

   ```sh
   cd questionX
   go mod init questionX
   go mod tidy
   ```


## Running the Application

Each question in this repository has an implementation in the `logic/` directory and can be executed through `main.go` or other test files.

To run the application, execute:

```sh
go run main.go
```

This will prompt for user input (if applicable) and utilize the corresponding logic from the `logic/` package to process and display the output.

## Running Tests

Unit tests for the implementations are located in the `logic/` directory. To run all tests, execute:

```sh
go test ./logic -v
```

This will run tests for all implemented questions, ensuring correctness and expected behavior.

## Implementation Details

- **Logic (`logic/` directory):**  
  Each question is implemented as a separate Go file in this directory.  
  - **Functions:** Each file contains a function specific to solving a particular question.
  - **Testing:** Each question has an associated test file (e.g., `questionX_test.go`) to validate its logic.

- **Application Entry (`main.go`):**  
  The `main.go` file serves as the entry point for executing specific questions. It may handle user input and call the appropriate functions from the `logic/` package.

## Contributing

Contributions are welcome! If you would like to add new exercises or improve existing code, please follow these steps:

1. Fork the repository.
2. Create a new branch for your question or bugfix.
3. Ensure that your code adheres to the existing style and includes appropriate tests.
4. Submit a pull request for review.

## Acknowledgements

This repository was developed as part of a pair-programming exercise for a Golang class, promoting collaboration, clean coding practices, and test-driven development.

## Contact

For any questions or suggestions, please contact [gokul.unnikrishnan@mca.christuniversity.in](mailto:gokul.unnikrishnan@mca.christuniversity.in).

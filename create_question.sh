#!/bin/bash
# Usage: ./create_question.sh question<number>
# Example: ./create_question.sh question6

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 question<number> (e.g. question1)"
  exit 1
fi

QUESTION_DIR="$1"
# Extract the numeric part from "questionX" (e.g., "6" from "question6")
QUESTION_NUM=$(echo "$QUESTION_DIR" | sed 's/[^0-9]*//g')

# Create directory structure
mkdir -p "$QUESTION_DIR/logic"

##############################
# Create README.md
##############################
cat > "$QUESTION_DIR/README.md" <<EOF
# **Question ${QUESTION_NUM}: [Insert Title Here]**

## **Problem Statement**

[Insert problem statement here.]

## **Requirements**

- The program should [describe the expected behavior].
- It must use [describe the programming concepts, e.g., loops, conditionals].
- The output should [describe what output is expected].

## **Example Execution**

\`\`\`sh
# Example:
Enter input: [sample input]
Output: [sample output]
\`\`\`

## **Implementation Guidelines**

- Use Go's built-in packages (e.g., \`fmt\` for input/output).
- Read user input using functions like \`fmt.Scanln\` or \`bufio.Scanner\`.
- Implement the necessary logic to meet the problem requirements.

## **Expected Output**

For example, if the input is \`[example input]\`, the output should be:

\`\`\`shell
[Expected output here]
\`\`\`

## **File Structure**

\`\`\`shell
${QUESTION_DIR}/
├── main.go
├── go.mod
├── logic/
│   ├── ${QUESTION_DIR}.go
│   └── ${QUESTION_DIR}_test.go
└── README.md
\`\`\`

## **Testing**

- A test file \`${QUESTION_DIR}_test.go\` is included to validate different cases.
- Use Go's \`testing\` package to write unit tests.
- To run the tests, execute:

\`\`\`sh
go test
\`\`\`

## **Run the Program**

To run the program, navigate to the \`${QUESTION_DIR}\` directory and execute:

\`\`\`sh
go run main.go
\`\`\`
EOF

##############################
# Create main.go
##############################
cat > "$QUESTION_DIR/main.go" <<EOF
package main

import (
    "fmt"
    // "your_module_path/${QUESTION_DIR}/logic" // update with your module path if needed
)

func main() {
    fmt.Println("Run tests or implement your logic in the logic package.")
}
EOF

##############################
# Create go.mod
##############################
cd $QUESTION_DIR; go mod init $QUESTION_DIR; cd ..

##############################
# Create logic/questionX.go
##############################
cat > "$QUESTION_DIR/logic/${QUESTION_DIR}.go" <<EOF
package logic

// SolveProblem is a placeholder function.
// Update the function signature and implementation as needed.
func SolveProblem(input string) string {
    // TODO: implement logic here
    return ""
}
EOF

##############################
# Create logic/questionX_test.go
##############################
cat > "$QUESTION_DIR/logic/${QUESTION_DIR}_test.go" <<EOF
package logic

import "testing"

func TestSolveProblem(t *testing.T) {
    result := SolveProblem("test input")
    expected := "" // define expected output here
    if result != expected {
        t.Errorf("Expected %q, got %q", expected, result)
    }
}
EOF

echo "Created $QUESTION_DIR with basic file structure."

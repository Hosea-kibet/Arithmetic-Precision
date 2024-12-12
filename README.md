# Arbitrary Precision Calculator

This project implements an arbitrary precision calculator in Golang. The primary goal of this application is to perform basic arithmetic operations (addition, subtraction, and multiplication) on integers of arbitrary size.  The rest of the operations  will be included. I was caught by time.

## Key Features

1. **Arbitrary Precision**: Supports operations on large  numbers without using libraries.
2. **Sign Handling**: Properly manages positive and negative numbers.
3. **Efficient Arithmetic**: Uses digit-by-digit processing for addition, subtraction, and multiplication.
4. **Interactive CLI**: Users can input expressions in a REPL and they are able to get immediate results.

## How It Works

- **Data Structure**: Numbers are stored as slices  (`[]int`) to handle arbitrary lengths.
- **Sign Management**: A `sign` field in the `BigInt` struct keeps track of the number's   whether posive or negative.
- **Operations**:
  - Addition and subtraction consider both the signs and magnitudes of the operands.
  - Multiplication is performed digit-by-digit, and carries are handled.

## Installation

To run this program, ensure you have  GO installed on your machine.

Clone the repository and navigate to the project directory:

Install Dependencies running go mod init

You will exec on the terminal  and you can person either of the operations stated above

## Usage

The calculator supports the following operations:

- Addition (`+`):

  ```
  calc> 12345678901234567890 + 98765432109876543210
  111111111011111111100
  ```

- Subtraction (`-`):

  ```
  calc> 98765432109876543210 - 12345678901234567890
  86419753208641975320
  ```

- Multiplication (`*`):

  ```
  calc> 123456789 * 987654321
  121932631112635269
  ```

To exit the program, type `exit`:

```bash
calc> exit
```

## Interesting Aspects of the Solution

1. **Digit-by-Digit Arithmetic**: The core logic for addition, subtraction, and multiplication processes numbers one digit at a time from the least significant to the most significant.
2. **Sign Management**: The program consistently handles mixed-sign arithmetic, ensuring correctness.
3. **Leading Zero Handling**: Results are normalized to remove leading zeros, maintaining readability.
4. **Scalability**: The slice-based design allows the program to handle extremely large integers.

## Code Structure

- **BigInt Struct**: Represents large integers with a sign and digits.
- **Addition and Subtraction**: Implemented to handle both matching and differing signs.
- **Multiplication**: Processes every pair of digits from the two operands and accumulates the result.
- **CLI**: The `main` function provides an interactive shell for user input.

## Future Improvements

- Add division and modulus operations.
- Implement better error handling for invalid inputs.
- Optimize multiplication for performance using algorithms like Karatsuba.
- Support floating-point numbers.

## Contribution

Feel free to fork this repository and submit pull requests for improvements. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.


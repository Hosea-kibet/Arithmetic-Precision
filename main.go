package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BigInt struct {
	sign   int
	digits []int
}

func NewBigInt(value string) BigInt {
	bigInt := BigInt{sign: 1}
	if strings.HasPrefix(value, "-") {
		bigInt.sign = -1
		value = value[1:]
	}
	for _, r := range value {
		bigInt.digits = append(bigInt.digits, int(r-'0'))
	}
	return bigInt
}

func (b BigInt) String() string {
	if b.sign == -1 {
		return "-" + digitsToString(b.digits)
	}
	return digitsToString(b.digits)
}

func digitsToString(digits []int) string {
	var result strings.Builder
	for _, d := range digits {
		result.WriteRune(rune(d + '0'))
	}
	return result.String()
}

// AddBigInts performs addition of two BigInts
func AddBigInts(a, b BigInt) BigInt {
	if a.sign == b.sign {
		return BigInt{
			sign:   a.sign,
			digits: addDigits(a.digits, b.digits),
		}
	}
	if compareDigits(a.digits, b.digits) >= 0 {
		return BigInt{
			sign:   a.sign,
			digits: subtractDigits(a.digits, b.digits),
		}
	}
	return BigInt{
		sign:   b.sign,
		digits: subtractDigits(b.digits, a.digits),
	}
}

func addDigits(a, b []int) []int {
	fmt.Println("The length of the a ", len(a))
	result := []int{}
	carry := 0
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b) // We need to loop through the larger length
	}

	for k := 0; k < maxLen; k++ {
		sum := carry
		if k < len(a) {
			sum += a[len(a)-1-k] // Add digit from 'a' (start from the least significant)
		}
		if k < len(b) {
			sum += b[len(b)-1-k] // Add digit from 'b' (start from the least significant)
		}

		result = append([]int{sum % 10}, result...) // Add the digit (sum % 10)
		carry = sum / 10                            // Update carry (sum / 10)
	}

	return result
}

func compareDigits(a, b []int) int {
	if len(a) != len(b) {
		return len(a) - len(b)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return 0
}

// SubtractBigInts performs subtraction of two BigInts
func SubtractBigInts(a, b BigInt) BigInt {
	if a.sign != b.sign {
		return BigInt{
			sign:   a.sign,
			digits: addDigits(a.digits, b.digits),
		}
	}

	if compareDigits(a.digits, b.digits) >= 0 {
		return BigInt{
			sign:   a.sign,
			digits: subtractDigits(a.digits, b.digits),
		}
	}
	return BigInt{
		sign:   -a.sign,
		digits: subtractDigits(b.digits, a.digits),
	}
}

func subtractDigits(a, b []int) []int {
	result := []int{}
	borrow := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 {
		diff := a[i] - borrow
		if j >= 0 {
			diff -= b[j]
			j--
		}
		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}
		result = append([]int{diff}, result...)
		i--
	}
	for len(result) > 1 && result[0] == 0 {
		result = result[1:]
	}
	return result
}

// MultiplyBigInts performs multiplication of two BigInts
func MultiplyBigInts(a, b BigInt) BigInt {
	result := make([]int, len(a.digits)+len(b.digits))
	for i := len(a.digits) - 1; i >= 0; i-- {
		for j := len(b.digits) - 1; j >= 0; j-- {
			mul := a.digits[i]*b.digits[j] + result[i+j+1]
			result[i+j+1] = mul % 10
			result[i+j] += mul / 10
		}
	}
	for len(result) > 1 && result[0] == 0 {
		result = result[1:]
	}
	return BigInt{
		sign:   a.sign * b.sign,
		digits: result,
	}
}

func main() {
	fmt.Println("Arbitrary Precision Calculator (Golang)")
	fmt.Println("Supported operations: + (add), - (subtract), * (multiply)")
	fmt.Println("Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("calc> ")
		if !scanner.Scan() {
			break
		}
		expr := scanner.Text()

		if expr == "exit" {
			break
		}

		parts := strings.Fields(expr)
		if len(parts) != 2 && len(parts) != 3 {
			fmt.Println("Invalid input. Usage: <num1> <op> <num2>")
			continue
		}

		num1 := NewBigInt(parts[0])
		op := parts[1]
		var num2 BigInt
		if len(parts) == 3 {
			num2 = NewBigInt(parts[2])
		}

		switch op {
		case "+":
			fmt.Println(AddBigInts(num1, num2))
		case "-":
			fmt.Println(SubtractBigInts(num1, num2))
		case "*":
			fmt.Println(MultiplyBigInts(num1, num2))
		default:
			fmt.Println("Unsupported operation")
		}
	}
}

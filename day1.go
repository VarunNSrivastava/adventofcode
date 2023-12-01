package main

import (
	"bufio"
	"fmt"
	"os"
)
 const FILENAME = "day1.txt"

 func safeSlice(s string, start, length int) string {
	runes := []rune(s)
    if start < 0 || start + length > len(runes) || length <= 0 {
        return "" 
    }
    return string(runes[start:start + length])
}

func isStringDigit(s string) (int, bool) { 
	digitMap := map[string]int{
        "zero":  0,
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    }
	digit, ok := digitMap[s]
	return digit, ok
}

func isRuneDigit(r rune) (int, bool) {
	if r >= '1' && r <= '9' {
		digit := int(r - '0')
		return digit, true
	}

	return 0, false
}


func main() {
	var filename = FILENAME
	if len(os.Args) == 2 {
        filename = os.Args[1]
    }

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // ensure files closed when function returns

	scanner := bufio.NewScanner(file) // declare variable with type inference
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		var first, last int
		// runes are alias for int32
		// you can print as rune or int32 via %c or %d, respectively
		for i, r := range line {
			if digit, found := isRuneDigit(r); found {
				last = digit
			} else if digit, found = isStringDigit(safeSlice(line, i, 3)); found {
				last = digit
			} else if digit, found = isStringDigit(safeSlice(line, i, 4)); found {
				last = digit
			} else if digit, found = isStringDigit(safeSlice(line, i, 5)); found {
				last = digit
			}
			if first == 0 {
				first = last
			}
		}
		total += last + 10*first
	}

	// Erros that occurred during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Printf("Total is %d\n", total)
}

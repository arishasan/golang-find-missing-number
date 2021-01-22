package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// Declare our main variables
	var s string
	// End of declaration

	fmt.Println("----------------------------------------------")

	// Read and do validation for inputted number from user
	fmt.Println("Enter list of non-zero positive numbers : ")
	
	for {

		fmt.Scan(&s) // Get the input from user

		sLength := len(s) // Calculate the length of it's inputted numbers

		if(sLength < 1 || s == ""){ // First validation

			fmt.Println("Please input valid numbers")

		}else{

			if(sLength < 3 || sLength > 1000){ // Logic will work if this criteria fulfiled

				fmt.Println("Length minimal of inputted number is 3 words, and Maximum is 1000 words")

			}else{

				holder := findMissingNumber(s)
				fmt.Println("----------------------------------------------")

				if holder == -1 {
					fmt.Println("We haven't found any missing number from your inputted numbers or there's error occured. Keep in mind, we programmed only to find a missing number.")
				}else{
					fmt.Println("We found it, the missing number is : " + strconv.Itoa(holder))
				}

				break

			}

		}

	}

	// End of input and validation

}

// For unit testing only

func testFindMissingNumber(s string) string {

	holder := findMissingNumber(s)
	// fmt.Println(holder)
	return strconv.Itoa(holder)

}

// End of unit testing

// Function for getting the digit/integer at position i with the length of main
func getDigit(s string, i int, main int) int {

	if i + main > len(s) {
		return -1
	}

	// Find value at index i with the length of main
	retVal := 0
	for l := 0; l < main; l++ {
		temp := bytesToInt(s[i + l]) - '0'
		// fmt.Println(bytesToInt(s[i+1]))
		if temp < 0 || temp > 9 {
			return -1
		}
		retVal = retVal * 10 + temp
	}

	return retVal

}
// End of function

// Logical function, search the missing number
func findMissingNumber(s string) int {

	maxDigits := 6

	// Iterate maxDigits as main
	for main := 1; main <= maxDigits; main++{

		// fmt.Println(main)

		// Get the value of first number with current length
		n := getDigit(s, 0, main)
		if n == -1 { // If n returns -1, break the loop. Consider it an error
			break
		}

		// fmt.Println(n)

		tempMissingNumber := -1 // Store the missing number

		isError := false // To indicate whether the iteration failed or no

		// Iterate the subsequent numbers after the first digit/previous number, we indicate it as variable n
		for i := main; i != len(s); i += 1 + int(math.Log10(float64(n))) {

			// fmt.Println(s[i])

			/*
			If we haven't yet found the missing number for the current iteration. Next number will be n + 2
			*/

			if tempMissingNumber == -1 && getDigit(s, i, 1 + int(math.Log10(float64(n+2)))) == n + 2 { // If returned data from getDigit equal or the same as n + 2, it means we finally found the missing number
				tempMissingNumber = n + 1
				n += 2
			} else if getDigit(s, i, 1 + int(math.Log10(float64(n+1)))) == n+1 { // If next value is equal or the same as n + 1 then do increment to n and do another cycle of iteration
				n++
			} else { // If there's nothing to iterate, break the loop and indicate it as an error
				isError = true
				break
			}

		}

		if isError == false {
			return tempMissingNumber
		}

	}

	return -1 // Not found (error) or there's no missing number

}
// End of function

// Function for converting type data
func bytesToString(data byte) string {
	return string(byte(data))
}

func bytesToInt(data byte) int {
	return int(byte(data))
}
// End of function
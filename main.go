package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	reload "reload/functions"
)

func main() {
	if (len(os.Args)) != 3 {
		fmt.Println("Error wrong arguments")
		log.Fatal()
	}
	arg := os.Args[1]  // sample text
	arg2 := os.Args[2] // result text

	// create result file
	file, err := os.Create(arg2)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Clear the file if it exist
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error clearing file:", err)
		return
	}

	content, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println(err)
		return
	}
	// free := strings.Fields(string(content))
	// freestr := strings.Join(free, " ")
	// contentR := reload.FixPunctuationSpacing(freestr) // to fix caping prev word remove this commmit change content to contentR below.

	arr := strings.Fields(string(content))
	arr = append(arr, "")

	counter := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] == "'" && counter == 0 {
			counter++
			arr[i+1] = "'" + arr[i+1]
			arr = reload.Rightshift(arr, i)
		}
		if arr[i] == "'" && counter == 1 && i != 0 {
			counter--
			arr[i-1] = arr[i-1] + "'"
			arr = reload.Rightshift(arr, i)
		}

		if arr[i] == "(hex)" && i != 0 {
			arr[i-1] = reload.ConvToDecimal(arr[i-1], 16)
			arr = reload.Rightshift(arr, i)
		}

		if arr[i] == "(bin)" && i != 0 {
			arr[i-1] = reload.ConvToDecimal(arr[i-1], 2)
			arr = reload.Rightshift(arr, i)
		}

		if arr[i] == "(up)" && i != 0 {
			if !reload.IsAlpha(arr[i-1]) {
				for j := i - 1; j > 0; j-- {
					if reload.IsAlpha(arr[j]) {
						arr[i-1] = strings.ToUpper(arr[i-1])
						arr = reload.Rightshift(arr, i)
					}
				}
			}
		}

		if arr[i] == "(low)" && i != 0 || strings.Contains(arr[i], "(low)") {
			if !reload.IsAlpha(arr[i-1]) {
				for j := i - 1; j > 0; j-- {
					if reload.IsAlpha(arr[j]) {
						arr[i-1] = strings.ToLower(arr[i-1])
						arr = reload.Rightshift(arr, i)
					}
				}
			}
		}

		if arr[i] == "(cap)" && i != 0 {
			if !reload.IsAlpha(arr[i-1]) {
				for j := i - 1; j > 0; j-- {
					if reload.IsAlpha(arr[j]) {
						arr[i-1] = strings.ToTitle(strings.ToLower(arr[i-1])) // strings.ToUpper(string(word[0])) + word[1:]
						arr = reload.Rightshift(arr, i)
					}
				}
			}
		}

		if arr[i] == "(up," && i != 0 {
			reload.ProcessPrev(arr, i, "up")
			arr = reload.Rightshift(arr, i)
			arr = reload.Rightshift(arr, i)
		}

		if arr[i] == "(low," && i != 0 {
			reload.ProcessPrev(arr, i, "low")
			arr = reload.Rightshift(arr, i)
			arr = reload.Rightshift(arr, i)
		}

		if arr[i] == "(cap," && i != 0 {
			reload.ProcessPrev(arr, i, "cap")
			arr = reload.Rightshift(arr, i)
			arr = reload.Rightshift(arr, i)
		}

		if arr[i] == "a" || arr[i] == "A" {
			vowels := "aeiouAEIOU"
			word := arr[i+1]
			if len(word) > 0 && strings.ContainsAny(string(word[0]), vowels) {
				arr[i] += "n"
			}
		}

	}

	// fix punctuations
	strr := strings.Join(arr, " ") // conv to string
	result1 := reload.FixPunctuationSpacing(strr)

	result := strings.Fields(string(result1)) // conv back to array
	result = append(result, "")
	counter = 1
	// loop again to make sure everything is correct
	for i := 0; i < len(result); i++ {
		worder := result[i]
		if result[i] == "'" && counter == 0 {
			counter++
			result[i+1] = "'" + result[i+1]
			result = reload.Rightshift(result, i)
		}

		if result[i] == "'" && counter == 1 && i != 0 {
			counter--
			result[i-1] = result[i-1] + "'"
			result = reload.Rightshift(result, i)
		}

		if result[i] == "(hex)" && i != 0 || (strings.EqualFold(result[i], "(hex)") && i != 0) {
			result[i-1] = reload.ConvToDecimal(result[i-1], 16) + string(worder[5:])
			result = reload.Rightshift(result, i)
		}

		if result[i] == "(bin)" && i != 0 || (strings.EqualFold(result[i], "(bin)") && i != 0) {
			result[i-1] = reload.ConvToDecimal(result[i-1], 2) + string(worder[5:])
			result = reload.Rightshift(result, i)
		}

		if result[i] == "(up)" && i != 0 || (strings.EqualFold(result[i], "(up)") && i != 0) {
			if !reload.IsAlpha(arr[i-1]) {
				for j := i - 1; j > 0; j-- {
					if reload.IsAlpha(arr[j]) {
						break
					}
				}
			}

			result[i-1] = strings.ToUpper(result[i-1]) + string(worder[4:])
			result = reload.Rightshift(result, i)
		}

		if result[i] == "(low)" && i != 0 || (strings.EqualFold(result[i], "(low)") && i != 0) {
			result[i-1] = strings.ToLower(result[i-1]) + string(worder[5:])
			result = reload.Rightshift(result, i)
		}

		if result[i] == "(cap)" && i != 0 || (strings.EqualFold(result[i], "(cap)") && i != 0) {
			// result[i-1] = strings.ToTitle(strings.ToLower(result[i-1])) // strings.ToUpper(string(word[0])) + word[1:]

			word := result[i-1]
			if word[0] == '\'' {
				result[i-1] = "'" + strings.ToUpper(string(word[1])) + strings.ToLower(string(word[2:])) + string(worder[5:])
				result = reload.Rightshift(result, i)
			} else {
				result[i-1] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:])) + string(worder[5:])
				result = reload.Rightshift(result, i)
			}
		}

		if result[i] == "a" {
			vowels := "aeiouAEIOU"
			word := result[i+1]
			if len(word) > 0 && strings.ContainsAny(string(word[0]), vowels) {
				result[i] = "an"
			}
		}

	}

	res := strings.Join(result, " ")
	res = res[:len(res)-1]
	fmt.Fprintln(file, res)                  // print in file
	fmt.Printf("File contents: \n%s\n", res) // print in terminal
	defer file.Close()
}

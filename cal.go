package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return characterMap[s[0]]
	}

	sum := characterMap[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if characterMap[s[i]] < characterMap[s[i+1]] {
			sum -= characterMap[s[i]]
		} else {
			sum += characterMap[s[i]]
		}
	}

	return sum
}

func main() {

	fmt.Println("Посчитать римскими цифрами - введите 1")
	fmt.Println("Посчитать арабскими цифрами - введите 0")
	var isRoman bool
	fmt.Scanln(&isRoman)

	fmt.Print("Enter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)

	parts := strings.Split(input, " ")
	var1 := parts[0]
	op := parts[1]
	var2 := parts[2]

	var var1_int int
	var var2_int int
	var err1 error
	var err2 error
	if isRoman == true {
		var1_int = romanToInt(var1)
		var2_int = romanToInt(var2)
	} else {
		var1_int, err1 = strconv.Atoi(var1)
		var2_int, err2 = strconv.Atoi(var2)

		if err1 != nil || err2 != nil {
			fmt.Println("Ошибка")
		}
	}

	if op == "-" {
		result := var1_int - var2_int
		fmt.Println("Итог:", result)
	} else if op == "+" {
		result := var1_int + var2_int
		fmt.Println("Итог:", result)
	} else if op == "*" {
		result := var1_int * var2_int
		fmt.Println("Итог:", result)
	} else if op == "/" {
		result := var1_int / var2_int
		fmt.Println("Итог:", result)
	} else {
		fmt.Println("Ошибка")
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(faktorisasi(4))
	pemisah(2468)
	fmt.Println(valid("{{}}"))
}

func faktorisasi(val int) int {
	result := 1
	for i := 1; i <= val; i++ {
		result = result * i
	}
	return result
}

func pemisah(num int) {
	nilai := strconv.Itoa(num)
	s := strings.Split(nilai, "")
	result := []string{s[0]}
	for i := 1; i < len(s); i++ {
		angka1, _ := strconv.Atoi(s[i-1])
		angka2, _ := strconv.Atoi(s[i])
		if angka1%2 == 0 && angka2%2 == 0 {
			result = append(result, "-")
			result = append(result, s[i])
		} else {
			result = append(result, s[i])
		}
	}
	str := strings.Join(result, "")
	fmt.Println(str)
}

func valid(val string) bool {
	var result bool
	if len(val)%2 != 0 {
		result = false
	} else {
		s := strings.Split(val, "")
		pembagi := len(s) / 2
		for i := 0; i < pembagi; i++ {
			if s[i] == "{" && s[(len(s)-1)-i] == "}" {
				result = true
			} else if s[i] == "(" && s[(len(s)-1)-i] == ")" {
				result = true
			} else if s[i] == "[" && s[(len(s)-1)-i] == "]" {
				result = true
			} else {
				result = false
			}
		}
	}
	return result
}

package main

import "strings"

func postProcess(word string) string {
	charArr := strings.Split(word, "")

	if len(word) > 1 {
		if string([]rune(word)[1]) == "ং" {
			charArr[0] = "ঙ"
		}

		if strings.HasPrefix(word, "ৰ্") {
			charArr[1] = ""
		}

		if strings.HasSuffix(word, "্") {
			charArr[len(charArr)-1] = " ।"
		}
	}

	word = strings.Join(charArr, "")

	charArr = strings.Split(word, "O")
	word = strings.Join(charArr, "ো")
	word = strings.ReplaceAll(word, "o", "")

	return word
}

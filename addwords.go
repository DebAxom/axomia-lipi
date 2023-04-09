package main

import "strings"

func addWords(first, last string) string {
	charArr_First := strings.Split(first, "")
	charArr_Last := strings.Split(last, "")
	last_char := charArr_First[len(charArr_First)-1]
	first_char := charArr_Last[0]

	if !isVowel(last_char) && first_char == "ৰ" && last_char != "ং" {
		charArr_Last[0] = "্ৰ"
	}

	if last_char == "ৰ" && !isVowel(first_char) {
		if first_char == "য়" {
			charArr_Last[0] = "য"
		}
		charArr_First[len(charArr_First)-1] = "ৰ্"
		return strings.Join(charArr_First, "") + strings.Join(charArr_Last, "")
	}

	// Converting য় to ্য in some situations.
	if first_char == "য়" && !isVowel(last_char) {
		charArr_Last[0] = "্য"
	}

	// Converting ং to ঙ in some situations.
	if last_char == "ং" && isVowel(first_char) {
		charArr_First[len(charArr_First)-1] = "ঙ"
	}

	// Adding kars of vowels.
	if isVowel(last_char) {

		if first_char == "আ" && last_char == "ই" {
			charArr_Last[0] = "য়া"
		}

		if first_char == "আ" && last_char == "উ" {
			charArr_Last[0] = "ৱা"
		}

		return strings.Join(charArr_First, "") + strings.Join(charArr_Last, "")
	}

	if isVowel(first_char) {
		if last_char == "গ" && first_char == "ও" || last_char == "স" && first_char == "ও" {
			charArr_Last[0] = "ো"
		} else {
			charArr_Last[0] = karMap[first_char]
		}

		return strings.Join(charArr_First, "") + strings.Join(charArr_Last, "")
	}

	return strings.Join(charArr_First, "") + strings.Join(charArr_Last, "")
}

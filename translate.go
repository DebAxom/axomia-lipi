package main

import (
	"math"
	"strings"
)

var vowels = [20]string{"অ", "আ", "ই", "ঈ", "উ", "ঊ", "এ", "ঐ", "ও", "ঔ", "া", "ি", "ী", "ু", "ূ", "ে", "ৈ", "ো", "ৌ", "o"}

func TranslateWord(word string) string {
	word = strings.ToLower(word)
	if val, WordExists := wordsMap[word]; WordExists {
		return val
	}

	word = Translate(word)
	word = postProcess(word)

	return word
}

func JoinChars(word string) string {
	word = strings.ReplaceAll(word, "oi", "ঐ")
	word = strings.ReplaceAll(word, "ou", "ঔ")
	word = strings.ReplaceAll(word, "ee", "i")
	word = strings.ReplaceAll(word, "th", "থ")
	word = strings.ReplaceAll(word, "bh", "ভ")
	word = strings.ReplaceAll(word, "cl", "ক্ল")
	word = strings.ReplaceAll(word, "kl", "ক্ল")
	word = strings.ReplaceAll(word, "jh", "ঝ")
	word = strings.ReplaceAll(word, "zh", "ঝ")
	word = strings.ReplaceAll(word, "sk", "স্ক")
	word = strings.ReplaceAll(word, "shn", "ষ্ণ")
	word = strings.ReplaceAll(word, "sh", "ষ")
	word = strings.ReplaceAll(word, "ch", "চ")
	word = strings.ReplaceAll(word, "khy", "ক্ষ")
	word = strings.ReplaceAll(word, "kh", "খ")
	word = strings.ReplaceAll(word, "gh", "ঘ")
	word = strings.ReplaceAll(word, "ndh", "n.dh")
	word = strings.ReplaceAll(word, "dh", "ধ")
	word = strings.ReplaceAll(word, "ph", "ফ")
	word = strings.ReplaceAll(word, "ng", "ং")
	word = strings.ReplaceAll(word, "nt", "n.t")
	word = strings.ReplaceAll(word, "tn", "t.n")
	word = strings.ReplaceAll(word, "sw", "s.w")
	word = strings.ReplaceAll(word, "hm", "h.m")
	word = strings.ReplaceAll(word, "mh", "m.h")
	word = strings.ReplaceAll(word, "jyo", "জ্যো")
	word = strings.ReplaceAll(word, "t,", "ট")
	word = strings.ReplaceAll(word, "n,", "ণ")
	word = strings.ReplaceAll(word, "o,", "O")
	// ো  ে

	return word
}

func Translate(word string) string {

	if val, WordExist := charMap[word]; WordExist {
		return val
	}

	word = JoinChars(word)

	charArr := strings.Split(word, "")
	strLen := len(charArr)

	if strLen == 1 {
		return word
	}

	half := int(math.Ceil(float64(strLen / 2)))
	firstHalf := charArr[0:half]
	secondHalf := charArr[half:strLen]

	firstHalfStr := strings.Join(firstHalf, "")
	secondHalfStr := strings.Join(secondHalf, "")

	return addWords(Translate(firstHalfStr), Translate(secondHalfStr))
}

func isVowel(str string) bool {
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == str {
			return true
		}
	}
	return false
}

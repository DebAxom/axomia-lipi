package main

import (
	"encoding/json"
	"os"
)

func returnMapfromJSON(jsonStr string) map[string]string {
	x := map[string]string{}
	json.Unmarshal([]byte(jsonStr), &x)
	return x
}

var charRawData, _ = os.ReadFile("./words/char.json")
var karRawData, _ = os.ReadFile("./words/kar.json")
var wordsRawData, _ = os.ReadFile("./words/words.json")

var charMap = returnMapfromJSON(string(charRawData))
var karMap = returnMapfromJSON(string(karRawData))
var wordsMap = returnMapfromJSON(string(wordsRawData))

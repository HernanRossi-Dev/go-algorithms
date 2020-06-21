package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	result := make([][]string, 0)
	if length < 2 {
		result = append(result, strs)
		return result
	}
	// HashMap has a map of interter mapping to list of list of strings for anagrams that match the hash for chaining
	codeHash := make(map[int][][]string)
	for i := 0; i < length; i++ {
		word := strs[i]
		charsCode := 0
		wordRune := []rune(word)
		for j := 0; j < len(word); j++ {
			charsCode = charsCode + int(wordRune[j])
		}
		if anagramLists, isPresent := codeHash[charsCode]; isPresent { // Check if the hash key already exists in the hashmap
			match := false
			for k := 0; k < len(anagramLists); k++ { // If exists in the hash map iterate throu
				if checkAnagram(word, anagramLists[k][0]) { // Do a hashtable check for anagram
					anagramLists[k] = append(anagramLists[k], word)
					codeHash[charsCode] = anagramLists
					match = true
					break
				}
			}
			if !match { //If no match found in list of anagrams add a new list
				newList := []string{word}
				anagramLists = append(anagramLists, newList)
				codeHash[charsCode] = anagramLists
			}
		} else { //There is no conflicting hash entry create a new key value
			newList := [][]string{{word}}
			codeHash[charsCode] = newList
		}
	}
	for _, anagramLists := range codeHash { //Merge all lists return result
		result = append(result, anagramLists...)
	}
	return result
}

func checkAnagram(str1 string, str2 string) bool {
	wordRune1 := []rune(str1)
	wordRune2 := []rune(str2)
	hashMap := make(map[rune]int)
	for i := 0; i < len(str1); i++ {
		char := wordRune1[i]
		if count, isPresent := hashMap[char]; isPresent {
			hashMap[char] = count + 1
		} else {
			hashMap[char] = 1
		}
	}
	for j := 0; j < len(str2); j++ {
		char := wordRune2[j]
		if count, isPresent := hashMap[char]; isPresent {
			count = count - 1
			if count == 0 {
				delete(hashMap, char)
			} else {
				hashMap[char] = count
			}
		}
	}
	if len(hashMap) > 0 {
		return false
	}
	return true
}

type testAnagram struct {
	input    []string
	expected [][]string
}

func main() {
	tests := []testAnagram{}
	testOne := testAnagram{
		input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		expected: [][]string{{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"}},
	}
	tests = append(tests, testOne)

	for _, test := range tests {
		result := groupAnagrams(test.input)
		fmt.Print("Test Result: ", result, " Expected: ", test.expected)
	}
}

// {
//   input: ["tho", "tin", "erg", "end", "pug", "ton", "alb", "mes", "job", "ads", "soy", "toe", "tap", "sen", "ape", "led", "rig", "rig", "con", "wac", "gog",
//     "zen", "hay", "lie", "pay", "kid", "oaf", "arc", "hay", "vet", "sat", "gap", "hop", "ben", "gem", "dem", "pie", "eco", "cub", "coy", "pep", "wot", "wee"],
//   expected: [
//     ["wee"], ["wot"], ["pep"], ["eco"], ["pie"], ["ben"], ["coy"], ["dem"], ["sat"],
//     ["arc"], ["vet"], ["pay"], ["cub"], ["ads"], ["ton"], ["kid"], ["pug"], ["oaf"],
//     ["ape"], ["led"], ["hay", "hay"], ["tho"], ["tin"], ["mes"], ["gap"], ["erg"],
//     ["lie"], ["alb"], ["tap"], ["end"], ["toe"], ["soy"], ["gem"], ["rig", "rig"],
//     ["hop"], ["sen"], ["con"], ["job"], ["wac"], ["gog"], ["zen"]]
// }

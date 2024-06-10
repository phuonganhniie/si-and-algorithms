package leetcode_648

import (
	"strings"

	"github.com/phuonganhniie/leetcode/helper"
)

/*
[Medium] 648. Replace Words
https://leetcode.com/problems/replace-words/description/
Created: 2024-06-07
Done   : 20 mins
Attempt: 4
---------------------NOTE---------------------
Time: O(n)
Space: O(n)
*/

/*
	----------- Hash Map -----------

Approach:
1. Create a map to store all the roots from dictionary.
2. Split sentence into words.
3. Loop to each word and found it in root map. If it has, replace root to this word and append this root to result array.
4. Continue til loop all of words in sentence.
*/
func replaceWordsMap(dictionary []string, sentence string) string {
	if len(dictionary) == 0 {
		return ""
	}

	rootMap := make(map[string]bool)
	for _, root := range dictionary {
		rootMap[root] = true
	}

	words := strings.Split(sentence, " ")

	var result []string
	for _, word := range words {
		replaceWord := word
		for i := 1; i <= len(word); i++ {
			wordChar := word[:i]
			if rootMap[wordChar] {
				replaceWord = wordChar
				break
			}
		}
		result = append(result, replaceWord)
	}
	return strings.Join(result, " ")
}

/*
	----------- Trie -----------

Approach:
*/
func replaceWordsTrie(dictionary []string, sentence string) string {
	if len(dictionary) == 0 {
		return ""
	}

	trie := helper.NewTrie()
	for _, root := range dictionary {
		trie.Insert(root)
	}
	trie.Visualize()

	words := strings.Split(sentence, " ")
	for i, word := range words {
		shortestWord := trie.Search(word)
		words[i] = shortestWord
	}
	return strings.Join(words, " ")
}

func replaceWords(dictionary []string, sentence string) string {
	wordsArray := strings.Split(sentence, " ")
	for i, word := range wordsArray {
		replacedWord := word
		for _, root := range dictionary {
			if strings.HasPrefix(word, root) && len(replacedWord) > len(root) {
				replacedWord = root
				wordsArray[i] = replacedWord
			}
		}
	}
	return strings.Join(wordsArray, " ")
}

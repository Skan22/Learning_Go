package main

import (
	"strings"
	
)

func countDistinctWords(messages []string) int {
	word_count_map := make(map[string]struct{})
	
	for _,sentence:=range messages{
		words:=strings.Fields(strings.ToLower(sentence))
		for _,word:=range words{
			if _,ok := word_count_map[word];!ok{
				word_count_map[word]=struct{}{}
			}else{
				continue
			}
		}



	}
	return len(word_count_map)
	
}

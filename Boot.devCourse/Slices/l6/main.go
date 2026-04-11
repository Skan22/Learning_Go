package main

import "unicode/utf8"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64,len(messages))
	for i:=0;i<len(messages);i++{
		costs[i] = 0.01*float64(utf8.RuneCountInString(messages[i])) 
	}
	return costs
}	

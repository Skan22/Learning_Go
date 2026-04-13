package main



func getNameCounts(names []string) map[rune]map[string]int {
	mapp := make(map[rune]map[string]int)
	
	for _,name:=range names {
		runes := []rune(name)
		if _,ok := mapp[runes[0]];!ok {
			mapp[runes[0]] = make(map[string]int) 
		}
		mapp[runes[0]][name] ++ 
		
	
}
return  mapp
}

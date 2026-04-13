package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	mapp := make(map[string]user)
	if len(names)!=len(phoneNumbers){
		return  nil,errors.New("invalid sizes")
	}else{
		for i,name:=range names {
			mapp[name] = user{name,phoneNumbers[i]}

		}
		return mapp,nil
	}
}

type user struct {	
	name        string
	phoneNumber int
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func (e email) cost() int {
	if !e.isSubscribed {
		return  5*utf8.RuneCountInString(e.body)
	}else{
		return 2*utf8.RuneCountInString(e.body)
	}
}

func (e email) format() string {
	msg := ""
	if e.isSubscribed {
		msg = "Subscribed"
	}else{
		msg = "Not Subscribed"
	}

	message := fmt.Sprintf("'%s' | %s",e.body,msg )
	return  message
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

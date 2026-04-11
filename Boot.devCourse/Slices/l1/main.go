package main

import "unicode/utf8"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
		messages := [3]string{primary,secondary,tertiary}
		costs := [3]int{utf8.RuneCountInString(primary),utf8.RuneCountInString(primary)+utf8.RuneCountInString(secondary),utf8.RuneCountInString(primary)+utf8.RuneCountInString(secondary)+utf8.RuneCountInString(tertiary)}

		return  messages,costs
	}

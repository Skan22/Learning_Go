package main

import (
	"errors"
	"unicode/utf8"
)

func validateStatus(status string) error {
	if len(status) == 0{
		return errors.New("status cannot be empty")
	}else if utf8.RuneCountInString(status) > 140 {
		return  errors.New("status exceeds 140 characters")

	}else {
		return nil
	}
	
}

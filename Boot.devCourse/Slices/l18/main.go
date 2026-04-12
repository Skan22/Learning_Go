package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password)>12 {
		return false
	}
	contains_digit :=false
	contains_uppercase :=false
	for _,j:=range password{
		contains_digit = ('0'<=j && j<= '9') || contains_digit
		contains_uppercase =('A' <= j && j<='Z')  || contains_uppercase
		if contains_digit && contains_uppercase{
			return true
		}
	}
	return false
}

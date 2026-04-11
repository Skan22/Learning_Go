package main

import "fmt"

func fizzbuzz() {
	for i := 0 ;  i <= 100 ;i++{
		text :=""
		switch  {
		case i%15==0:
			text += "fizzbuzz"
		case i%3==0:
			text +="fizz"
		case i%5==0:
			text +="buzz"
		default:
			text +=fmt.Sprint(i)	
		}
		fmt.Println(text)
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}

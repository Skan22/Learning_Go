package main

import "fmt"



func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string,len(emails))

		for idx,i:=range emails {
			fmt.Println(idx)
			ch <- i 

			}
return  ch

}
package main

import (
	
	"fmt"
	"net"

	"sync"
)

func worker(ports chan int , wg *sync.WaitGroup){
	for p:=range ports {
		adress := fmt.Sprintf("supcom.tn:%d",p)
		conn,err := net.Dial("tcp",adress)
		if err !=nil{
			continue
		}else{
			fmt.Println("Connected to: ",adress)
			conn.Close()
		}
		
		wg.Done()
	}
}
func main()  {
	ports := make(chan int,65535)
	var wg sync.WaitGroup
	for i:=0;i<cap(ports);i++{
		go worker(ports,&wg)

	}
	for i:=1;i<=65535;i++{
		wg.Add(1)
		ports <- i 
	}
	wg.Wait()
	close(ports)

}
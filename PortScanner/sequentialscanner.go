package main

import (
	"fmt"
	"net"
	"strconv"
)

func notmain()  {
	for i:=0;i<1024;i++{	
		address :="scanme.nmap.org:"+strconv.Itoa(i)
			
		conn,err := net.Dial("tcp",address)
			if err != nil {
				continue
			}
		conn.Close()
		fmt.Printf("%d open\n",i)
	}
}
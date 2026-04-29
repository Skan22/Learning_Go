package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	b:= make([]byte,1024)

	for {
		size ,err := conn.Read(b[:])
		if err ==io.EOF{
			log.Println("Client disconnected")
			break
		}
		if err!=nil{
			log.Println("Unexpected Error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))
		log.Println("Writing Data")

		if _,err :=conn.Write(b[0:size]);err!=nil{
			log.Fatalln("Unable to Write Data")
		}
		log.Println("Data Written Successfully")
	
	
	}
}



func main() {
	listener,err := net.Listen("tcp",":20080")
	if err!=nil {
		log.Fatalln("Unable to bind port")
	}
	log.Println("Listening on port 0.0.0.0:20080")
	for {
		// main thread blocks here on listener.Accept()
		conn,err:=listener.Accept()
		log.Println("Received Connection")
		if err!=nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}

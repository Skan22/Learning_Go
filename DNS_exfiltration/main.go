package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/layers"
)

func main(){
	devices,err:=pcap.FindAllDevs()

	if err !=nil {
		fmt.Printf("Error Reading Devices")
	}
	fmt.Println("Machine devices :")
	for _,device := range devices {
		fmt.Println("- ",device.Name)
	}


	handle,err := pcap.OpenLive("any",65535,true,pcap.BlockForever)
	if err !=nil{
		log.Fatal("Error Reading packet:",err)
	}
	defer handle.Close()

	err = handle.SetBPFFilter("udp and port 53")
		if err !=nil{
		log.Fatal("Error Setting Filters :",err)
	}
	packetsource :=gopacket.NewPacketSource(handle,handle.LinkType())
	for packet := range packetsource.Packets() {
    	fmt.Println(packet)
	}


}

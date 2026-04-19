package config

import (
	"os"
	"log"
)

var (
	DeviceName 	string
	es_index 	string
	es_docType 	string
	es_server	string
	err      	error
	handle   	*pcap.Handle
	InetAddr 	string
	SrcIP    	string
	DstIP    	string
)

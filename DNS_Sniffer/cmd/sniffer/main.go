package main 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)


var(
	deviceName string 
	InetAdrr string
	
	handle *pcap.Handle 
	
	err error 
	
	SrcIP string
	DestIP string
	// DB params go here (kafka or clickhouse )
)


type DnsMsg struct {
	Timestamp string
	SourceIP	string
	DestinationIP string
	DNS_Query string
	DNS_Answer []string
	DNS_Answer_TTL []string
	NumberOfAnswers string
	DNS_Response_Code string
	DnsOpCode string

}
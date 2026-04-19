package capture

import (
	//internal imports
	"log"

	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/DNS_Sniffer/internal/config"
	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/internal/config"
	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/internal/dns"

	//exernal imports
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

	func main(){
		handle,err := pcap.OpenLive(config.DeviceName,1600,false,pcap.BlockForever)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()
		




	}



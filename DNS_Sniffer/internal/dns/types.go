package dns



type DnsRecord struct {
	Timestamp string
	SourceIP	string
	DestinationIP string
	DNS_Query string
	DNS_Answer []string
	DNS_Answer_TTL []string
	DNS_Response_Code string
	DnsOpCode string

}
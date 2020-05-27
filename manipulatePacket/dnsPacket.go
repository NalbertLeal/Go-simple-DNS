package manipulatePacket

import "errors"

type DNSPacket struct {
	Header      *DnsHeader
	Questions   []*DnsQuestion
	Answers     []*DnsRecord
	Authorities []*DnsRecord
	Resources   []*DnsRecord
}

func NewDNSPacket() *DNSPacket {
	dnsPacket := new(DNSPacket)

	dnsPacket.Header = new(DnsHeader)

	return dnsPacket
}

func DNSPacketfromBuffer(rawDnsPackage *RawDnsPackage) (*DNSPacket, error) {
	dnsPacket := NewDNSPacket()
	dnsPacket.Header.Read(rawDnsPackage)

	var i uint16
	for i = 0; i < dnsPacket.Header.Questions; i++ {
		question := NewDnsQuestion("", 0)
		err := question.Read(rawDnsPackage)
		if err != nil {
			return nil, errors.New("> Error: While Creating the DNS packet (reading the questions).\n" + err.Error())
		}
		dnsPacket.Questions = append(dnsPacket.Questions, question)
	}

	for i = 0; i < dnsPacket.Header.Answers; i++ {
		record, err := NewDnsRecord(rawDnsPackage)
		if err != nil {
			return nil, errors.New("> Error: While Creating the DNS packet (reading the answers).\n" + err.Error())
		}
		dnsPacket.Answers = append(dnsPacket.Answers, record)
	}

	for i = 0; i < dnsPacket.Header.AuthoritativeEntries; i++ {
		record, err := NewDnsRecord(rawDnsPackage)
		if err != nil {
			return nil, errors.New("> Error: While Creating the DNS packet (reading the authorities).\n" + err.Error())
		}
		dnsPacket.Authorities = append(dnsPacket.Authorities, record)
	}

	for i = 0; i < dnsPacket.Header.ResourceEntries; i++ {
		record, err := NewDnsRecord(rawDnsPackage)
		if err != nil {
			return nil, errors.New("> Error: While Creating the DNS packet (reading the resources).\n" + err.Error())
		}
		dnsPacket.Resources = append(dnsPacket.Resources, record)
	}

	return dnsPacket, nil
}

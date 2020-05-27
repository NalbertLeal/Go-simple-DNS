package manipulatePacket

import (
	"errors"
	"strconv"
)

type DnsRecord struct {
	Domain     string
	Address    uint32 // ONLY TO QUERY TYPE A
	QueryType  uint16
	DataLength uint16 // ONLY TO QUERY TYPE DIFFERENT OF A
	Ttl        uint32
}

func NewDnsRecord(rawDnsPackage *RawDnsPackage) (*DnsRecord, error) {
	domain := ""
	domain, err := rawDnsPackage.ReadQueryName()
	if err != nil {
		return nil, errors.New("> Error: While reading the domain in function NewDnsRecord.\n" + err.Error())
	}

	queryType, err := rawDnsPackage.ReadUInt16()
	if err != nil {
		return nil, errors.New("> Error: While reading the query type in function NewDnsRecord.\n" + err.Error())
	}

	// skip class
	rawDnsPackage.ReadUInt16()

	ttl, err := rawDnsPackage.ReadUInt32()
	if err != nil {
		return nil, errors.New("> Error: While reading the ttl in function NewDnsRecord.\n" + err.Error())
	}

	dataLength, err := rawDnsPackage.ReadUInt16()
	if err != nil {
		return nil, errors.New("> Error: While reading the data length in function NewDnsRecord.\n" + err.Error())
	}

	dnsRecord := new(DnsRecord)
	switch queryType {
	case 1: // TYPE A
		dnsRecord.Domain = domain
		dnsRecord.Address, err = rawDnsPackage.ReadUInt32()
		if err != nil {
			return nil, errors.New("> Error: While reading the address in function NewDnsRecord.\n" + err.Error())
		}
		dnsRecord.QueryType = queryType
		dnsRecord.Ttl = ttl
		break
	default:
		dnsRecord.Domain = domain
		dnsRecord.QueryType = queryType
		dnsRecord.DataLength = dataLength
		dnsRecord.Ttl = ttl
	}

	return dnsRecord, nil
}

func (self *DnsRecord) AddressString() string {
	address := self.Address

	byte1 := int((address & 4278190080) >> 24)
	byte2 := int((address & 16711680) >> 16)
	byte3 := int((address & 65280) >> 8)
	byte4 := int(address & 255)

	return strconv.Itoa(byte1) + "." + strconv.Itoa(byte2) + "." + strconv.Itoa(byte3) + "." + strconv.Itoa(byte4)
}

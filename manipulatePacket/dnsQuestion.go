package manipulatePacket

import (
	"errors"
)

type DnsQuestion struct {
	Name      string
	Querytype uint16
}

func NewDnsQuestion(name string, queryType uint16) *DnsQuestion {
	dnsQuestion := new(DnsQuestion)

	dnsQuestion.Name = name
	dnsQuestion.Querytype = queryType

	return dnsQuestion
}

func (self *DnsQuestion) Read(rawDnsPackage *RawDnsPackage) error {
	var err error
	self.Name, err = rawDnsPackage.ReadQueryName()
	if err != nil {
		return errors.New(">  Error: While read DNS question name.\n" + err.Error())
	}
	self.Querytype, err = rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New(">  Error: While read DNS question query type.\n" + err.Error())
	}

	// skip the class
	rawDnsPackage.ReadUInt16()

	return nil
}

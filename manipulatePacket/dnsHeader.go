package manipulatePacket

import (
	"errors"
)

type DnsHeader struct {
	Id uint16 // 16 bits

	RecursionDesired    bool  // 1 bit
	TrucetedMessage     bool  // 1 bit
	AuthoritativeAnswer bool  // 1 bit
	Opcode              uint8 // 4 bits
	Response            bool  // 1 bit

	Rescode            uint8 // 4 bits constant from file resultCode.go
	ChekingDisabled    bool  // 1 bit
	AuthedData         bool  // 1 bit
	Z                  bool  // 1 bit
	RecursionAvaliable bool  // 1 bit

	Questions            uint16 // 16 bit
	Answers              uint16 // 16 bit
	AuthoritativeEntries uint16 // 16 bit
	ResourceEntries      uint16 // 16 bit
}

func NewDnsHeader() *DnsHeader {
	dnsHeader := new(DnsHeader)
	return dnsHeader
}

func (self *DnsHeader) Read(rawDnsPackage *RawDnsPackage) error {
	var err error
	self.Id, err = rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New("> Error: While get the transaction id. Function: NewDnsHeader.\n" + err.Error())
	}

	flags, err := rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New("> Error: While get the flags. Function: NewDnsHeader.\n" + err.Error())
	}

	byte1 := byte(flags >> 8)
	byte2 := byte(flags & 0xFF)

	self.RecursionDesired = (byte1 & 1) > 0
	self.TrucetedMessage = (byte1 & 2) > 0
	self.AuthoritativeAnswer = (byte1 & 4) > 0
	self.Opcode = uint8(byte1 & 120)
	self.Response = (byte1 & 128) > 0

	self.Rescode = uint8(byte2 & 15)
	self.ChekingDisabled = int8(byte2&16) > 0
	self.AuthedData = int8(byte2&32) > 0
	self.Z = int8(byte2&64) > 0
	self.RecursionAvaliable = int8(byte2&128) > 0

	self.Questions, err = rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New("> Error: While get the DNS header questions. Function: NewDnsHeader.\n" + err.Error())
	}
	self.Answers, err = rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New("> Error: While get the DNS header answers. Function: NewDnsHeader.\n" + err.Error())
	}
	self.AuthoritativeEntries, err = rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New("> Error: While get the DNS header authoritative entries. Function: NewDnsHeader.\n" + err.Error())
	}
	self.ResourceEntries, err = rawDnsPackage.ReadUInt16()
	if err != nil {
		return errors.New("> Error: While get the DNS header resource entries. Function: NewDnsHeader.\n" + err.Error())
	}

	return nil
}

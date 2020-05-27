package manipulatePacket

import (
	"errors"
	"strings"
)

type RawDnsPackage struct {
	Data     []byte
	Position uint
}

func NewRawDnsPackage(data []byte) *RawDnsPackage {
	dnsPackage := new(RawDnsPackage)

	dnsPackage.Data = data
	dnsPackage.Position = 0

	return dnsPackage
}

func (self *RawDnsPackage) Step(steps uint) {
	self.Position += steps
}

func (self *RawDnsPackage) Seek(position uint) {
	self.Position = position
}

func (self *RawDnsPackage) Read() (byte, error) {
	if self.Position > 511 {
		return 0x00, errors.New("> Error: Invalid dns package input size.  The function that caused the error was: RawDnsPackage.Read.")
	}
	response := self.Data[self.Position]
	self.Position += 1
	return response, nil
}

func (self *RawDnsPackage) Get(position uint) (byte, error) {
	if position > 511 {
		return 0x00, errors.New("> Error: Invalid dns package input size.  The function that caused the error was: RawDnsPackage.Get.")
	}
	return self.Data[position], nil
}

func (self *RawDnsPackage) GetRange(start uint, length uint) ([]byte, error) {
	if start+length > 511 {
		return nil, errors.New("> Error: Invalid dns package input size. The function that caused the error was: RawDnsPackage.GetRange.")
	}
	return self.Data[start : start+length], nil
}

func (self *RawDnsPackage) ReadUInt16() (uint16, error) {
	byte1, err := self.Read()
	if err != nil {
		return 0, errors.New("> Error: The dns package size is 512, after call self.Read the read position passed 512. The function that caused the error was: RawDnsPackage.ReadInt16.")
	}
	byte2, err := self.Read()
	if err != nil {
		return 0, errors.New("> Error: The dns package size is 512, after call self.Read the read position passed 512. The function that caused the error was: RawDnsPackage.ReadInt16.")
	}

	var response uint16
	response = uint16(byte1)<<8 | uint16(byte2)
	return response, nil
}

func (self *RawDnsPackage) ReadUInt32() (uint32, error) {
	byte1, err := self.Read()
	if err != nil {
		return 0, errors.New("> Error: The dns package size is 512, after call self.Read the read position passed 512. The function that caused the error was: RawDnsPackage.ReadInt32.")
	}
	byte2, err := self.Read()
	if err != nil {
		return 0, errors.New("> Error: The dns package size is 512, after call self.Read the read position passed 512. The function that caused the error was: RawDnsPackage.ReadInt32.")
	}
	byte3, err := self.Read()
	if err != nil {
		return 0, errors.New("> Error: The dns package size is 512, after call self.Read the read position passed 512. The function that caused the error was: RawDnsPackage.ReadInt32.")
	}
	byte4, err := self.Read()
	if err != nil {
		return 0, errors.New("> Error: The dns package size is 512, after call self.Read the read position passed 512. The function that caused the error was: RawDnsPackage.ReadInt32.")
	}

	var response uint32
	response = uint32(byte1)<<24 | uint32(byte2)<<16 | uint32(byte3)<<8 | uint32(byte4)
	return response, nil
}

func (self *RawDnsPackage) ReadQueryName() (string, error) {
	response := ""

	positionNow := self.Position
	jumped := false

	delimeter := ""
	for {
		// Here is at the beginning of the label.
		// Label start with label length byte.
		length, err := self.Get(positionNow)
		if err != nil {
			return "", errors.New("> Error: While reading the domain/label length.\n" + err.Error())
		}

		// If the 2 most significant bits of length are set, if represents a
		// jump to some other offset.
		if (length & 0xC0) == 0xC0 {
			if !jumped {
				self.Seek(positionNow + 2)
			}

			// Read another byte, calculate offset and
			// jump by updating the positionNow
			byte2, err := self.Get(positionNow + 1)
			if err != nil {
				return "", errors.New("> Error: While reading the domain.\n" + err.Error())
			}

			offset := ((int16(length) ^ 0xC0) << 8) | int16(byte2)
			positionNow = uint(offset)

			jumped = true
		} else {
			positionNow += 1

			if length == 0 {
				break
			}

			response += delimeter

			strBuffer, err := self.GetRange(positionNow, uint(length))
			if err != nil {
				return "", errors.New("> Error: While reading the domain.\n" + err.Error())
			}
			response += strings.ToLower(string(strBuffer))

			delimeter = "."

			positionNow += uint(length)
		}
	}

	if !jumped {
		self.Seek(positionNow)
	}

	return response, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"

	mp "./manipulatePacket"
)

func main() {

	// file, err := os.Open("response_packet.txt")
	file, err := os.Open("testFiles/response_packet.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	fikeStats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var size int64 = fikeStats.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rawDnsPackage := mp.NewRawDnsPackage(bytes)

	dnsPacket, err := mp.DNSPacketfromBuffer(rawDnsPackage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(dnsPacket.Questions); i++ {
		address := dnsPacket.Answers[i].Address

		fmt.Println("> Google address:")
		fmt.Println("\tAddress as uint32 (4 bytes, each one with a 0..255 value): ", address)
		fmt.Println("\tAddress parsed from uint32 to IP string: ", dnsPacket.Answers[i].AddressString())
	}
}

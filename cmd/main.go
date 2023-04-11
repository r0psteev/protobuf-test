package main

import (
	//"fmt"
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	sliverpb "example.com/protobuf-test/proto"
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg) 
	}
}
func saveToFile(name string, data []byte) {
	err := os.WriteFile(name, data, 0666)
	failOnError("Failed writing to file", err)
}

func main() {
	test := sliverpb.Ping{
		Nonce: 0xbeef,
		Timeout: 2,
		BeaconID: "uuid-bot",
	}

	// can i marshal it too ?
	// marshal to JSON
	data, err := protojson.Marshal(&test)
	failOnError("failed to Marshal to JSON", err)
	saveToFile("ping.json", data)


	// marshal to wire-format binary
	data2, err := proto.Marshal(&test)
	failOnError("failed to Marshal to JSON", err)
	saveToFile("ping.wire", data2)

	// marshal to binary little endian
	// binary.Write(w io.Writer, order ByteOrder, data any) error
	header := new(bytes.Buffer)
	err = binary.Write(header, binary.LittleEndian, uint32(0xdeadbeef))
	failOnError("failed to write to header buffer", err)
	err = binary.Write(header, binary.BigEndian, uint32(0x1337cafe))
	failOnError("failed to write to header buffer", err)
	err = binary.Write(header, binary.BigEndian, data2)
	failOnError("failed to write to header buffer", err)

	saveToFile("ping.wire.lte", header.Bytes())
}

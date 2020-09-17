package tlsx

import (
	"encoding/hex"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TestClientHello(t *testing.T) {

	packet := "48d343aac4b8f018982a38be0800450002390000400040063dddc0a8b20d58c62f66fa5401bb22afc03a4f66c79a80180814af8900000101080a571eb2fab5d8c2d71603010200010001fc03030c4c5a78621a9d1f687fda02e40b01897bc32fefdd8f66612360cb40f186e29f2075aae50aca7bd3d7db205ce25ddc409a902578c8b5b6b1eb1f1cbe19cc02a45a0034130113021303c02cc02bc024c023c00ac009cca9c030c02fc028c027c014c013cca8009d009c003d003c0035002fc008c012000a0100017fff010001000000001a00180000156463382e733234302e6d656574726963732e6e657400170000000d0018001604030804040105030203080508050501080606010201000500050100000000001200000010000e000c02683208687474702f312e31000b00020100003300260024001d00200bd78e1307f42e2e1ce25309a2191a31f8436c270476f7808171d787c7d2b25f002d00020101002b0009080304030303020301000a000a0008001d001700180019001500c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	data, err := hex.DecodeString(packet)
	if err != nil {
		t.Fatal(err)
	}

	p := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.Default)
	hello := GetClientHello(p)
	if hello == nil {
		t.Fatal("hello should not be nil")
	}

	if hello.Version != 769 {
		t.Fatal("version should be 769 (TLS 1.0), got", hello.Version)
	}

	for i, e := range []uint16{4865, 4866, 4867, 49196, 49195, 49188, 49187, 49162, 49161, 52393, 49200, 49199, 49192, 49191, 49172, 49171, 52392, 157, 156, 61, 60, 53, 47, 49160, 49170, 10} {
		if uint16(hello.CipherSuites[i]) != e {
			t.Fatal("incorrect cipher, got", hello.AllExtensions[i], "expected", e)
		}
	}

	for i, e := range []uint16{65281, 0, 23, 13, 5, 18, 16, 11, 51, 45, 43, 10, 21} {
		if hello.AllExtensions[i] != e {
			t.Fatal("incorrect extension, got", hello.AllExtensions[i], "expected", e)
		}
	}
}

func TestClientHelloBasic(t *testing.T) {

	packet := "48d343aac4b8f018982a38be0800450002390000400040063dddc0a8b20d58c62f66fa5401bb22afc03a4f66c79a80180814af8900000101080a571eb2fab5d8c2d71603010200010001fc03030c4c5a78621a9d1f687fda02e40b01897bc32fefdd8f66612360cb40f186e29f2075aae50aca7bd3d7db205ce25ddc409a902578c8b5b6b1eb1f1cbe19cc02a45a0034130113021303c02cc02bc024c023c00ac009cca9c030c02fc028c027c014c013cca8009d009c003d003c0035002fc008c012000a0100017fff010001000000001a00180000156463382e733234302e6d656574726963732e6e657400170000000d0018001604030804040105030203080508050501080606010201000500050100000000001200000010000e000c02683208687474702f312e31000b00020100003300260024001d00200bd78e1307f42e2e1ce25309a2191a31f8436c270476f7808171d787c7d2b25f002d00020101002b0009080304030303020301000a000a0008001d001700180019001500c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	data, err := hex.DecodeString(packet)
	if err != nil {
		t.Fatal(err)
	}

	p := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.Default)
	hello := GetClientHelloBasic(p)
	if hello == nil {
		t.Fatal("hello should not be nil")
	}

	if hello.Version != 769 {
		t.Fatal("version should be 769 (TLS 1.0), got", hello.Version)
	}

	for i, e := range []uint16{4865, 4866, 4867, 49196, 49195, 49188, 49187, 49162, 49161, 52393, 49200, 49199, 49192, 49191, 49172, 49171, 52392, 157, 156, 61, 60, 53, 47, 49160, 49170, 10} {
		if uint16(hello.CipherSuites[i]) != e {
			t.Fatal("incorrect cipher, got", hello.AllExtensions[i], "expected", e)
		}
	}

	for i, e := range []uint16{65281, 0, 23, 13, 5, 18, 16, 11, 51, 45, 43, 10, 21} {
		if hello.AllExtensions[i] != e {
			t.Fatal("incorrect extension, got", hello.AllExtensions[i], "expected", e)
		}
	}
}

func BenchmarkGetClientHello(b *testing.B) {

	packet := "48d343aac4b8f018982a38be0800450002390000400040063dddc0a8b20d58c62f66fa5401bb22afc03a4f66c79a80180814af8900000101080a571eb2fab5d8c2d71603010200010001fc03030c4c5a78621a9d1f687fda02e40b01897bc32fefdd8f66612360cb40f186e29f2075aae50aca7bd3d7db205ce25ddc409a902578c8b5b6b1eb1f1cbe19cc02a45a0034130113021303c02cc02bc024c023c00ac009cca9c030c02fc028c027c014c013cca8009d009c003d003c0035002fc008c012000a0100017fff010001000000001a00180000156463382e733234302e6d656574726963732e6e657400170000000d0018001604030804040105030203080508050501080606010201000500050100000000001200000010000e000c02683208687474702f312e31000b00020100003300260024001d00200bd78e1307f42e2e1ce25309a2191a31f8436c270476f7808171d787c7d2b25f002d00020101002b0009080304030303020301000a000a0008001d001700180019001500c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	data, err := hex.DecodeString(packet)
	if err != nil {
		b.Fatal(err)
	}

	p := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.Default)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		hello := GetClientHello(p)
		if hello == nil {
			b.Fatal("hello should not be nil")
		}
	}
}

func BenchmarkGetClientHelloBasic(b *testing.B) {

	packet := "48d343aac4b8f018982a38be0800450002390000400040063dddc0a8b20d58c62f66fa5401bb22afc03a4f66c79a80180814af8900000101080a571eb2fab5d8c2d71603010200010001fc03030c4c5a78621a9d1f687fda02e40b01897bc32fefdd8f66612360cb40f186e29f2075aae50aca7bd3d7db205ce25ddc409a902578c8b5b6b1eb1f1cbe19cc02a45a0034130113021303c02cc02bc024c023c00ac009cca9c030c02fc028c027c014c013cca8009d009c003d003c0035002fc008c012000a0100017fff010001000000001a00180000156463382e733234302e6d656574726963732e6e657400170000000d0018001604030804040105030203080508050501080606010201000500050100000000001200000010000e000c02683208687474702f312e31000b00020100003300260024001d00200bd78e1307f42e2e1ce25309a2191a31f8436c270476f7808171d787c7d2b25f002d00020101002b0009080304030303020301000a000a0008001d001700180019001500c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	data, err := hex.DecodeString(packet)
	if err != nil {
		b.Fatal(err)
	}

	p := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.Default)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		hello := GetClientHelloBasic(p)
		if hello == nil {
			b.Fatal("hello should not be nil")
		}
	}
}

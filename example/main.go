package main

import (
	"fmt"
	"net"

	"github.com/endobit/oui"
)

func main() {
	const addr = "00:00:0f:01:02:03"

	// Lookup the vendor of an Ethernet MAC.

	fmt.Println(oui.Vendor(addr))

	// Or, lookup the vendor of a net.HardwareAddr.

	mac, err := net.ParseMAC(addr)
	if err != nil {
		fmt.Printf("%q not a mac\n", addr)
		panic(err)
	}

	fmt.Println(oui.VendorFromMAC(mac))
}

// Package oui provides methods to lookup the Vendor string of an
// Organizationally Unique Identifier. For Ethernet MACs the OUI is
// the first 24 bits.
//
// Vendor lookup is done statically from a code generated `struct`
// using a sanitized oui database (https://linuxnet.ca/ieee/oui).
package oui

import (
	"net"
	"strings"
)

//go:generate ./gen

// Vendor returns the Vendor of the OUI s, or the empty string if
// not found. The value s can be either just the OUI or a mac address,
// it can also include or exclude the ":" delimiters. Any of the
// following are valid inputs:
//     0A:0B:0C
//     0a:0b:0c
//     0A0B0C
//     0a0b0c
//     0A:0B:0C:01:02:03
//     0a:0b:0c:01:02:03
//     0A0B0C010203
//     0a0b0c010203
func Vendor(s string) string {
	s = strings.ReplaceAll(s, ":", "")

	switch {
	case len(s) < 6:
		return ""
	case len(s) > 6:
		s = s[0:6]
	}

	s = strings.ToLower(s)

	id, ok := database.ouis[s]
	if !ok {
		return ""
	}

	v, ok := database.vendors[id]
	if !ok {
		return ""
	}

	return v
}

// VendorFromMAC returns the hardware vendor of a net.HardwareAddr.
// Both 48 and 64 bits addresses are valid, as the OUI is the first 24
// bits of both.
func VendorFromMAC(hw net.HardwareAddr) string {
	return Vendor(hw.String())
}

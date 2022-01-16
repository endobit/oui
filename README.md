[![GitHub Actions](https://img.shields.io/github/workflow/status/endobit/oui/test)](https://github.com/endobit/oui/actions?query=workflow%3Atest)
[![Go Report Card](https://goreportcard.com/badge/github.com/endobit/oui)](https://goreportcard.com/report/github.com/endobit/oui)
[![Codecov](https://codecov.io/gh/endobit/oui/branch/main/graph/badge.svg)](https://codecov.io/gh/endobit/oui)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/endobit/oui)


# Vendor lookup from MAC prefixes

A Go module that provides methods to lookup the Vendor string of an
Organizationally Unique Identifier. For Ethernet MACs the OUI is
encoded in the first 24 bits.

Vendor lookup is done statically from a code generated `struct` using
a sanitized oui [database](https://linuxnet.ca/ieee/oui).

---

## Usage

```go
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
		t.Errorf("%v", mac)
	}

	fmt.Println(oui.VendorFromMAC(mac))
}
```

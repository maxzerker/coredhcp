package main

/*
 * Sample DHCPv6 client to test on the local interface
 */

import (
	"flag"
	"log"
	"net"

	"github.com/insomniacslk/dhcp/dhcpv6"
	"github.com/insomniacslk/dhcp/iana"
)

func main() {
	flag.Parse()

	var macString string
	if len(flag.Args()) > 0 {
		macString = flag.Arg(0)
	} else {
		macString = "00:11:22:33:44:55"
	}

	c := dhcpv6.NewClient()
	c.LocalAddr = &net.UDPAddr{
		IP:   net.ParseIP("::1"),
		Port: 546,
	}
	c.RemoteAddr = &net.UDPAddr{
		IP:   net.ParseIP("::1"),
		Port: 547,
	}
	log.Printf("%+v", c)

	mac, err := net.ParseMAC(macString)
	if err != nil {
		log.Fatal(err)
	}
	duid := dhcpv6.Duid{
		Type:          dhcpv6.DUID_LLT,
		HwType:        iana.HwTypeEthernet,
		Time:          dhcpv6.GetTime(),
		LinkLayerAddr: mac,
	}

	conv, err := c.Exchange("lo", dhcpv6.WithClientID(duid))
	for _, p := range conv {
		log.Print(p.Summary())
	}
	if err != nil {
		log.Fatal(err)
	}
}

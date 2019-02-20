package gocqlec2multiregiontranslator

import (
	"net"
	gocql "github.com/gocql/gocql"
)

type EC2MultiRegionAddressTranslator struct {
        gocql.AddressTranslator
}

func (translator EC2MultiRegionAddressTranslator) Translate(addr net.IP, port int) (net.IP, int) {
	return net.ParseIP("10.10.10.10"), 1
}

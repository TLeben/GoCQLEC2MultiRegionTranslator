package gocqlec2multiregiontranslator

import (
	"net"
	gocql "github.com/gocql/gocql"
)

type EC2MultiRegionAddressTranslator struct {
        gocql.AddressTranslator
}

func (translator EC2MultiRegionAddressTranslator) Translate(addr net.IP, port int) (net.IP, int) {
	hostname, _ := net.LookupAddr(addr.String())
	privateip, _ := net.LookupHost(hostname[0])
	return net.ParseIP(privateip[0]), port
} 

package main

import (
	"github.com/zachaller/GoCQLEC2MultiRegionTranslator/gocqlec2multiregiontranslator"
	"net"
	"fmt"
)

func main() {
	var a = gocqlec2multiregiontranslator.EC2MultiRegionAddressTranslator{}
	fmt.Println(a.Translate(net.ParseIP("34.244.79.21"), 9042))
}

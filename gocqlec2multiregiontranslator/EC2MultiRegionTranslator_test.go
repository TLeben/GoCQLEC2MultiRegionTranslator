package gocqlec2multiregiontranslator
import (
    "testing"
    "net"
    "fmt"
)

func TestLookup(t *testing.T) {
  var a = EC2MultiRegionAddressTranslator{}
  out, outp := a.Translate(net.ParseIP("34.244.79.21"), 9042)
  fmt.Println(out, outp)
}

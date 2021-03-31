package main

import (
	ipt "github.com/coreos/go-iptables/iptables"
	"os"
)

func main() {
	tables, _ := ipt.New()
	chains, _ := tables.ListChains(os.Args[1])
	for _, chain := range chains {
		println(chain)
	}
}

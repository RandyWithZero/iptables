package main

import (
	ipt "github.com/coreos/go-iptables/iptables"
	"os"
)

func main() {
	tables, _ := ipt.New()
	s := os.Args[1]
	chains, _ := tables.ListChains(s)
	for _, chain := range chains {
		println(chain)
		list, _ := tables.List(s, chain)
		println(list)
	}
}

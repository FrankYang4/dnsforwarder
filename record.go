package main

import (
	"net"
)

type Record struct {
	Domain string
	Addrs  []net.IP
	Ttl    int
}

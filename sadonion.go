package main

import (
	"fmt"
	"os"
	"golang.org/x/net/proxy"
	"strings"
	"strconv"
)

services := map[int]string {
	22 :	"SSH Server",
	53 :	"DNS Server",
	70 :	"Gopher Server",
	80 :    "HTTP Server",
	443 :	"HTTPS Server"
}

const (
	TOR_ADDR = "127.0.0.1:9050"
)

func oErr() {
	fmt.Fprintln(os.Stderr, os.Args[1], " is not a valid hidden service address.")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: sadonion <dot onion address>")
		os.Exit(1)
	}

	if len(os.Args[1]) != 22 {
		oErr()
	}

	if strings.HasSuffix(os.Args[1], ".onion") == false {
		oErr()
	}

	dialer, err := proxy.SOCKS5("tcp", TOR_ADDR, nil, proxy.Direct)

	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to Tor SOCKS5 proxy: ", err)
		os.Exit(1)
	}

	for x := 1; x < 1000; x++ {
		strport := strconv.Itoa(x)
		conn, dialerr := dialer.Dial("tcp", os.Args[1]  + ":" + strport)

		if dialerr == nil {
			defer conn.Close()
			fmt.Println("Port ", strport, " is live!")
		}

	}
}

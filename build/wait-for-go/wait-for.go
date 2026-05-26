package main

import (
	"flag"
	"os"
	"strings"
	"time"
)

// Usage example
// go run build/wait-for-go/wait-for.go -u -t 10 localhost:8125 && go run main.go
func canConnect(host, port, protocol string) bool { _ = "STUB: not implemented"; return false }

// When using UDP do a quick check to see if something is listening on the
// given port to return an error as soon as possible.

func main() {
	protocol := "tcp"
	udp := flag.Bool("u", false, "check for udp")
	timeout := flag.Int("t", 60, "Timeout in seconds")
	flag.Parse()
	if *udp {
		protocol = "udp"
	}
	hostport := flag.Args()

	hostportArray := strings.SplitN(hostport[0], ":", 2)
	host := hostportArray[0]
	port := hostportArray[1]
	for index := 0; index < *timeout; index++ {
		connected := canConnect(host, port, protocol)
		if connected {
			os.Exit(0)
		}
		time.Sleep(time.Second)
	}
	os.Exit(1)
}

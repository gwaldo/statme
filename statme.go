// Send metrics

// Intended Usage:
// statme [ -s | --statsd | -g | --graphite ] [ -tcp | -udp ] [ -p port ] [ -h host ] metric value type
//	or
// statme [ -c | --config file ] metric value type
//
// It would be helpful to be able to have some defaults or specify a config file
//
//
//
// http://golang.org/pkg/net/
// http://stackoverflow.com/questions/17915952/golang-tcp-socket-does-not-send-message-after-write-immediately
// 
// https://www.hostedgraphite.com/accounts/profile/
// http://docs.hostedgraphite.com/languageguide/lg_shell.html
//
// http://spf13.com/presentation/first-go-app?utm_content=buffer2ee4e&utm_medium=social&utm_source=twitter.com&utm_campaign=buffer

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "carbon.hostedgraphite.com:2003")
	if err != nil {
		// handle error
		fmt.Println("There was an error")
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "706120df-49e5-46ef-b3cb-f4460f3f3975.testing.golang.tcp 1")
}
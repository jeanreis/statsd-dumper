package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	build := "dev"

	port := os.Getenv("STATSD_PORT")
	if port == "" {
		port = "8125"
	}

	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port specified", port)
	}

	socket, err := net.ListenUDP("udp", &net.UDPAddr{Port: portAsInt, IP: net.ParseIP("0.0.0.0")})

	if err != nil {
		log.Fatalf("Could not start listening on %s", socket.LocalAddr().String())
	}

	defer socket.Close()

	fmt.Printf("Starting statsd-dumper (build %s) on port %s...\n\n", build, socket.LocalAddr().String())

	for {
		buffer := make([]byte, 1024)
		length, _, err := socket.ReadFromUDP(buffer[:])
		if err != nil {
			continue
		}
		go serve(buffer[:length])
	}

}

func serve(buffer []byte) {
	fmt.Printf("%s", buffer)
}

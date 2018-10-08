package main

import "net"

// CacheDefault - Default number of records to Cache before expiring a record
const CacheDefault = 100

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		// conn, err := listen.Accept()
		// if err != nil {
		// 	// handle error
		// }
		// go handleConnection(conn)
	}

}

package main

import "net"

// CacheDefault - Default number of records to Cache before expiring a record
const CacheDefault = 100

type p2pServerHandler interface {
	OnRecieve(request string) (string, error)
	OnResolve(request string, response string)
	OnFail(request string, response string)
}

type p2pServer struct {
	handler p2pServerHandler
}

func (s p2pServer) recieve() string {
	return ""
}

func (s p2pServer) send(ip string, port int, message string) {

}

func (s p2pServer) handle() {
	request := s.recieve()
	response, err := s.handler.OnRecieve(request)
	if err != nil && response == "" {
		s.send("", 0, response)
		s.handler.OnResolve(request, response)
	} else {
		s.handler.OnFail(request, response)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

}

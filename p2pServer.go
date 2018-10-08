package main

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

package main

// EchoServerHandler - p2pServerHandler which simply echo's/prints what it does and return its ip
type EchoServerHandler struct {
}

// OnRecieve - return own server's ip upton reciecing a request
func (s EchoServerHandler) OnRecieve(request string) (string, error) {
	return "", nil
}

// OnResolve - return own server's ip upton reciecing a request
func (s EchoServerHandler) OnResolve(request string, response string) {
}

// OnFail - return own server's ip upton reciecing a request
func (s EchoServerHandler) OnFail(request string, response string) {
}

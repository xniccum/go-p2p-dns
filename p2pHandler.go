package main

type p2pServerHandler interface {
	OnRecieve(request string) (string, error)
	OnResolve(request string, response string)
	OnFail(request string, response string)
}

# go-p2p-dns

A simple P2P DNS implementation in Go; the is also made flexible enough such that diffrent servers can have diffrent implimentations, allowing anyone to set up a server & bussinesses to implement particular one as they wish

## TODO

* Generic Database for datastore cache
* Examples to implement
  * Echo Server - on which simply echos its own IP no matter what
  * Proxy Server - one which, in the even of failure, defaults to actual current DNS
  * Full Server - one which actually checks its peers in the event of failure
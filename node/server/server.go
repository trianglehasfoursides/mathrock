package server

import "github.com/quic-go/quic-go/http3"

var Server = &http3.Server{
	Addr:    "",
	Port:    9000,
	Handler: route(),
}

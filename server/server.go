package server

type Server interface {
	NewServer()
	SetAddr(addr string)
	Start()
}

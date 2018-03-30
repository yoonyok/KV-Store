package structs

import "net/rpc"

type Store struct {
	Address   string
	RPCClient *rpc.Client
	IsLeader  bool
}

type WriteRequest struct {
	Key   int
	Value string
}

type ACK struct {
	Acknowledged bool
	Key          int
	Value        string
	Error        error
}

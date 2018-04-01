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
}

type CandidateInfo struct {
	LogLength         int
	NumberOfCommitted int
}

type LogEntry struct {
	Index       int
	Key         int
	Value       string
	IsCommitted bool
}

type LogEntries struct {
	Current  LogEntry
	Previous LogEntry
}

type StoreInfo struct {
	Address  string
	IsLeader bool
}

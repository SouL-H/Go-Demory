package main

import (
	"context"
	"io"
	"sync"

	"github.com/hashicorp/raft"
)

type demory struct{
	DB map[string]string
	mutex sync.RWMutex 
}

func NewDemory() *demory{
	return &demury{
		
	}
}

func(d demory) Apply(log *raft.Log) interface{}{
	panic("implement")
}
func(d *demory) Snapshot()(raft.FSMSnapshot,error){
	panic("implement")
}
func(d *demory) Restore(closer io.ReadCloser) error{
	panic("implement")
}

type rpcInterface struct{
	raft *raft.Raft
	demory *demory
	proto.UnimplementedDemoryServer
}
func Put(ctx context.Context, req *proto.Putrequest)(*emptypb.Empty,error)


func (r rpcInterface) Get(ctx context.Context,req *proto.GetRequest)
var _ raft.FSM = &demory{}
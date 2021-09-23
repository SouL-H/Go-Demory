package main

import (
	"log"
	"os"
	"path/filepath"

	boltdb "github.com/boltdb/bolt"
	"github.com/hashicorp/raft"
)
func main(){
	config :=raft.DefaultConfig()
	config.LocalID=raft.ServerID("test")
	fsm:=NewDemory()
	basedir :=filepath.Join("/tmp","test")
	mkdirErr :=os.MkdirAll(basedir,os.ModePerm)
	if mkdirErr != nil{
		log.Fatalf("mkdir error  %v", mkdirErr)
	}
	boltdb.NewBoltStore()
	logStore, logStoreErr:= boltdb.NewBoltStore(filepath.Join(basedir,"logs.dat"))
	if logStoreErr !=nil{
		log.Fatalf("logStore error %v", logStoreErr)
	}
	
	
	raft.NewRaft(config,fsm,logStore,stableStore,snapshotStore,transport)
}
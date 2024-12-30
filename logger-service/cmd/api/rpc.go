package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"logger-service/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

type RPCResult struct {
	Resp string
}

func init() {
	gob.Register(RPCPayload{})
}

func (r *RPCServer) LogInfo(rpc *RPCPayload, resp *RPCResult) error {
	fmt.Println("Starting logging rpc- service here...!")
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      rpc.Name,
		Data:      rpc.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println("error while writing data in mongodb", err)
		return err
	}

	res := RPCResult{
		Resp: "Logging from rpc server " + rpc.Name,
	}
	*resp = res
	return nil
}

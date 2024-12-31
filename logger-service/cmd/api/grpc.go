package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"logger-service/logs/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServicesServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	res := &logs.LogResponse{Result: "logged"}
	return res, nil
}

func (app *Config) gRPCListener() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC : %v", err)
	}

	ser := grpc.NewServer()

	logs.RegisterLogServicesServer(ser, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port : %s", gRpcPort)

	err = ser.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

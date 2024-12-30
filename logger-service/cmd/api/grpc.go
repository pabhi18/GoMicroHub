package main

import (
	"context"
	"logger-service/data"
	"logger-service/logs/logs"
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

	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

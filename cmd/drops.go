package main

import (
	"bytes"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type Drop struct {
	interval time.Duration
	url      string
}

type dropper struct {
	logger *zap.Logger
}

func (d dropper) Drop(drop Drop) {
	ticker := time.NewTicker(drop.interval)
	for _ = range ticker.C {
		ping := Ping{
			ID:        uuid.NewV4().String(),
			Message:   "ping",
			TimeStamp: time.Now().Format(time.RFC3339Nano),
		}
		pingJson, err := json.Marshal(ping)
		if err != nil {
			d.logger.Error("error on ping", zap.Error(err))
		}
		resp, err := http.Post(drop.url, "application/json", bytes.NewBuffer(pingJson))
		if err != nil {
			d.logger.Error("got error from api", zap.Error(err))
		}
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			d.logger.Error("absent response body", zap.Error(err))
		}
		var pong Ping
		json.Unmarshal(responseData, &pong)
		d.logger.Info("got ping", zap.String("ping_uuid", pong.ID), zap.String("message", pong.Message))
	}
}

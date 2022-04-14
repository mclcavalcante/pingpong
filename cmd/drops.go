package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
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
			continue
		}
		resp, err := http.Post(drop.url, "application/json", bytes.NewBuffer(pingJson))
		if err != nil {
			d.logger.Error("got error from api", zap.Error(err))
			continue
		}
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			d.logger.Error("absent response body", zap.Error(err))
			continue
		}
		var pong Ping
		json.Unmarshal(responseData, &pong)
		d.logger.Info("got ping", zap.String("ping_uuid", pong.ID), zap.String("message", pong.Message))
	}
}

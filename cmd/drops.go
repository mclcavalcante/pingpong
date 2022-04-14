package main

import (
	"encoding/json"
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
		resp, err := http.Get(drop.url)
		if err != nil {
			d.logger.Error("got error from api", zap.Error(err))
		}
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			d.logger.Error("absent response body", zap.Error(err))
		}
		var ping Ping
		json.Unmarshal(responseData, &ping)
		d.logger.Info("got ping", zap.String("ping_uuid", ping.ID), zap.String("message", ping.Message))
	}
}

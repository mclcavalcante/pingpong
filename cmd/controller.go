package main

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type Ping struct {
	ID        string `json:"ID"`
	Message   string `json:"message"`
	TimeStamp string `json:"@timestamp"`
}

type controller struct {
	logger *zap.Logger
}

func (c *controller) pingFunc(rw http.ResponseWriter, req *http.Request) {
	logger := c.logger
	defer func(start time.Time) {
		duration := time.Since(start)
		logger.Info("ping function runtime", zap.Duration("duration", duration))
	}(time.Now())
	reqBody, _ := ioutil.ReadAll(req.Body)
	var ping Ping
	json.Unmarshal(reqBody, &ping)
	u1 := uuid.NewV4()
	pong := Ping{
		ID:        u1.String(),
		Message:   "pong",
		TimeStamp: time.Now().Format(time.RFC3339Nano),
	}
	json.NewEncoder(rw).Encode(pong)
	rw.WriteHeader(200)
}

func (c *controller) pongFunc(w http.ResponseWriter, req *http.Request) {
	defer func(start time.Time) {
		duration := time.Since(start)
		c.logger.Info("ping function runtime", zap.Duration("duration", duration))
	}(time.Now())
	reqBody, _ := ioutil.ReadAll(req.Body)
	var pong Ping
	json.Unmarshal(reqBody, &pong)
	u1, err := uuid.FromString(pong.ID)
	if err != nil {
		c.logger.Error("uuid not valid", zap.String("incorrect_uid", pong.ID))
		error := Ping{
			Message:   "id not valid must be v4 uid",
			TimeStamp: time.Now().Format(time.RFC3339Nano),
		}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(400)
		return
	}
	splash := Ping{
		ID:        u1.String(),
		Message:   "splash",
		TimeStamp: time.Now().Format(time.RFC3339Nano),
	}
	json.NewEncoder(w).Encode(splash)
	w.WriteHeader(200)
}

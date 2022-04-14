package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	c := controller{logger: logger}
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/ping", c.pingFunc).Methods("POST")
	myRouter.HandleFunc("/pong", c.pongFunc).Methods("POST")

	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.Fatal("Failed to start invalid port", zap.Error(err))
	}
	logger.Info("Started server", zap.Int("port", portInt))

	d := dropper{logger: logger}
	dropList := os.Getenv("DROPS_PRESCRIPTION")
	if len(dropList) > 0 {
		drops := strings.Split(dropList, ";")
		for _, prescription := range drops {
			pres := strings.Split(prescription, "->")
			interval, err := time.ParseDuration(pres[0])
			if err != nil {
				logger.Error("Couldn't load drop skipping", zap.String("invalid_prescription",
					prescription), zap.Error(err))
			}
			drop := Drop{
				interval: interval,
				url:      pres[1],
			}
			go d.Drop(drop)
		}
	}

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter)
	if err != nil {
		logger.Fatal("failed in listen and serve", zap.Error(err))
		return
	}

}

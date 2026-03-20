package main

import (
	"time"

	"monitoring-system/scripts/agent/collector"
	"monitoring-system/scripts/agent/config"
	"monitoring-system/scripts/agent/sender"
)

func main() {
	cfg := config.Load()

	s := sender.NewSender(cfg.APIURL)

	for {
		metric := collector.Collect()

		s.Send(metric)

		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}

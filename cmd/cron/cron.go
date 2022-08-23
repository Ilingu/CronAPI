package cron

import (
	"bytes"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

type cronManager struct {
}

func NewCronManager() *cronManager {
	return &cronManager{}
}

var totalCron int

func (manager *cronManager) NewCronJob(frequency string, callbackUrl string) (*cron.Cron, error) {
	var c = cron.New()
	cronId := totalCron + 1

	err := c.AddFunc(frequency, func() {
		resp, err := http.Post(callbackUrl, "text/plain", bytes.NewBuffer([]byte{})) // Webhook

		// if client doen't respond to the webhook call, stop the cronJob
		if err != nil || resp.StatusCode != http.StatusOK || resp.Header.Get("Continue") != "true" {
			c.Stop()
			totalCron--
			log.Printf("[LOG]: Cron #%d Stopped\n", cronId)
		}
	})
	if err != nil {
		return nil, err
	}
	c.Start()

	totalCron++
	log.Printf("[LOG]: Cron #%d Started\n", cronId)
	return c, nil
}

package cron

import (
	"bytes"
	"net/http"

	"github.com/robfig/cron"
)

type cronManager struct {
}

func NewCronManager() *cronManager {
	return &cronManager{}
}

func (manager *cronManager) NewCronJob(frequency string, callbackUrl string) *cron.Cron {
	var c = cron.New()

	c.AddFunc(frequency, func() {
		resp, err := http.Post(callbackUrl, "text/plain", bytes.NewBuffer([]byte{})) // Webhook

		// if client doen't respond to the webhook call, stop the cronJob
		if err != nil || resp.StatusCode != http.StatusOK || resp.Header.Get("Continue") != "true" {
			c.Stop()
		}
	})
	c.Start()

	return c
}

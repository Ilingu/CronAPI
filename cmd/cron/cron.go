package cron

import (
	"bytes"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

type CronJob struct {
	Id         int
	Identifier string
	Cron       *cron.Cron
}

var ActiveCron, lastCronId = map[string]CronJob{}, 0

func NewCronJob(frequency string, callbackUrl string) (*CronJob, error) {
	if activeJob, active := ActiveCron[callbackUrl]; active {
		return &activeJob, nil
	}

	var c = cron.New()
	err := c.AddFunc(frequency, func() {
		resp, err := http.Post(callbackUrl, "text/plain", bytes.NewBuffer([]byte{})) // Webhook

		// if client doen't respond to the webhook call, stop the cronJob
		if err != nil || resp.StatusCode != http.StatusOK || resp.Header.Get("Continue") != "true" {
			DeleteCronJob(callbackUrl)
		}
	})

	if err != nil {
		return nil, err
	}

	cronInstance := CronJob{Id: lastCronId + 1, Identifier: callbackUrl, Cron: c}
	cronInstance.StartCron()
	return &cronInstance, nil
}

func DeleteCronJob(callbackUrl string) bool {
	activeJob, active := ActiveCron[callbackUrl]
	if !active || activeJob == (CronJob{}) {
		return false
	}

	activeJob.StopCron()
	return true
}

func (c CronJob) StartCron() {
	c.Cron.Start()
	ActiveCron[c.Identifier] = c
	lastCronId = c.Id

	log.Printf("[LOG]: Cron #%d Started. ID: %s\n", c.Id, c.Identifier)
}
func (c CronJob) StopCron() {
	c.Cron.Stop()
	delete(ActiveCron, c.Identifier)

	log.Printf("[LOG]: Cron #%d Stopped. ID: %s\n", c.Id, c.Identifier)
}

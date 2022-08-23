package main

import (
	"cron-api/cmd/cron"
	"net/http"
)

type Routes struct {
}

var cronManager = cron.NewCronManager()

func (*Routes) addJob(res http.ResponseWriter, req *http.Request) {
	body, err := ReadBody(req)
	if err != nil {
		WriteResponse(&res, http.StatusBadRequest, "Invalid Body")
		return
	}

	_, err = cronManager.NewCronJob(body.CronFrequency, body.CallbackUrl)
	if err != nil {
		WriteResponse(&res, http.StatusInternalServerError, "Failed to summon the CronJob")
		return
	}

	WriteResponse(&res, http.StatusCreated, "Job Successfully Created!")
}

package main

import (
	"cron-api/cmd/cron"
	"cron-api/cmd/utils"
	"net/http"
	"net/url"
)

type Routes struct {
}

func (*Routes) addJob(res http.ResponseWriter, req *http.Request) {
	body, err := ReadBody(req)
	if err != nil {
		WriteResponse(&res, http.StatusBadRequest, "Invalid Body")
		return
	}

	_, err = cron.NewCronJob(body.CronFrequency, body.CallbackUrl)
	if err != nil {
		WriteResponse(&res, http.StatusInternalServerError, "Failed to summon the CronJob")
		return
	}

	WriteResponse(&res, http.StatusCreated, "Job Successfully Created!")
}

func (*Routes) delJob(res http.ResponseWriter, req *http.Request) {
	rawIdentifier := req.URL.Query().Get("identifier")
	identifier, err := url.QueryUnescape(rawIdentifier)
	if err != nil || !utils.IsValidUrl(identifier) {
		WriteResponse(&res, http.StatusInternalServerError, "Invalid identifier")
		return
	}

	ok := cron.DeleteCronJob(identifier)
	if !ok {
		WriteResponse(&res, http.StatusInternalServerError, "Failed to delete the CronJob")
		return
	}

	WriteResponse(&res, http.StatusCreated, "Job Successfully Deleted!")
}

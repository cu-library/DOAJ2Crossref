package main

import (
	"time"
)

type TemplateData struct {
	HeadData
}

type HeadData struct {
	DOIBatch       int64
	Timestamp      int64
	DepositorName  string
	DepositorEmail string
	Registrant     string
}

func NewTemplateData(depositorName, depositorEmail, registrant string) TemplateData {
	return TemplateData{
		HeadData{
			DOIBatch:       time.Now().UTC().Unix(),
			Timestamp:      time.Now().UTC().UnixNano(),
			DepositorName:  depositorName,
			DepositorEmail: depositorEmail,
			Registrant:     registrant,
		},
	}
}

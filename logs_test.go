/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/docker/machine/log"
	oaocs "github.com/Noxaro/oneandone-cloudserver-api"
)

func TestIntegrationLogs(t *testing.T) {

	logsLastH := getLogs(t, "LAST_HOUR")
	logsLast24H := getLogs(t, "LAST_24H")
	logsLast7D := getLogs(t, "LAST_7D")
	logsLast30D := getLogs(t, "LAST_30D")
	logsLast365D := getLogs(t, "LAST_365D")


	for index, _ := range logsLastH {
		assert.Equal(t, logsLastH[index].Id, logsLast24H[index].Id)
		assert.Equal(t, logsLastH[index].Id, logsLast7D[index].Id)
		assert.Equal(t, logsLastH[index].Id, logsLast30D[index].Id)
		assert.Equal(t, logsLastH[index].Id, logsLast365D[index].Id)
	}

	for index, _ := range logsLast24H {
		assert.Equal(t, logsLast24H[index].Id, logsLast7D[index].Id)
		assert.Equal(t, logsLast24H[index].Id, logsLast30D[index].Id)
		assert.Equal(t, logsLast24H[index].Id, logsLast365D[index].Id)
	}

	for index, _ := range logsLast7D {
		assert.Equal(t, logsLast7D[index].Id, logsLast30D[index].Id)
		assert.Equal(t, logsLast7D[index].Id, logsLast365D[index].Id)
	}

	for index, _ := range logsLast30D {
		assert.Equal(t, logsLast30D[index].Id, logsLast365D[index].Id)
	}

	var i = 0
	for index, _ := range logsLast7D {
		i++
		if i == 60 {
			break
		}
		cLog, err := GetAPI().GetLog(logsLast7D[index].Id)
		log.Debug(logsLast7D[index])
		log.Debug(cLog)

		assert.Nil(t, err)
		assert.Equal(t, logsLast7D[index].Id, cLog.Id)
		assert.Equal(t, logsLast7D[index].StartDate, cLog.StartDate)
		assert.Equal(t, logsLast7D[index].EndDate, cLog.EndDate)
		assert.Equal(t, logsLast7D[index].Duration, cLog.Duration)
		assert.Equal(t, logsLast7D[index].Status.Percent, cLog.Status.Percent)
		assert.Equal(t, logsLast7D[index].Status.State, cLog.Status.State)
		assert.Equal(t, logsLast7D[index].Action, cLog.Action)
		assert.Equal(t, logsLast7D[index].Type, cLog.Type)
		assert.Equal(t, logsLast7D[index].Resource.Id, cLog.Resource.Id)
		assert.Equal(t, logsLast7D[index].Resource.Name, cLog.Resource.Name)
		assert.Equal(t, logsLast7D[index].User.Id, cLog.User.Id)
		assert.Equal(t, logsLast7D[index].User.Name, cLog.User.Name)
		assert.Equal(t, logsLast7D[index].CloudPanelId, cLog.CloudPanelId)

		time.Sleep(1 * time.Second)
	}
}

func makeMap(logs []oaocs.Log) (map[string]oaocs.Log){
	result := make(map[string]oaocs.Log)
	for index,_ := range logs {
		result[logs[index].Id] = logs[index]
	}
	return result
}

func getLogs(t *testing.T, dRange string) (map[string]oaocs.Log) {
	result, err := GetAPI().GetLogs(dRange)
	assert.Nil(t, err)
	return makeMap(result)
}
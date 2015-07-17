/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestIntegrationLogs(t *testing.T) {
	logs, err := GetAPI().GetLogs("LAST_30D")
	assert.Nil(t, err)
	var i = 0
	for index, _ := range logs {
		i++
		if i == 30 {
			break
		}
		log, err := GetAPI().GetLog(logs[index].Id)
		assert.Nil(t, err)
		assert.Equal(t, logs[index].Id, log.Id)
		assert.Equal(t, logs[index].StartDate, log.StartDate)
		assert.Equal(t, logs[index].EndDate, log.EndDate)
		assert.Equal(t, logs[index].Duration, log.Duration)
		assert.Equal(t, logs[index].Status.Percent, log.Status.Percent)
		assert.Equal(t, logs[index].Status.State, log.Status.State)
		assert.Equal(t, logs[index].Action, log.Action)
		assert.Equal(t, logs[index].Type, log.Type)
		assert.Equal(t, logs[index].Resource.Id, log.Resource.Id)
		assert.Equal(t, logs[index].Resource.Name, log.Resource.Name)
		assert.Equal(t, logs[index].User.Id, log.User.Id)
		assert.Equal(t, logs[index].User.Name, log.User.Name)
		assert.Equal(t, logs[index].CloudPanelId, log.CloudPanelId)

		time.Sleep(1 * time.Second)
	}
}
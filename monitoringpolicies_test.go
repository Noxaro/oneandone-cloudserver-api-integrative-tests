/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	oaocs "github.com/Noxaro/oneandone-cloudserver-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationCreateMonitoringPolicy(t *testing.T) {
	monitoringPolicy := oaocs.MonitoringPolicy{
		Name:  "Integration Test Monitoring Policy",
		Email: "",
		Agent: false,
		Thresholds: oaocs.MonitoringThreshold{
			Cpu: oaocs.MonitoringLevel{
				Warning: oaocs.MonitoringValue{
					Value: 90,
					Alert: false,
				},
				Critical: oaocs.MonitoringValue{
					Value: 95,
					Alert: false,
				},
			},
			Ram: oaocs.MonitoringLevel{
				Warning: oaocs.MonitoringValue{
					Value: 90,
					Alert: false,
				},
				Critical: oaocs.MonitoringValue{
					Value: 95,
					Alert: false,
				},
			},
			Disk: oaocs.MonitoringLevel{
				Warning: oaocs.MonitoringValue{
					Value: 80,
					Alert: false,
				},
				Critical: oaocs.MonitoringValue{
					Value: 90,
					Alert: false,
				},
			},
			Transfer: oaocs.MonitoringLevel{
				Warning: oaocs.MonitoringValue{
					Value: 1000,
					Alert: false,
				},
				Critical: oaocs.MonitoringValue{
					Value: 2000,
					Alert: false,
				},
			},
			InternalPing: oaocs.MonitoringLevel{
				Warning: oaocs.MonitoringValue{
					Value: 50,
					Alert: false,
				},
				Critical: oaocs.MonitoringValue{
					Value: 100,
					Alert: false,
				},
			},
		},
	}

	createdMonitoringPolicy, err := GetAPI().CreateMonitoringPolicy(monitoringPolicy)
	assert.Nil(t, err)

	monitoringPolicies, err := GetAPI().GetMonitoringPolicies()
	assert.Nil(t, err)

	if len(monitoringPolicies) == 0 {
		t.Error("")
	}

	for index, _ := range monitoringPolicies {
		time.Sleep(1 * time.Second)
		monitoringPolicy, err := GetAPI().GetMonitoringPolicy(monitoringPolicies[index].Id)
		assert.Nil(t, err)
		assert.Equal(t, monitoringPolicies[index].Id, monitoringPolicy.Id)
		assert.Equal(t, monitoringPolicies[index].Name, monitoringPolicy.Name)
		assert.Equal(t, monitoringPolicies[index].Description, monitoringPolicy.Description)
		assert.Equal(t, monitoringPolicies[index].Default, monitoringPolicy.Default)
		assert.Equal(t, monitoringPolicies[index].State, monitoringPolicy.State)
		assert.Equal(t, monitoringPolicies[index].CreationDate, monitoringPolicy.CreationDate)
		assert.Equal(t, monitoringPolicies[index].Email, monitoringPolicy.Email)
		assert.Equal(t, monitoringPolicies[index].Agent, monitoringPolicy.Agent)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Cpu.Warning.Value, monitoringPolicy.Thresholds.Cpu.Warning.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Cpu.Warning.Alert, monitoringPolicy.Thresholds.Cpu.Warning.Alert)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Cpu.Critical.Value, monitoringPolicy.Thresholds.Cpu.Critical.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Cpu.Critical.Alert, monitoringPolicy.Thresholds.Cpu.Critical.Alert)

		assert.Equal(t, monitoringPolicies[index].Thresholds.Ram.Warning.Value, monitoringPolicy.Thresholds.Ram.Warning.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Ram.Warning.Alert, monitoringPolicy.Thresholds.Ram.Warning.Alert)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Ram.Critical.Value, monitoringPolicy.Thresholds.Ram.Critical.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Ram.Critical.Alert, monitoringPolicy.Thresholds.Ram.Critical.Alert)

		assert.Equal(t, monitoringPolicies[index].Thresholds.Transfer.Warning.Value, monitoringPolicy.Thresholds.Transfer.Warning.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Transfer.Warning.Alert, monitoringPolicy.Thresholds.Transfer.Warning.Alert)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Transfer.Critical.Value, monitoringPolicy.Thresholds.Transfer.Critical.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.Transfer.Critical.Alert, monitoringPolicy.Thresholds.Transfer.Critical.Alert)

		assert.Equal(t, monitoringPolicies[index].Thresholds.InternalPing.Warning.Value, monitoringPolicy.Thresholds.InternalPing.Warning.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.InternalPing.Warning.Alert, monitoringPolicy.Thresholds.InternalPing.Warning.Alert)
		assert.Equal(t, monitoringPolicies[index].Thresholds.InternalPing.Critical.Value, monitoringPolicy.Thresholds.InternalPing.Critical.Value)
		assert.Equal(t, monitoringPolicies[index].Thresholds.InternalPing.Critical.Alert, monitoringPolicy.Thresholds.InternalPing.Critical.Alert)

	}

	createdMonitoringPolicy.Delete()

}

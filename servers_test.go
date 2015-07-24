package main

import (
	oaocs "github.com/Noxaro/oneandone-cloudserver-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var serverId = ""

func TestIntegrativeCreateServer(t *testing.T) {

	serverName := "IT Test Server " + GetTimestamp()
	latestAppliance, err := GetAPI().ServerApplianceFindNewest("Linux", "Ubuntu", "Minimal", 64, true)
	assert.Nil(t, err)
	config := oaocs.ServerCreateData{
		Name:        serverName,
		ApplianceId: latestAppliance.Id,
		Hardware: oaocs.Hardware{
			Vcores:            1,
			CoresPerProcessor: 1,
			Ram:               1,
			Hdds: []oaocs.Hdd{
				oaocs.Hdd{
					Size:   40,
					IsMain: true,
				},
			},
		},
		PowerOn: true,
	}

	server, err := GetAPI().CreateServer(config)
	assert.Nil(t, err)
	server.WaitForState("POWERED_ON")

	serverId = server.Id
}

func TestIntegrationServerGet(t *testing.T) {
	servers, err := GetAPI().GetServers()
	assert.Nil(t, err)
	for index, _ := range servers {
		time.Sleep(1 * time.Second)
		server, err := GetAPI().GetServer(servers[index].Id)
		assert.Nil(t, err)
		assert.Equal(t, servers[index].Id, server.Id)
		assert.Equal(t, servers[index].Description, server.Description)
		assert.Equal(t, servers[index].Hardware.CoresPerProcessor, server.Hardware.CoresPerProcessor)
		assert.Equal(t, servers[index].Hardware.Ram, server.Hardware.Ram)
		assert.Equal(t, servers[index].Hardware.Vcores, server.Hardware.Vcores)
		assert.Equal(t, servers[index].Image.Id, server.Image.Id)
		assert.Equal(t, servers[index].Image.Name, server.Image.Name)
		assert.Equal(t, servers[index].Status.Percent, server.Status.Percent)
		assert.Equal(t, servers[index].Status.State, server.Status.State)
	}
}

func TestIntegrationServerLifecycle(t *testing.T) {
	server, err := GetAPI().GetServer(serverId)
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_ON", server.Status)

	server.Shutdown(false)
	server.WaitForState("POWERED_OFF")
	state, err := server.GetStatus()
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_OFF", state)

	server.Start()
	server.WaitForState("POWERED_ON")
	state, err = server.GetStatus()
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_ON", state)

	server.Shutdown(true)
	server.WaitForState("POWERED_OFF")
	state, err = server.GetStatus()
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_OFF", state)

	server.Start()
	server.WaitForState("POWERED_ON")
	state, err = server.GetStatus()
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_ON", state)

	server.Reboot(false)
	server.WaitForState("POWERED_ON")
	state, err = server.GetStatus()
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_ON", state)

	server.Reboot(true)
	server.WaitForState("POWERED_ON")
	state, err = server.GetStatus()
	assert.Nil(t, err)
	assert.Equal(t, "POWERED_ON", state)
}

func TestIntegrationServerDelete(t *testing.T) {
	server, err := GetAPI().GetServer(serverId)
	assert.Nil(t, err)
	result, err := server.Delete()
	assert.Nil(t, err)
	result.WaitUntilDeleted()
}
